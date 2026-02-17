// Package webrtc provides all the related functions for WebRTC communication
package webrtc

import (
	"context"
	"log"
	"strings"

	// "github.com/PiterWeb/RemoteController/src/plugins"
	// "github.com/PiterWeb/RemoteController/src/devices/audio"
	"github.com/PiterWeb/RemoteController/src/devices/gamepad"
	"github.com/PiterWeb/RemoteController/src/devices/keyboard"
	"github.com/PiterWeb/RemoteController/src/devices/mouse"
	"github.com/PiterWeb/RemoteController/src/net/webrtc/streaming_signal"
	"github.com/pion/webrtc/v4"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var defaultSTUNServers = []string{"stun:stun.l.google.com:19305", "stun:stun.l.google.com:19302", "stun:stun.ipfire.org:3478"}

const ERROR_ANSWER = "ERROR"

func InitHost(ctx context.Context, ICEServers []webrtc.ICEServer, offerEncodedWithCandidates string, answerResponse chan<- string, triggerEnd <-chan struct{}, pidChan <-chan uint32) {

	candidates := []webrtc.ICECandidateInit{}

	if len(ICEServers) == 0 {
		ICEServers = []webrtc.ICEServer{
			{
				URLs: defaultSTUNServers,
			},
		}
	}

	streaming_signal.WhipConfig.ICEServers.Store(&ICEServers)

	// Prepare the configuration
	config := webrtc.Configuration{
		ICEServers: ICEServers,
	}

	closedConnChan := make(chan struct{})

	defer func() {
		if err := recover(); err != nil {
			answerResponse <- ERROR_ANSWER
			closedConnChan <- struct{}{}
		}
	}()

	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		return
	}

	defer func() {
		if err := peerConnection.Close(); err != nil {
			log.Printf("cannot close peerConnection: %v\n", err)
		}
	}()

	// audioTrack, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypePCMA, Channels: 2}, "audio", "app-audio")

	// if err != nil {
	// 	panic(err)
	// } else if _, err := peerConnection.AddTrack(audioTrack); err != nil {
	// 	panic(err)
	// }

	// Defer close of the wails signaling channel
	defer runtime.EventsOff(ctx, "streaming-signal-server")

	// Reload plugins in case a new plugin was added or configuration changed
	// plugins.ReloadPlugins()

	// Register data channel creation handling
	peerConnection.OnDataChannel(func(d *webrtc.DataChannel) {

		gamepad.HandleGamepad(d)
		streaming_signal.HandleStreamingSignal(ctx, d)
		keyboard.HandleKeyboard(d)
		mouse.HandleMouse(d)
		// plugins.HandleServerPlugins(d)

	})

	peerConnection.OnICECandidate(func(c *webrtc.ICECandidate) {

		if c == nil {
			answerResponse <- signalEncode(*peerConnection.LocalDescription()) + ";" + signalEncode(candidates)
			return
		}

		candidates = append(candidates, (*c).ToJSON())

	})

	// Set the handler for Peer connection state
	// This will notify you when the peer has connected/disconnected
	peerConnection.OnConnectionStateChange(func(s webrtc.PeerConnectionState) {
		log.Printf("Peer Connection State has changed: %s\n", s.String())

		runtime.EventsEmit(ctx, "connection_state", s.String())

		if s == webrtc.PeerConnectionStateFailed {
			closedConnChan <- struct{}{}

			peerConnection.Close()

		}

		if s == webrtc.PeerConnectionStateClosed {
			closedConnChan <- struct{}{}
		}
	})

	offerEncodedWithCandidatesSplited := strings.Split(offerEncodedWithCandidates, ";")

	offer := webrtc.SessionDescription{}
	signalDecode(offerEncodedWithCandidatesSplited[0], &offer)

	var receivedCandidates []webrtc.ICECandidateInit

	signalDecode(offerEncodedWithCandidatesSplited[1], &receivedCandidates)

	if err := peerConnection.SetRemoteDescription(offer); err != nil {
		panic(err)
	}

	for _, candidate := range receivedCandidates {
		if err := peerConnection.AddICECandidate(candidate); err != nil {
			log.Println(err)
			continue
		}
	}

	// Create an answer to send to the other process
	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		panic(err)
	}

	// Sets the LocalDescription, and starts our UDP listeners
	err = peerConnection.SetLocalDescription(answer)
	if err != nil {
		panic(err)
	}

	// var audioCtx context.Context
	var cancelAudioCtx context.CancelFunc = func() {}

	for {
		select {
		case pid := <-pidChan:
			log.Printf("Audio pid: %d\n", pid)
			cancelAudioCtx()
			// Pid value to not stream audio
			if pid == 0 {
				continue
			}
			// audioCtx, cancelAudioCtx = context.WithCancel(context.WithValue(context.Background(), "pid", pid))
			// go func() {
				// if err := audio.HandleAudio(audioCtx, audioTrack); err != nil {
				// 	log.Println(err)
				// }
			// }()
		case <-triggerEnd: // Block until cancel by user
			answerResponse <- ERROR_ANSWER
			cancelAudioCtx()
			return
		case <-closedConnChan: // Block until failed/clossed peerconnection
			answerResponse <- ERROR_ANSWER
			cancelAudioCtx()
			return
		}
	}

}
