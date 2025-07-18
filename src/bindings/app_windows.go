package bindings

import (
	"github.com/PiterWeb/RemoteController/src/devices/audio"
	"github.com/PiterWeb/RemoteController/src/devices/gamepad"
)

func (a *App) OpenViGEmWizard() (err string) {

	return gamepad.OpenViGEmWizard().Error()

}

func (a *App) SetAudioPid(pid uint32) {
	pidAudioChan <- pid
}

func (a *App) GetAudioProcess() []audio.AudioProcess {
	return audio.GetAudioProcess()
}
