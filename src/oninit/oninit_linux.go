package oninit

import (
	"embed"
	"log"
	"net/http"

	LRPSignals "github.com/PiterWeb/LibreRemotePlaySignals/v1"
	net_http "github.com/PiterWeb/RemoteController/src/net/http"
	"github.com/PiterWeb/RemoteController/src/net/websocket"
)

func Execute(assets embed.FS) error {

	serverPort := 8080

	httpServerMux := http.NewServeMux()
	
	websocket.SetupWebsocketHandler(httpServerMux)

	ips_channel := make(chan []string, 1)
	errChan := make(chan error, 2)
	defer close(errChan)
	defer close(ips_channel)

	go func() {
		err := net_http.InitHTTPAssets(httpServerMux, serverPort, assets)

		if err != nil {
			errChan <- err
		}
	}()

	easyConnectPort := uint16(8081)
	
	go func() {

		options := LRPSignals.ServerOptions{
			Port: easyConnectPort,
		}
		
		log.Println("Easy Connect Server started on port 8081")
		err := LRPSignals.InitServer(options, ips_channel)
		if err != nil {
			errChan <- err
		}
	}()

	return <-errChan

}
