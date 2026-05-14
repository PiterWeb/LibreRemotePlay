package mouse

import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func moveMouseHandler(screens []runtime.Screen) (func(x, y, videoW, videoH int32), func() error) {
	
	return func(x, y, videoW, videoH int32) {
		
	}, func() error {
		return nil
	}
}