// Websocket server
package events

import (
	"LibreRemotePlay/src/cli"
	"LibreRemotePlay/src/state"
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/coder/websocket"
)

func setupListenHandler(app state.App, eventChan chan internalEvent, mux *http.ServeMux) {

	acceptOptions := &websocket.AcceptOptions{}
	
	if app.Dev {
		acceptOptions.OriginPatterns = []string{"localhost:5173"}
	}

	if cli.GetConfig().GetNetworkVisible() {
		acceptOptions.OriginPatterns = []string{}
		acceptOptions.InsecureSkipVerify = true	
	}

	type topicSlice struct {
		M sync.RWMutex
		Map map[uint16]*websocket.Conn
		N uint16
	}
	
	topics := sync.Map{}
	
	mux.HandleFunc("GET /{topic}", func(w http.ResponseWriter, r *http.Request) {

		clientGone := r.Context().Done()
		
		c, err := websocket.Accept(w, r, acceptOptions)
		
		if err != nil {
			log.Println(err)
			return
		}

		defer c.CloseNow()
		
		topic := topicT(r.PathValue("topic"))
		log.Println("Init listener topic: ", topic)
		
		topicSliceAny, _ := topics.LoadOrStore(topic, &topicSlice{Map: map[uint16]*websocket.Conn{}})
		topicMap := topicSliceAny.(*topicSlice)

		topicMap.M.Lock()
		
		log.Printf("Topic %s: %d\n", topic, topicMap.N)
		topicN := topicMap.N
		topicMap.N += 1
		topicMap.Map[topicN] = c
		
		topicMap.M.Unlock()
		
		defer func() {
			topicMap.M.Lock()
			defer topicMap.M.Unlock()

			delete(topicMap.Map, topicN)
		}()
		
	    for {
	        select {
	        case <-clientGone:
	            log.Println("Client disconnected from event ", topic)
	            return
			case <-app.Ctx.Done():
				return
			default:
				c.Read(app.Ctx)
				return
			}
	    }

	})

	go func() {
		for e := range eventChan {
			topicMapAny, _ := topics.LoadOrStore(e.event, &topicSlice{Map: map[uint16]*websocket.Conn{}})		
			topicMap := topicMapAny.(*topicSlice)
			log.Printf("Received event %s: %s\n", e.event, e.payload)

			func() {
				topicMap.M.RLock()
				defer topicMap.M.RUnlock()
				log.Println(topicMap.Map)
				for _, c := range topicMap.Map {
					err := c.Write(context.Background(), websocket.MessageText, e.payload)
					
					if err != nil {
						log.Println(err)
						return
					}
				}
			}()
		}
	}()

}