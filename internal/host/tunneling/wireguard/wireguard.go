package wireguard

import (
	"sync"

	"golang.org/x/sys/windows"
	"golang.zx2c4.com/wireguard/tun"
)

type Tunnel struct {
	device tun.Device
}

var (
	instanceMutex = &sync.Mutex{}
	instance      *Tunnel
)

func NewTunnel() (tunnel *Tunnel, err error) {
	instanceMutex.Lock()
	defer instanceMutex.Unlock()

	if instance == nil {
		tun.WintunTunnelType = "G-Net"
		if guid, e := windows.GUIDFromString(WireguardAdapterGuid); e != nil {
			err = e
		} else {
			tun.WintunStaticRequestedGUID = &guid
			instance = &Tunnel{}
		}
	}

	tunnel = instance
	return
}
