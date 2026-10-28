package events

import (
	"io"
	"log"
	"net/http"
)

func setupEmitHandler(eventChan chan<- internalEvent, mux *http.ServeMux) {

	mux.HandleFunc("POST /{topic}", func(w http.ResponseWriter, r *http.Request) {

		topic := topicT(r.PathValue("topic"))

		rawBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		defer r.Body.Close()
		
		eventChan <- internalEvent{
			event: topic,
			payload: rawBody,
		}

		log.Printf("Emmit event %s: %s\n", topic, string(rawBody))
		
	})
	
}