package keyboard

import "strings"

func mapJSKeyToRobotGo(jsKey string) string {
	keyMap := map[string]string{
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
		"Insert":     "insert",
		"Delete":     "delete",
		"Home":       "home",
		"End":        "end",
		"PageUp":     "pageup",
		"PageDown":   "pagedown",
	}

	// Map (F1 - F12) keys
	if strings.HasPrefix(jsKey, "F") && len(jsKey) > 1 {
		return strings.ToLower(jsKey) // "F1" → "f1"
	}

	if val, exists := keyMap[jsKey]; exists {
		return val
	}
	return ""
}
