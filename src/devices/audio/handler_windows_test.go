package audio

import (
	"context"
	"testing"
	"time"

	"github.com/pion/webrtc/v4"
)

func TestHandleAudio(t *testing.T) {

	audioProcess := GetAudioProcess()

	ctx, cancelCtx := context.WithCancel(context.WithValue(context.Background(), "pid", audioProcess[len(audioProcess)-1].Pid))

	audioTrack, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypePCMA, Channels: 2}, "audio", "app-audio")

	if err != nil {
		t.Errorf("Error creating audio track: %s", err.Error())
	}

	go func() {
		t.Error(HandleAudio(ctx, audioTrack))
	}()

	time.Sleep(5 * time.Second)

	cancelCtx()

}
