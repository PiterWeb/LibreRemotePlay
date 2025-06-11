package oninit

import (
	"embed"
	"net/http"
	
	LRPSignals "github.com/PiterWeb/LibreRemotePlaySignals/v1"
	net_http "github.com/PiterWeb/RemoteController/src/net/http"
	"github.com/PiterWeb/RemoteController/src/net/websocket"
)

func Execute(assets embed.FS) error {

	serverPort := 8080

	httpServerMux := http.NewServeMux()
	
	websocket.SetupWebsocketHandler(httpServerMux)

	errChan := make(chan error)
	defer close(errChan)

	go func() {
		err := net_http.InitHTTPAssets(httpServerMux, serverPort, assets)

		errChan <- err
	}()

	easyConnectPort := uint16(8081)

	go func() {
		ips_channel := make(chan []string)
		defer close(ips_channel)

		err := LRPSignals.InitServer(easyConnectPort, ips_channel)
		errChan <- err
	}()

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil

}
