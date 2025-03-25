package oninit

import (
	"embed"

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

	return nil

}
