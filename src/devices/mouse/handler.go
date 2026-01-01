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

func HandleMouse(d *webrtc.DataChannel) {

	if d.Label() != "mouse" {
		return
	}

	d.OnOpen(func() {
		
		robotgo.MouseSleep = 100
		
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
			
			switch btnState {
				case mouseDown:
				state = "down"
				case mouseUp:
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
			
			x := make([]byte, 2)
			_, err := msgBuf.Read(x)
			
			if err != nil {
				return
			}
						
			y := make([]byte, 2)
			_, err = msgBuf.Read(y)
			
			if err != nil {
				return
			}
			
			xNum := binary.BigEndian.Uint16(x)
			yNum := binary.BigEndian.Uint16(y)
			
			// log.Printf("Mouse x: %d, y:%d\n", int(xNum), int(yNum))
			
			robotgo.Move(int(xNum), int(yNum))
			
		}

	})

}
