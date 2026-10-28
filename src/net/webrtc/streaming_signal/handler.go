package streaming_signal

import (
	"context"
	"log"

	"LibreRemotePlay/src/net/events"

	"github.com/pion/webrtc/v4"
)

func HandleStreamingSignal(ctx context.Context, streamingSignalChannel *webrtc.DataChannel) {

	if streamingSignalChannel.Label() != "streaming-signal" {
		return
	}

	go handleWhipOffer(ctx, streamingSignalChannel)
	
	localCtx, cancelLocalCtx := context.WithCancel(ctx)
	defer cancelLocalCtx()
	
	events.EventsOn(localCtx, events.EVENT_STREAMING_SIGNAL_SERVER, func(data []byte) {
		if WhipConfig.Enabled.IsEnabled() {
			return
		}

		err := streamingSignalChannel.SendText(string(data))
		if err != nil {
			log.Println(err)
			cancelLocalCtx()
			return
		}
	})
	
	streamingSignalChannel.OnMessage(func(msg webrtc.DataChannelMessage) {

		// TODO: see if msg is reaching here ****
		
		log.Println(string(msg.Data))
		
		if WhipConfig.Enabled.IsEnabled() {
			handleWhipAnswer(msg.Data)
		} else {
			err := events.EventsEmit(ctx, events.EVENT_STREAMING_SIGNAL_CLIENT ,msg.Data)
			if err != nil {
				log.Println(err)
				return
			}
		}

	})

}
