package devices

import "sync"

type DeviceEnabledI interface {
	Toogle()
	IsEnabled() bool
}

type DeviceEnabled struct {
	enabled bool
	m       sync.Mutex
	DeviceEnabledI
}

func (d *DeviceEnabled) Toogle() {
	d.m.Lock()
	defer d.m.Unlock()
	d.enabled = !d.enabled
}

func (d *DeviceEnabled) IsEnabled() bool {
	d.m.Lock()
	defer d.m.Unlock()
	return d.enabled
}

func (d *DeviceEnabled) Enable() *DeviceEnabled {
	d.m.Lock()
	defer d.m.Unlock()
	d.enabled = true
	return d
}

func (d *DeviceEnabled) Disable() *DeviceEnabled {
	d.m.Lock()
	defer d.m.Unlock()
	d.enabled = false
	return d
}
