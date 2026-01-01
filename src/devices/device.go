package devices

import (
	"sync/atomic"
)

type DeviceEnabled struct {
	enabled atomic.Bool
}

func (d *DeviceEnabled) Toogle() {
	for {
		old := d.enabled.Load()
		if d.enabled.CompareAndSwap(old, !old) {
			break
		}
	}
}

func (d *DeviceEnabled) IsEnabled() bool {
	return d.enabled.Load()
}

func (d *DeviceEnabled) Enable() *DeviceEnabled {
	d.enabled.Store(true)
	return d
}

func (d *DeviceEnabled) Disable() *DeviceEnabled {
	d.enabled.Store(false)
	return d
}
