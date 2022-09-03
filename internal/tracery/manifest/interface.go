package manifest

import (
	"github.com/jmg292/G-Net/internal/datagram/management/edict"
	"github.com/jmg292/G-Net/internal/tracery/manifest/entry"
	"github.com/jmg292/G-Net/pkg/identity/certificate"
)

type Manifest interface {
	GetIdentity([]byte) (*certificate.Identity, error)
	GetNetworkExits([]byte) ([]*edict.NetworkExit, error)
	GetNetworkRoutes([]byte) ([]*edict.NetworkRoute, error)
	GetDeviceManifest([]byte) (*entry.DeviceManifest, error)
}
