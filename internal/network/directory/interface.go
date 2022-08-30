package directory

import "github.com/jmg292/G-Net/internal/datagram/management/warrant"

type NetworkDirectory interface {
	GetAdminWarrantById([]byte) warrant.Administration
	GetDeviceWarrantByID([]byte) warrant.DeviceWarrant
}
