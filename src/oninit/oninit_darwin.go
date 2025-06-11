package oninit

import (
	"embed"

	LRPSignals "github.com/PiterWeb/LibreRemotePlaySignals/v1"
)

func Execute(assets embed.FS) error {

	easyConnectPort := uint16(8081)
	ips_channel := make(chan []string)
	defer close(ips_channel)

	err := LRPSignals.InitServer(easyConnectPort, ips_channel)

	return err

}
