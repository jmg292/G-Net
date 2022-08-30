package entries

import (
	"time"

	"github.com/jmg292/G-Net/internal/datagram/management/publication"
	"github.com/jmg292/G-Net/pkg/keyring"
)

type DeviceManifest struct {
	Identity               keyring.PublicKeyRing
	NetworkInfo            publication.NetworkInfo
	NetworkInfoLastUpdated time.Time
	PeeringInfo            publication.PeeringInfo
	PeeringInfoLastUpdated time.Time
}
