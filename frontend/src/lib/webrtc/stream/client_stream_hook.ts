import { exportStunServers } from '$lib/webrtc/stun_servers';
import { setConsumingStream, type SignalingData } from '$lib/webrtc/stream/stream_signal_hook.svelte';
import { exportTurnServers } from '$lib/webrtc/turn_servers';
import { getSortedVideoCodecs} from './stream_config.svelte';
import LANMode from '$lib/webrtc/lan_mode.svelte';

let peerConnection: RTCPeerConnection | undefined;
let inboundStream: MediaStream | null = null;

function initStreamingPeerConnection() {
	if (peerConnection) {
		peerConnection.close();
	}

	peerConnection = new RTCPeerConnection({
		iceServers: LANMode.enabled ? [] : [...exportStunServers(), ...exportTurnServers()]
	});
}

async function CreateClientStream(
	signalingChannel: RTCDataChannel,
	videoElement: HTMLVideoElement
) {
	initStreamingPeerConnection();

	if (!videoElement || !peerConnection) throw new Error('Error creating stream');

	const videoTransceiver = peerConnection.addTransceiver("video");
  const audioTransceiver = peerConnection.addTransceiver("audio");
  
  monitorAndAdaptAudioCodec(audioTransceiver.sender)
	videoTransceiver.setCodecPreferences(getSortedVideoCodecs());

	peerConnection.onconnectionstatechange = () => {
		if (!peerConnection) return;

		const connectionTerminatedOptions: RTCPeerConnectionState[] = ["disconnected", "failed", "closed"]

		if (connectionTerminatedOptions.includes(peerConnection.connectionState)) {
			CloseStreamClientConnection()
		}
	};

	peerConnection.onicecandidate = (e) => {
		if (!e.candidate) return;

		const data: SignalingData = {
			type: 'candidate',
			candidate: e.candidate.toJSON(),
			role: 'client'
		};

		signalingChannel.send(JSON.stringify(data));
	};

	peerConnection.ontrack = (ev) => {

		if (ev.streams && ev.streams[0]) {
			ev.streams[0].getTracks().forEach(t => t.addEventListener("ended", () => {CloseStreamClientConnection()}, true) )
			videoElement.srcObject = ev.streams[0];
			videoElement.play();
		} else {
			if (!inboundStream) {
				inboundStream = new MediaStream();
				videoElement.srcObject = inboundStream;
				videoElement.play();
			}
			ev.track.addEventListener("ended", () => {CloseStreamClientConnection()}, true)
			inboundStream.addTrack(ev.track);
			inboundStream.getTracks().forEach(t => t.addEventListener("ended", () => {CloseStreamClientConnection()}, true))
		}
	};


	const offer = await peerConnection.createOffer({
		offerToReceiveAudio: true,
		offerToReceiveVideo: true,
		iceRestart: true
	});

	await peerConnection.setLocalDescription(offer);

	const data: SignalingData = {
		type: 'offer',
		offer: offer,
		role: 'client'
	};

	signalingChannel.send(JSON.stringify(data));

	signalingChannel.onmessage = async (e) => {
		const { type, answer, candidate, role } = JSON.parse(e.data) as SignalingData;

		if (!peerConnection) {
			return;
		}

		if (role !== 'host') {
			return;
		}

		switch (type) {
			case 'answer':
				if (!answer) return;
				await peerConnection.setRemoteDescription(answer);
				break;
			case 'candidate':
				try {await peerConnection.addIceCandidate(candidate)} catch {/** */}
				break;
		}
	};


}

// Example of monitoring and adapting audio codec parameters
function monitorAndAdaptAudioCodec(audioSender: RTCRtpSender) {
  // Monitor network conditions
  setInterval(async () => {
    const stats = await audioSender.getStats();
    let availableBitrate = 0;

    // Calculate packet loss using remote-inbound-rtp stats
    // (outbound-rtp doesn't know about lost packets)
    stats.forEach(report => {
      if (report.type === 'candidate-pair' && report.state === 'succeeded') {
        availableBitrate = report.availableOutgoingBitrate;
      }
    });

    // Adapt Opus parameters based on conditions
    const parameters = audioSender.getParameters();

    // Find Opus codec in parameters
    const opusEncodingIdx = parameters.encodings.findIndex(() => 
      parameters.codecs.find(c => c.mimeType.toLowerCase() === 'audio/opus')
    );

    if (opusEncodingIdx >= 0) {
      // Adjust bitrate based on available bandwidth
      if (availableBitrate > 0) {
        // Leave headroom for other traffic
        const targetBitrate = Math.min(128000, availableBitrate * 0.7);
        parameters.encodings[opusEncodingIdx].maxBitrate = targetBitrate;
      }

      // Note: We're using standard RTCRtpEncodingParameters rather than
      // non-standard properties like 'networkPriority' which are
      // Chromium-only and behind flags

      // Apply the changes
      audioSender.setParameters(parameters);
    }
  }, 2000); // Check every 2 seconds
}

function CloseStreamClientConnection() {
	setConsumingStream(false)
	if (!peerConnection) return;
	peerConnection.close();
	peerConnection = undefined;
}

export { CreateClientStream, CloseStreamClientConnection };
