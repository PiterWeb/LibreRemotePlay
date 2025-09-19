package streaming_signal

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/PiterWeb/RemoteController/src/devices"
)

type whipConfig struct {
	Port uint16
	OfferChan chan string
	AnswerChan chan string
	Enabled devices.DeviceEnabled
}

var WhipConfig *whipConfig = &whipConfig{}

func InitWhipServer(config whipConfig) error {
	
	var answerChan <-chan string = config.AnswerChan
	var offerChan chan<- string = config.OfferChan
	
	defer func() {
		close(config.OfferChan)
		close(config.AnswerChan)
	}()
	
	httpServerMux := http.NewServeMux()
	
	httpServerMux.HandleFunc("POST /whip", func(w http.ResponseWriter, r *http.Request) {
		
		if !WhipConfig.Enabled.IsEnabled() {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Whip not enabled"))
			return
		}

		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Fatal error"))
			}
		}()
		
		r.Header.Add("Access-Control-Allow-Origin", "*")
		r.Header.Add("Access-Control-Allow-Methods", "POST")
		r.Header.Add("Access-Control-Allow-Headers", "*")
		r.Header.Add("Access-Control-Allow-Headers", "Authorization")
		
		offer, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading request body"))
		}

		log.Printf("Offer received in whip: %s\"n", string(offer))
		offerChan <- string(offer)
		
		answer, ok := <- answerChan
		
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error getting an answer"))
		}
		
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(answer))
		
	})
	
	httpServer := &http.Server{
		Handler: httpServerMux,
		Addr:    fmt.Sprintf("127.0.0.1:%d", config.Port),
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 30,
	}
	
	err := httpServer.ListenAndServe()
	return err
	
}