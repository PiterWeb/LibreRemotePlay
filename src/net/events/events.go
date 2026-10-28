package events

import (
	"LibreRemotePlay/src/cli"
	"LibreRemotePlay/src/state"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/coder/websocket"
)

type internalEvent struct {
	event topicT
	payload []byte
	n uint16
}

type topicT string

const (
	EVENT_CONNECTION_STATE topicT = "connection_state"
	EVENT_ERROR topicT = "ERROR" 
	EVENT_STREAMING_SIGNAL_CLIENT topicT = "streaming-signal-client"
	EVENT_STREAMING_SIGNAL_SERVER topicT = "streaming-signal-server"
)

func SetupEventHandler(app state.App, mux *http.ServeMux) {

	eventChan := make(chan internalEvent)

	go func() {
		<-app.Ctx.Done()
		close(eventChan)
	}()
	
	listenMux := http.NewServeMux()

	mux.Handle("/listen/", http.StripPrefix("/listen", listenMux))
	setupListenHandler(app, eventChan, listenMux)

	emitMux := http.NewServeMux()

	mux.Handle("/emmit/", http.StripPrefix("/emmit", emitMux))
	setupEmitHandler(eventChan, emitMux)
}

func createWS(ctx context.Context, eventName topicT) (*websocket.Conn, error) {
	wsUrl := fmt.Sprintf("ws://localhost:%d/ws/%s", cli.GetConfig().GetHTTPPort(), eventName)

	wsClient, _, err := websocket.Dial(ctx, wsUrl, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return  wsClient, nil
}

func EventsEmit(ctx context.Context, eventName topicT, data []byte) error {

	ws, err := createWS(ctx, eventName)
	if err != nil {
		return err
	}

	defer ws.Close(websocket.StatusNormalClosure, "")
	
	return ws.Write(ctx, websocket.MessageText, data)

}

func EventsOn(ctx context.Context, eventName topicT, callback func (data []byte)) error {

	ws, err := createWS(ctx, eventName)
	if err != nil {
		return err
	}

	go func() {

		defer ws.Close(websocket.StatusNormalClosure, "")
		
		for {
			t, msg ,err := ws.Read(ctx)

			if err != nil {
				log.Println(err)
				return
			}

			if t != websocket.MessageText {
				continue
			}

			callback(msg)
		
		}
		
	}()

	return nil
}

func EventsOnce(ctx context.Context, eventName topicT, callback func(data []byte)) {
	localCtx, cancelLocalCtx := context.WithCancel(ctx)
	EventsOn(localCtx, eventName, func(data []byte) {
		defer cancelLocalCtx()
		callback(data)
	})
}