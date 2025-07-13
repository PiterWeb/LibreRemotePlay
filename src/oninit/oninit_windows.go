package oninit

import (
	"embed"
	"log"

	LRPSignals "github.com/PiterWeb/LibreRemotePlaySignals/v1"
	"github.com/PiterWeb/RemoteController/src/devices/gamepad"
)

func Execute(assets embed.FS) error {
	err := gamepad.InitViGEm()

	if err != nil {
		return err
	}

	easyConnectPort := uint16(8081)
	ips_channel := make(chan []string, 1)
	defer close(ips_channel)

	log.Println("Easy Connect Server started on port 8081")
	options := LRPSignals.ServerOptions{
		Port: easyConnectPort,
	}
	err = LRPSignals.InitServer(options, ips_channel)

	return err

}
