package gamepad

import (
	"math"
)

type ID byte

type All [ControllerCount]XInputState

type XInputState struct {
	ID        ID
	Connected bool
	Packet    uint32
	Raw       RawControls
}

type RawControls struct {
	Buttons      Button
	LeftTrigger  uint8
	RightTrigger uint8
	ThumbLX      int16
	ThumbLY      int16
	ThumbRX      int16
	ThumbRY      int16
}

func (state XInputState) toXInput(virtualState *ViGEmState) {

	virtualState.DwPacketNumber = uint32(state.Packet)
	virtualState.Gamepad.updateFromRawState(state.Raw)
}

func (state *XInputState) Pressed(button Button) bool { return state.Raw.Buttons&button != 0 }

type Thumb struct{ X, Y, Magnitude float32 }

func (state *XInputState) RectDPad() (thumb Thumb) {
	if state.Pressed(DPadUp) {
		thumb.Y += 1
	}
	if state.Pressed(DPadDown) {
		thumb.Y -= 1
	}
	if state.Pressed(DPadLeft) {
		thumb.X -= 1
	}
	if state.Pressed(DPadRight) {
		thumb.X += 1
	}
	if thumb.X != 0 || thumb.Y != 0 {
		thumb.Magnitude = 1
	}
	return
}

func (state *XInputState) RoundDPad() (thumb Thumb) {
	thumb = state.RectDPad()
	if thumb.X != 0 && thumb.Y != 0 {
		thumb.X *= isqrt2
		thumb.Y *= isqrt2
	}
	return
}

func round16(rx, ry, deadzone int16) (thumb Thumb) {
	//TODO: use sqrt32
	fx, fy := float64(rx), float64(ry)
	thumb.Magnitude = float32(math.Sqrt(fx*fx + fy*fy))

	thumb.X = float32(rx) / thumb.Magnitude
	thumb.Y = float32(ry) / thumb.Magnitude

	if thumb.Magnitude > float32(deadzone) {
		if thumb.Magnitude > 32767 {
			thumb.Magnitude = 32767
		}
		thumb.Magnitude = (thumb.Magnitude - float32(deadzone)) / float32(32767-deadzone)
	} else {
		thumb.Magnitude = 0
	}

	thumb.X *= thumb.Magnitude
	thumb.Y *= thumb.Magnitude

	return
}

func (state *XInputState) RoundLeft() Thumb {
	return round16(state.Raw.ThumbLX, state.Raw.ThumbLY, LeftThumbDeadZone)
}

func (state *XInputState) RoundRight() Thumb {
	return round16(state.Raw.ThumbRX, state.Raw.ThumbRY, RightThumbDeadZone)
}

func linear16(v, deadzone int16) float32 {
	if v < -deadzone {
		return float32(v+deadzone) / float32(32767-deadzone)
	}
	if v > deadzone {
		return float32(v-deadzone) / float32(32767-deadzone)
	}
	return 0
}

func rect16(rx, ry, deadzone int16) (thumb Thumb) {
	thumb.X = linear16(rx, deadzone)
	thumb.Y = linear16(ry, deadzone)
	if thumb.X != 0 && thumb.Y != 0 {
		thumb.Magnitude = 1
	}
	return
}

func (state *XInputState) RectLeft() Thumb {
	return rect16(state.Raw.ThumbLX, state.Raw.ThumbLY, LeftThumbDeadZone)
}

func (state *XInputState) RectRight() Thumb {
	return rect16(state.Raw.ThumbRX, state.Raw.ThumbRY, RightThumbDeadZone)
}

// func (state *XInputState) Vibrate(left, right uint16) {
// 	if !state.Connected {
// 		return
// 	}
// 	Vibrate(state.ID, &Vibration{left, right})
// }

type Vibration struct {
	LeftMotor  uint16
	RightMotor uint16
}

const (
	ControllerCount    = ID(4)
	TriggerThreshold   = 30
	LeftThumbDeadZone  = 7849
	RightThumbDeadZone = 8689

	sqrt2  = 1.4142135623730950488
	isqrt2 = 1 / sqrt2
)

type Button uint16

const (
	DPadUp    Button = 0x0001
	DPadDown  Button = 0x0002
	DPadLeft  Button = 0x0004
	DPadRight Button = 0x0008

	Start Button = 0x0010
	Back  Button = 0x0020

	LeftThumb  Button = 0x0040
	RightThumb Button = 0x0080

	LeftShoulder  Button = 0x0100
	RightShoulder Button = 0x0200

	ButtonA Button = 0x1000
	ButtonB Button = 0x2000
	ButtonX Button = 0x4000
	ButtonY Button = 0x8000
)
