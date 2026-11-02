// Websocket server
package events

import (
	"LibreRemotePlay/src/cli"
	"LibreRemotePlay/src/state"
	"context"
	"log"
	"net/http"
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
					
			if app.Dev {
				return r.URL.Host == "localhost:5173"
			}
			
			return r.URL.Port() == strconv.FormatUint(uint64(cli.GetConfig().GetHTTPPort()), 10)
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
			c.Close()
		}()
		
		go func() {
			for {
				select {
					case <-localCtx.Done():
						return
					default:
						func() {
							listener.EventCond.L.Lock()
							defer listener.EventCond.L.Unlock()
						
							listener.EventCond.Wait()
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
						}()
				}
			}
		}()
		
	    for {
	        select {
			case <-localCtx.Done():
				return
			default:
				_, a, _ := c.ReadMessage()
				log.Printf("Received cancel(%s) from topic %s\n", string(a), topic)
				return
			}
	    }
				
	})
	
	go func() {
		for e := range eventChan {
			func(e internalEvent) {
				topicsMapLock.RLock()
				defer topicsMapLock.RUnlock()

				listener, ok := topicsMap[e.event]
				if !ok {
					return
				}

				listener.EventCond.L.Lock()
				listener.EventData = e.payload
				listener.EventCond.L.Unlock()	
				
				listener.EventCond.Broadcast()
				
			}(e)	
		}
	}()

}