package streaming_signal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/PiterWeb/RemoteController/src/devices"
	"github.com/pion/webrtc/v4"
)

type whipConfig struct {
	Port       uint16
	OfferChan  chan string
	AnswerChan chan string
	Enabled    *devices.DeviceEnabled
	ICEServers atomic.Pointer[[]webrtc.ICEServer]
}

var whipConfigEnabled = &devices.DeviceEnabled{}

var WhipConfig = &whipConfig{
	Port:       8082,
	OfferChan:  make(chan string),
	AnswerChan: make(chan string),
	Enabled:    whipConfigEnabled.Disable(),
	ICEServers: atomic.Pointer[[]webrtc.ICEServer]{},
}

func InitWhipServer(config *whipConfig) error {

	var answerChan <-chan string = config.AnswerChan
	var offerChan chan<- string = config.OfferChan

	defer func() {
		close(config.OfferChan)
		close(config.AnswerChan)
	}()

	httpServerMux := http.NewServeMux()

	httpServerMux.HandleFunc("/whip", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST")
		w.Header().Add("Access-Control-Allow-Headers", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Authorization")
		
		if r.Method == http.MethodOptions {
			return
		}
		
		if !WhipConfig.Enabled.IsEnabled() {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Whip not enabled"))
			return
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Fatal error %s\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				_, _ =w.Write([]byte("Fatal error"))
				return
			}
		}()

		for _, s := range *WhipConfig.ICEServers.Load() {
			for _, url := range s.URLs {
				w.Header().Add("Link", fmt.Sprintf("<%s>; rel=\"ice-server\"", url))
			}
		}

		offer, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Error reading request body"))
			return
		}

		if len(offer) == 0 {
			log.Println("Offer received in whip is too short")
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Error request body too short"))
			return
		}
		
		log.Printf("Offer received in whip: %s\n", string(offer))
		offerChan <- string(offer)

		log.Println("Waiting for answer in whip")
		
		rawAnswer, ok := <-answerChan

		if !ok {
			log.Println("Answer channel closed in whip")
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Error getting an answer"))
			return
		}

		type answerT struct {
			Answer struct {
				SDP string
			}
		}

		answerStruct := answerT{}

		err = json.Unmarshal([]byte(rawAnswer), &answerStruct)

		if err != nil {
			log.Printf("Error parsing answer on whip %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Error getting an answer"))
			return
		}

		log.Printf("Sending answer for whip: %s\n", answerStruct.Answer.SDP)

		w.Header().Set("Content-Type", "application/sdp")
		w.Header().Add("Location", "/whip")

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(answerStruct.Answer.SDP))

	})

	httpServer := &http.Server{
		Handler:      httpServerMux,
		Addr:         fmt.Sprintf("127.0.0.1:%d", config.Port),
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
	}

	err := httpServer.ListenAndServe()
	return err

}

func handleWhipOffer(streamingSignalChannel *webrtc.DataChannel) {

	for offer := range WhipConfig.OfferChan {
		if WhipConfig.Enabled.IsEnabled() {

			offerMap := map[string]any{}

			offerSdp := webrtc.SessionDescription{
				Type: webrtc.SDPTypeOffer, SDP: offer,
			}

			offerMap["offer"] = offerSdp
			offerMap["role"] = "host"
			offerMap["type"] = "offer"

			offerJSONBytes, err := json.Marshal(offerMap)

			if err != nil {
				log.Printf("Error encoding whip offer: %s\n", err)
				continue
			}

			if err := streamingSignalChannel.SendText(string(offerJSONBytes)); err != nil {
				break
			}

		}
	}

}

func handleWhipAnswer(msg []byte) {

	type dataT struct {
		Type string
	}

	dataStruct := dataT{}

	err := json.Unmarshal(msg, &dataStruct)

	if err != nil {
		return
	}

	if dataStruct.Type == "answer" {
		WhipConfig.AnswerChan <- string(msg)
	}

}
