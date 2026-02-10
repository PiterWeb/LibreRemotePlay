package oninit

import (
	"embed"
	"log"
	
	LRPSignals "github.com/PiterWeb/LibreRemotePlaySignals/v1"
	"github.com/PiterWeb/RemoteController/src/cli"
	
)

func Execute(assets embed.FS) error {

	ips_channel := make(chan []string, 1)
	defer close(ips_channel)

	log.Println("Easy Connect Server started on port 8081")
	options := LRPSignals.ServerOptions{
		Port: cli.GetConfig().GetEasyConnectPort(),
	}
	err := LRPSignals.InitServer(options, ips_channel)

	return err

}
