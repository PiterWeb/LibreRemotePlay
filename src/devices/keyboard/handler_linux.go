package keyboard

import (
	"log"
	"strings"
	"sync"

	"LibreRemotePlay/src/devices"

	"github.com/jbdemonte/virtual-device/keyboard"
	"github.com/jbdemonte/virtual-device/linux"
	"github.com/pion/webrtc/v4"
)

var KeyboardEnabled = new(devices.DeviceEnabled).Disable()

var keyMapLinux = map[string]linux.Key{
	"Enter":      linux.KEY_ENTER,
	"Escape":     linux.KEY_ESC,
	"Backspace":  linux.KEY_BACKSPACE,
	"Tab":        linux.KEY_TAB,
	" ":          linux.KEY_SPACE,
	".":		  linux.KEY_DOT,
	",":		  linux.KEY_COMMA,
	"ArrowUp":    linux.KEY_UP,
	"ArrowDown":  linux.KEY_DOWN,
	"ArrowLeft":  linux.KEY_LEFT,
	"ArrowRight": linux.KEY_RIGHT,
	"Shift":      linux.KEY_LEFTSHIFT,
	"Control":    linux.KEY_LEFTCTRL,
	"Alt":        linux.KEY_LEFTALT,
	"CapsLock":   linux.KEY_CAPSLOCK,
	"Q":		  linux.KEY_Q,
	"W":		  linux.KEY_W,
	"E":		  linux.KEY_E,
	"R":		  linux.KEY_R,
	"T":		  linux.KEY_T,
	"Y":		  linux.KEY_Y,
	"U":		  linux.KEY_U,
	"I":		  linux.KEY_I,
	"O":		  linux.KEY_O,
	"P":		  linux.KEY_P,
	"A":		  linux.KEY_A,
	"S":		  linux.KEY_S,
	"D":		  linux.KEY_D,
	"F":		  linux.KEY_F,
	"G":		  linux.KEY_G,
	"H":		  linux.KEY_H,
	"J":		  linux.KEY_J,
	"K":		  linux.KEY_K,
	"L":		  linux.KEY_L,
	"Z":		  linux.KEY_Z,
	"X":		  linux.KEY_X,
	"C":		  linux.KEY_C,
	"V":		  linux.KEY_V,
	"B":		  linux.KEY_B,
	"N":		  linux.KEY_N,
	"M":		  linux.KEY_M,
	"F1":		  linux.KEY_F1,
	"F2":		  linux.KEY_F2,
	"F3":		  linux.KEY_F3,
	"F4":		  linux.KEY_F4,
	"F5":		  linux.KEY_F5,
	"F6":		  linux.KEY_F6,
	"F7":		  linux.KEY_F7,
	"F8":		  linux.KEY_F8,
	"F9":		  linux.KEY_F9,
	"F10":		  linux.KEY_F10,
	"F11":		  linux.KEY_F11,
	"F12":		  linux.KEY_F12,
}

func HandleKeyboard(d *webrtc.DataChannel) {

	if d.Label() != "keyboard" {
		return
	}

	var virtualKeyboard keyboard.VirtualKeyboard
	
	d.OnOpen(func() {
				
		log.Println("keyboard data channel is open")
		virtualKeyboard = keyboard.NewGenericKeyboard()
		err := virtualKeyboard.Register()
		if err != nil {
			log.Printf("Error registering virtual keyboard: %s\n", err)
			virtualKeyboard = nil
		}
	})

	keyState := make(map[linux.Key]bool)
	keyStateMutex := new(sync.Mutex)
	
	d.OnClose(func() {

		if virtualKeyboard == nil {
			return
		}
		
		err := virtualKeyboard.Unregister()

		if err != nil {
			log.Println(err)
		}
		
	})

	d.OnMessage(func(msg webrtc.DataChannelMessage) {

		if virtualKeyboard == nil {
			return
		}
		
		if !KeyboardEnabled.IsEnabled() {
			return
		}
		
		if !msg.IsString || msg.Data == nil {
			return
		}
		
		keyParts := strings.Split(string(msg.Data), "_")

		if len(keyParts) < 2 {
			return
		}
		
		key, ok := mapJSKeyToLinux(keyParts[0])

		if !ok {
			log.Println("keyboard key not found: ", keyParts[0])
			return
		}
		
		if keyParts[1] == "1" {
			keyStateMutex.Lock()
			defer keyStateMutex.Unlock()
			if keyState[key] {
				return
			}
			keyState[key] = true
			virtualKeyboard.PressKey(key)
			virtualKeyboard.SyncReport()
		} else {
			keyStateMutex.Lock()
			defer keyStateMutex.Unlock()
			if !keyState[key] {
				return
			}
			keyState[key] = false
			virtualKeyboard.ReleaseKey(key)
			virtualKeyboard.SyncReport()
		}

	})
}

func mapJSKeyToLinux(jsKey string) (key linux.Key, exists bool) {

	// Handle numbers
	if len(jsKey) == 1 {
		isZero := jsKey[0] == '0'
		isOtherNumber := jsKey[0] >= '1' && jsKey[0] <= '9'

		if isZero {
			return linux.KEY_0, true
		} else if isOtherNumber {
			return linux.KEY_1 + linux.Key(jsKey[0] - '1'), true
		} else {
			jsKey = strings.ToUpper(jsKey)
		}
	}
	
	linuxKey, ok := keyMapLinux[jsKey]
	
	if ok {
		return linuxKey, true
	}


	
	return linux.Key(0), false
}