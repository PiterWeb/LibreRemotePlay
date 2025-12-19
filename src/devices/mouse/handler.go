package mouse

import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/PiterWeb/RemoteController/src/devices"
	"github.com/go-vgo/robotgo"
	"github.com/pion/webrtc/v4"
)

var MouseEnabled = new(devices.DeviceEnabled).Disable()

func HandleMouse(d *webrtc.DataChannel) error {

	if d.Label() != "mouse" {
		return nil
	}

	d.OnOpen(func() {
		log.Println("mouse data channel is open")
	})

	d.OnMessage(func(msg webrtc.DataChannelMessage) {

		if !MouseEnabled.IsEnabled() {
			return
		}

		if msg.Data == nil || msg.IsString {
			return
		}
		
		if len(msg.Data) <= 1 {
			return
		}

		msgBuf := bytes.NewBuffer(msg.Data)
		
		msgType, err := msgBuf.ReadByte()
		
		if err != nil {
			return
		}
		
		// Handle click event
		if msgType == typeMsgClick {
			
			clickBtn, err := msgBuf.ReadByte()
			
			if err != nil {
				return
			}
			
			btnState, err := msgBuf.ReadByte()
			
			if err != nil {
				return
			}
			
			var state string
			
			if btnState == mouseDown {
				state = "down"
			} else if btnState == mouseUp {
				state = "up"
			}
			
			switch clickBtn {
				case mouseLeft:
				robotgo.Toggle("left", state)
				case mouseCentral:
				robotgo.Toggle("center", state)
				case mouseRight:
				robotgo.Toggle("right", state)
			}
			
		} else if msgType == typeMsgMove { // Handle move event
			
			// TODO: do some logging to check the values
			
			x, err := binary.ReadUvarint(msgBuf)
			
			if err != nil {
				return
			}
			
			y, err := binary.ReadUvarint(msgBuf)
			
			if err != nil {
				return
			}
			
			robotgo.Move(int(x), int(y))
			
		}

	})

	return nil

}
