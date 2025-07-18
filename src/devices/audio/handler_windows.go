package audio

import (
	"context"
	"fmt"
	"time"

	"github.com/PiterWeb/RemoteController/src/bin"
	"github.com/amenzhinsky/go-memexec"
	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media"
)

func HandleAudio(ctx context.Context, track *webrtc.TrackLocalStaticSample) error {
	
	appLoopbackExe, err := memexec.New(bin.StdoutPCMApplicationLoopback_exe)
	if err != nil {
		return err
	}
	defer appLoopbackExe.Close()

	pid := ctx.Value("pid").(uint32)

	cmd := appLoopbackExe.Command(fmt.Sprintf("%d", pid))

	stdoutPipe, err := cmd.StdoutPipe()

	if err != nil {
		return  err
	}

	buff := make([]byte, 1024)

	tickerChan := time.NewTicker(time.Millisecond * 20).C

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-tickerChan:
			if n, err := stdoutPipe.Read(buff); err != nil {
				return err
			} else if n == 0 {
				continue
			}
	
			if err = track.WriteSample(media.Sample{Data: buff, Duration: time.Millisecond * 20}); err != nil {
				return err
			}
		}
	}

}