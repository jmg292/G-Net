package warrant

import (
	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/internal/datagram/management"
)

type DeviceWarrant struct {
	DeviceName                string
	OwnerFingerprint          []byte
	SigningCertificate        []byte
	SignatureAlgorithm        piv.Algorithm
	AuthenticationCertificate []byte
	AuthenticationAlgorithm   piv.Algorithm
}

func (*DeviceWarrant) Type() management.Datagram {
	return management.DeviceWarrant
}
