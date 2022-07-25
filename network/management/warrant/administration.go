package warrant

import (
	"github.com/jmg292/G-Net/network/management"

	"github.com/go-piv/piv-go/piv"
)

type Administration struct {
	AdminName                 string
	SigningCertificate        []byte
	SignatureAlgorithm        piv.Algorithm
	AuthenticationCertificate []byte
	AuthenticationAlgorithm   piv.Algorithm
}

func (*Administration) PacketType() management.PacketType {
	return management.AdminWarrant
}
