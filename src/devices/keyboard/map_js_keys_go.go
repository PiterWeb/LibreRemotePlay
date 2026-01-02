package keyboard

import "strings"

var keyMap = map[string]string{
	"Enter":      "enter",
	"Escape":     "esc",
	"Backspace":  "backspace",
	"Tab":        "tab",
	" ":          "space",
	"ArrowUp":    "up",
	"ArrowDown":  "down",
	"ArrowLeft":  "left",
	"ArrowRight": "right",
	"Shift":      "shift",
	"Control":    "ctrl",
	"Alt":        "alt",
	"CapsLock":   "capslock",
}

func mapJSKeyToRobotGo(jsKey string) (key string, exists bool) {

	// Map (F1 - F12) keys
	if strings.HasPrefix(jsKey, "F") && len(jsKey) > 1 {
		return strings.ToLower(jsKey), true // "F1" → "f1"
	}

	if val, exists := keyMap[jsKey]; exists {
		return val, true
	}

	if len(jsKey) == 1 {
		return jsKey, true
	}

	return "", false
}
