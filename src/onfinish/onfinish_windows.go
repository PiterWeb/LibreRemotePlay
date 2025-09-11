package onfinish

import "github.com/PiterWeb/RemoteController/src/devices/gamepad"

func Execute() error {

	if err := gamepad.CloseViGEmDLL(); err != nil {
		return err
	}

	return nil
}