// Websocket server
package events

import (
	"LibreRemotePlay/src/cli"
	"LibreRemotePlay/src/state"
	"context"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/gorilla/websocket"
)

type EventListener struct {
	EventCond *sync.Cond
	EventData []byte
}

func setupListenHandler(app state.App, eventChan chan internalEvent, mux *http.ServeMux) {

	upgrader := websocket.Upgrader{
	    ReadBufferSize:  1024,
	    WriteBufferSize: 1024,
	    CheckOrigin: func(r *http.Request) bool {
			if cli.GetConfig().GetNetworkVisible() {
				return true
			}

			origin := r.Header.Get("Origin")

			originUrl, err := url.Parse(origin)
			if err != nil {
				return false
			}
			
			if app.Dev {
				return originUrl.Host == "localhost:5173"
			}
			
			return originUrl.Port() == strconv.FormatUint(uint64(cli.GetConfig().GetHTTPPort()), 10)
		},
	}

	topicsMap := map[topicT]*EventListener{}
	topicsMapLock := sync.RWMutex{}
	
	mux.HandleFunc("GET /{topic}", func(w http.ResponseWriter, r *http.Request) {	
		once := r.URL.Query().Get("once") == "true"

		c, err := upgrader.Upgrade(w, r ,nil)
		
		if err != nil {
			log.Println(err)
			return
		}

		defer c.Close()
		
		topic := topicT(r.PathValue("topic"))
		log.Println("Init listener topic: ", topic)
		defer log.Println("Client disconnected from event ", topic)

		topicsMapLock.Lock()

		listener, ok := topicsMap[topic]
		if !ok {
			listener = &EventListener{
				EventCond: sync.NewCond(&sync.Mutex{}),
				EventData: []byte{},
			}
			topicsMap[topic] = listener
		}
		
		topicsMapLock.Unlock()

		closed := atomic.Bool{}
		
		localCtx, cancelLocalCtx := context.WithCancel(r.Context())
		defer cancelLocalCtx()

		go func() {
			<-localCtx.Done()
			closed.Store(true)
		}()
		
		go func() {
			for {
				select {
					case <-localCtx.Done():
						return
					default:
						listener.EventCond.L.Lock()
						listener.EventCond.Wait()
						listener.EventCond.L.Unlock()
						
						if closed.Load() {
							return
						}	
						
						err = c.WriteMessage(websocket.TextMessage, listener.EventData)
						if err != nil {
							cancelLocalCtx()
						}
						
						if once {
							cancelLocalCtx()
						}
				}
			}
		}()
		
	    for {
	        select {
			case <-localCtx.Done():
				return
			default:
				_, _, _ = c.ReadMessage()
				log.Printf("Received cancel from topic %s\n", topic)
				return
			}
	    }
				
	})
	
	go func() {
		for e := range eventChan {
			topicsMapLock.RLock()

			listener, ok := topicsMap[e.event]
			if !ok {
				topicsMapLock.RUnlock()
				continue
			}
			
			topicsMapLock.RUnlock()

			listener.EventCond.L.Lock()
			listener.EventData = e.payload
			listener.EventCond.L.Unlock()	
			
			listener.EventCond.Broadcast()
		}
	}()

}