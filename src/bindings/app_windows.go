package bindings

import (
	"github.com/PiterWeb/RemoteController/src/devices/gamepad"
)

func (a *App) OpenViGEmWizard() (err string) {

	return gamepad.OpenViGEmWizard().Error()

}
