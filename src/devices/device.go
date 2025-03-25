package devices

import (
	"sync/atomic"
)

type DeviceEnabledI interface {
	Toogle()
	IsEnabled() bool
}

type DeviceEnabled struct {
	enabled int32
	DeviceEnabledI
}

func (d *DeviceEnabled) Toogle() {
	atomic.StoreInt32(&d.enabled, 1-atomic.LoadInt32(&d.enabled))
}

func (d *DeviceEnabled) IsEnabled() bool {
	return atomic.LoadInt32(&d.enabled) == 1
}

func (d *DeviceEnabled) Enable() *DeviceEnabled {
	atomic.StoreInt32(&d.enabled, 1)
	return d
}

func (d *DeviceEnabled) Disable() *DeviceEnabled {
	atomic.StoreInt32(&d.enabled, 0)
	return d
}
