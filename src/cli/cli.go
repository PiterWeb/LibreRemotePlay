package cli

import (
	"flag"
)

func init() {
	networkVisible := flag.Bool("network_visible", false, "Indicate to use all network interfaces (0.0.0.0) instead of loopback-only (127.0.0.1)")
	httpPort := flag.Uint("port", 8080, "Port used for serving http Web client")
	easyConnectPort := flag.Uint("easyport", 8081, "Port used for Easy Connect Server")
	
	flag.Parse()

	config = configT{
		networkVisible: *networkVisible,
		httpPort:       uint16(*httpPort),
		easyConnectPort: uint16(*easyConnectPort),
	}
}

type configT struct {
	networkVisible bool
	httpPort       uint16
	easyConnectPort uint16
}

func (c configT) GetNetworkVisible() bool {
	return c.networkVisible
}

func (c configT) GetHTTPPort() uint16 {
	return c.httpPort
}

func (c configT) GetEasyConnectPort() uint16 {
	return c.easyConnectPort
}

var config configT

func GetConfig() configT {
	return config
}
