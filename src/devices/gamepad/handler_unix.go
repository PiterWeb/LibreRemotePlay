//go:build unix

package gamepad

import (
	"log"

	"github.com/jbdemonte/virtual-device/gamepad"
	"github.com/json-iterator/go"
	"github.com/pion/webrtc/v3"
)

func HandleGamepad(gamepadChannel *webrtc.DataChannel) {

	if gamepadChannel.Label() != "controller" {
		return
	}

	var virtualGamepad gamepad.VirtualGamepad

	// Create a virtual device
	gamepadChannel.OnOpen(func() {

		var err error

		virtualGamepad, err = generateVirtualDevice()

		if err != nil {
			log.Println(err)
			log.Println("VirtualGamepad is not defined, gamepad disabled by default")
			GamepadEnabled.Disable()
		}

	})

	defer func() {
		if err := recover(); err != nil {
			if virtualGamepad != nil {
				virtualGamepad.Unregister()
			}
		}
	}()

	lastPad := GamepadAPIXState{
		Connected: false,
	}

	var actualPad GamepadAPIXState
	
	// Update the virtual device
	gamepadChannel.OnMessage(func(msg webrtc.DataChannelMessage) {

		if !GamepadEnabled.IsEnabled() {
			return
		}

		err := jsoniter.ConfigFastest.Unmarshal(msg.Data, &actualPad)
		
		if err != nil {
			log.Println(err)
			return
		}

		updateVirtualDevice(virtualGamepad, actualPad, lastPad)

		lastPad = actualPad

	})

	// Free the virtualGamepad
	gamepadChannel.OnClose(func() {

		if virtualGamepad == nil {
			return
		}

		err := virtualGamepad.Unregister()

		if err != nil {
			log.Println(err)
		}

	})
}

func generateVirtualDevice() (gamepad.VirtualGamepad, error) {

	g := gamepad.NewXBox360()

	err := g.Register()
	if err != nil {
		return nil, err
	}

	return g, nil

}

func updateVirtualDevice(virtualGamepad gamepad.VirtualGamepad, actualPad GamepadAPIXState, lastPad GamepadAPIXState) {

	for i, v := range actualPad.Axes {
		if actualPad.Axes[i] == lastPad.Axes[i] {
			continue
		}

		switch i {
		case 0:
			virtualGamepad.MoveLeftStickX(float32(-fixLYAxis(v)))
		case 1:
			virtualGamepad.MoveLeftStickY(float32(v))
		case 2:
			virtualGamepad.MoveRightStickX(float32(v))
		case 3:
			virtualGamepad.MoveRightStickY(float32(-fixRYAxis(v)))
		}

	}

	for i := range actualPad.Buttons {
		if actualPad.Buttons[i].Pressed == lastPad.Buttons[i].Pressed {
			continue
		}

		if actualPad.Buttons[i].Pressed {
			virtualGamepad.Press(buttonAPIXStateToVirtualGamepadButton[i])
		} else {
			virtualGamepad.Release(buttonAPIXStateToVirtualGamepadButton[i])
		}

	}

}

var buttonAPIXStateToVirtualGamepadButton = map[int]gamepad.Button{
	0:  gamepad.ButtonSouth,
	1:  gamepad.ButtonEast,
	2:  gamepad.ButtonWest,
	3:  gamepad.ButtonNorth,
	4:  gamepad.ButtonL1,
	5:  gamepad.ButtonR1,
	6:  gamepad.ButtonL2,
	7:  gamepad.ButtonR2,
	8:  gamepad.ButtonSelect,
	9:  gamepad.ButtonStart,
	10: gamepad.ButtonL3,
	11: gamepad.ButtonR3,
	12: gamepad.ButtonUp,
	13: gamepad.ButtonDown,
	14: gamepad.ButtonLeft,
	15: gamepad.ButtonRight,
}
