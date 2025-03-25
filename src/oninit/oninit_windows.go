package oninit

import (
	"embed"

	"github.com/PiterWeb/RemoteController/src/devices/gamepad"
)

func Execute(assets embed.FS) error {
	err := gamepad.InitViGEm()

	if err != nil {
		return err
	}

	return nil
}
