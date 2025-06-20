import { showToast, ToastType } from '$lib/toast/toast_hook';
import { get } from 'svelte/store';
import { setStreaming, type SignalingData } from '$lib/webrtc/stream/stream_signal_hook.svelte';
import { _ } from 'svelte-i18n';
import { exportStunServers } from '../stun_servers';
import { exportTurnServers } from '../turn_servers';
import { IS_RUNNING_EXTERNAL } from '$lib/detection/onwebsite';
import { DEFAULT_IDEAL_FRAMERATE, DEFAULT_MAX_FRAMERATE, FIXED_RESOLUTIONS, RESOLUTIONS } from './stream_config.svelte';
import ws from '$lib/websocket/ws';
import log from '$lib/logger/logger';
import LANMode from '$lib/webrtc/lan_mode.svelte';

let peerConnection: RTCPeerConnection | undefined;

function initStreamingPeerConnection() {
	if (peerConnection) {
		peerConnection.close();
	}

	peerConnection = new RTCPeerConnection({
		iceServers: LANMode.enabled ? [] : [...exportStunServers(), ...exportTurnServers()]
	});
}

let stream: MediaStream | undefined
let unlistenerStreamingSignal: (() => void) | undefined

async function getDisplayMediaStream(resolution: FIXED_RESOLUTIONS = FIXED_RESOLUTIONS.resolution720p, idealFrameRate = DEFAULT_IDEAL_FRAMERATE, maxFramerate = DEFAULT_MAX_FRAMERATE) {
	try {
		const mediastream = await navigator.mediaDevices.getDisplayMedia({
			video: { 
				frameRate: { ideal: idealFrameRate, max: maxFramerate },
				...(RESOLUTIONS.get(resolution) ?? {}),
				noiseSuppression: true, 
				autoGainControl: true,
			},
			audio: true,
		});

		return mediastream;
	} catch (e) {
		showToast(get(_)('error-starting-streaming'), ToastType.ERROR);
		return undefined;
	}
}

export function StopStreaming() {
	try {
		unlistenerStreamingSignal?.()
		unlistenerStreamingSignal = undefined
		setStreaming(false)
		stream?.getTracks().forEach(t => t.stop()) 

		if (!peerConnection) return;

		peerConnection.close();
		peerConnection = undefined;

		showToast(get(_)('streaming-stopped'), ToastType.SUCCESS);
	} catch (e) {
		showToast(get(_)('error-stopping-streaming'), ToastType.ERROR);
	}
}

export function CreateHostStream(resolution: FIXED_RESOLUTIONS = FIXED_RESOLUTIONS.resolution720p, idealFrameRate = DEFAULT_IDEAL_FRAMERATE, maxFramerate = DEFAULT_MAX_FRAMERATE) {
	initStreamingPeerConnection();

	if (!peerConnection) {
		throw new Error('Error creating stream');
	}

	peerConnection.onconnectionstatechange = async () => {
		if (!peerConnection) return;

		if (peerConnection.connectionState === 'connected') {
			showToast(get(_)('connected'), ToastType.SUCCESS);
			return;
		}

		const connectionTerminatedOptions: RTCPeerConnectionState[] = ["disconnected", "failed", "closed"]

		if (connectionTerminatedOptions.includes(peerConnection.connectionState)) {
			StopStreaming()
			return
		}
	};

	peerConnection.onicecandidate = async (event) => {
		if (event.candidate) {
			const data: SignalingData = {
				type: 'candidate',
				candidate: event.candidate.toJSON(),
				role: 'host'
			};

			if (IS_RUNNING_EXTERNAL) return ws().send(JSON.stringify(data));
			
			const { EventsEmit } = await import('$lib/wailsjs/runtime/runtime');
			EventsEmit('streaming-signal-server', JSON.stringify(data));
			return;
		}

		log('ICE gathering complete');

		const answer = peerConnection?.localDescription?.toJSON();
		const data: SignalingData = {
			type: 'answer',
			answer,
			role: 'host'
		};

		if (IS_RUNNING_EXTERNAL) return ws().send(JSON.stringify(data));

		const { EventsEmit } = await import('$lib/wailsjs/runtime/runtime');
		EventsEmit('streaming-signal-server', JSON.stringify(data));

		return
	};

	let offerArrived = false

	async function onSignalArrive(data: string) {
		if (!peerConnection) return;


		const { type, offer, candidate, role } = JSON.parse(data) as SignalingData;

		if (role !== 'client') return;

		if (type === "candidate") {
			try {peerConnection.addIceCandidate(candidate)} catch {/** */}
			return
		}

		if (type !== 'offer') return;
		if (!offer || offerArrived) return;

		try {

			await peerConnection.setRemoteDescription(offer);

		} catch (e) {
			// TODO: manage error
			log(e, {err: true})
			return
		}
		offerArrived = true;

		stream = await getDisplayMediaStream(resolution, idealFrameRate, maxFramerate);

		stream?.getTracks().forEach((track) => {
			track.addEventListener(
				'ended',
				() => {
					StopStreaming();
				},
				true
			)
			if (!stream) return;
			peerConnection?.addTrack(track, stream);
		});

		try {
			
			await peerConnection.setLocalDescription(await peerConnection.createAnswer());

		} catch (e) {
			log(e, {err: true})
			return
		}
		
	}

	if (IS_RUNNING_EXTERNAL) {
		const cllbck = (ev: MessageEvent<string>) =>  onSignalArrive(ev.data)
		ws().addEventListener("message", cllbck)
		unlistenerStreamingSignal = () => ws().removeEventListener("message", cllbck)
		return
	}

	(async () => {
		const { EventsOn } = await import('$lib/wailsjs/runtime/runtime');
		unlistenerStreamingSignal = EventsOn('streaming-signal-client', (data: string) => onSignalArrive(data));
	})()

}