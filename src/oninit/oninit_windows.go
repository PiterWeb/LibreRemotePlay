package oninit

import (
	"embed"

	LRPSignals "github.com/PiterWeb/LibreRemotePlaySignals/v1"
	"github.com/PiterWeb/RemoteController/src/devices/gamepad"
)

func Execute(assets embed.FS) error {
	err := gamepad.InitViGEm()

	if err != nil {
		return err
	}

	easyConnectPort := uint16(8081)
	ips_channel := make(chan []string)
	defer close(ips_channel)

	err = LRPSignals.InitServer(easyConnectPort, ips_channel)

	return err

}
