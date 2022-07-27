package directory

import "github.com/jmg292/G-Net/internal/datagrams/management/warrant"

type NetworkDirectory interface {
	GetAdminWarrantById([]byte) warrant.Administration
	GetDeviceWarrantByID([]byte) warrant.DeviceWarrant
}
