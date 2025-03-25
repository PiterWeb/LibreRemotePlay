package devices

import (
	"sync/atomic"
)

type DeviceEnabledI interface {
	Toogle()
	IsEnabled() bool
}

type DeviceEnabled struct {
	enabled atomic.Int32
	DeviceEnabledI
}

func (d *DeviceEnabled) Toogle() {
	d.enabled.Store(1 - d.enabled.Load())
}

func (d *DeviceEnabled) IsEnabled() bool {
	return d.enabled.Load() == 1
}

func (d *DeviceEnabled) Enable() *DeviceEnabled {
	d.enabled.Store(1)
	return d
}

func (d *DeviceEnabled) Disable() *DeviceEnabled {
	d.enabled.Store(0)
	return d
}
