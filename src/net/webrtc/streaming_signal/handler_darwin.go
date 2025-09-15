package streaming_signal

import (
	"context"

	"github.com/pion/webrtc/v4"
)

func HandleStreamingSignal(ctx context.Context, streamingSignalChannel *webrtc.DataChannel) {

	if streamingSignalChannel.Label() != "streaming-signal" {
		return
	}

}
