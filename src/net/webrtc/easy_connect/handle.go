package easyconnect

import "github.com/PiterWeb/LibreRemotePlaySignals/v1"

func HandleHostEasyConnect(s LRPSignals.ServerT, ID uint16, genHostCode func (LRPSignals.ClientCodeT) (LRPSignals.HostCodeT, error)) error {

	clientCode, err := LRPSignals.ReceiveClientCode(s, ID)

	if err != nil {
		return err
	}

	hostCode, err := genHostCode(clientCode)

	if err != nil {
		return err
	}

	err = LRPSignals.SendHostCode(s, hostCode, ID)

	return err
}

func HandleClientEasyConnect(s LRPSignals.ServerT, client_code LRPSignals.ClientCodeT, ID uint16) (LRPSignals.HostCodeT, error) {

	host_code, err := LRPSignals.SendClientCode(s, client_code, ID)

	if err != nil {
		return LRPSignals.HostCodeT{}, err
	}

	return host_code, nil
}