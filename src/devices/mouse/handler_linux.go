package mouse

import (
	"fmt"
	"math"

	"github.com/jbdemonte/virtual-device/mouse"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func moveMouseHandler(screens []runtime.Screen) (func(x, y, videoW, videoH int32), func() error) {
	
	width := int32(1920)
	height := int32(1080)
	
	for _, s := range screens {
		if !s.IsPrimary {
			continue
		}
		
		width = int32(s.Size.Width)
		height = int32(s.Size.Height)
	}
	
	m := mouse.NewGenericMouse()
	
	m.Register()
	
	var lastX int32 = math.MaxInt32
	var lastY int32 = math.MaxInt32
	
	return func(x, y, videoW, videoH int32) {
		
		fmt.Printf("LastX: %d, X:%d\n", lastX, x)
		
		if (lastX == math.MaxInt32 || lastY == math.MaxInt32) {
			lastX = x
			lastY = y
			return
		}
		
		tempX := x
		tempY := y
		x -= lastX
		y -= lastY
		
		x = x * (width/videoW)
		y = y * (height/videoH)
		
		
		fmt.Printf("Move x:%d y:%d\n", x, y)
		
		m.Move(x, y)
		
		lastX = tempX
		lastY = tempY
		
	}, m.Unregister
}