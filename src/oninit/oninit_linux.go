package oninit

import (
	"embed"

	LRPSignals "github.com/PiterWeb/LibreRemotePlaySignals/v1"
	"github.com/PiterWeb/RemoteController/src/net/http"
	"github.com/PiterWeb/RemoteController/src/net/websocket"
)

func Execute(assets embed.FS) error {

	serverPort := 8080

	websocket.SetupWebsocketHandler()

	err := http.InitHTTPAssets(serverPort, assets)

	if err != nil {
		return err
	}

	easyConnectPort := uint16(8081)
	ips_channel := make(chan []string)

	err = LRPSignals.InitServer(easyConnectPort, ips_channel)

	defer close(ips_channel)

	return err

}
