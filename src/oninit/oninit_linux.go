package oninit

import (
	"embed"
	"log"
	"net/http"

	LRPSignals "github.com/PiterWeb/LibreRemotePlaySignals/v1"
	"github.com/PiterWeb/RemoteController/src/cli"
	net_http "github.com/PiterWeb/RemoteController/src/net/http"
	"github.com/PiterWeb/RemoteController/src/net/websocket"
)

func Execute(assets embed.FS) error {

	httpServerMux := http.NewServeMux()
	
	websocket.SetupWebsocketHandler(httpServerMux)

	ips_channel := make(chan []string, 1)
	errChan := make(chan error, 2)

	go func() {
		err := net_http.InitHTTPAssets(httpServerMux, assets)

		if err != nil {
			errChan <- err
		}
	}()
	
	go func() {

		options := LRPSignals.ServerOptions{
			Port: cli.GetConfig().GetEasyConnectPort(),
		}
		
		log.Printf("Easy Connect Server started on port %d\n", options.Port)
		err := LRPSignals.InitServer(options, ips_channel)
		if err != nil {
			errChan <- err
		}
	}()

	return <-errChan

}
