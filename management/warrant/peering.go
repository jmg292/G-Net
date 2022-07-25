package warrant

import (
	"github.com/jmg292/G-Net/management"

	"github.com/go-piv/piv-go/piv"
)

type Peering struct {
	DeviceName                string
	OwnerFingerprint          []byte
	SigningCertificate        []byte
	SignatureAlgorithm        piv.Algorithm
	AuthenticationCertificate []byte
	AuthenticationAlgorithm   piv.Algorithm
}

func (*Peering) PacketType() management.PacketType {
	return management.DeviceWarrant
}
