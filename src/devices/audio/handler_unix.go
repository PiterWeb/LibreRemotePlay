//go:build unix
package audio

import (
	"github.com/pion/webrtc/v3"
	"context"
)

func HandleAudio(ctx context.Context, track *webrtc.TrackLocalStaticSample) error {
	return nil
}

func GetAudioProcess() []AudioProcess {
	return []AudioProcess{
		{
			Name: "None",
			Pid: 0,
		},
	}
}