package warrant

import (
	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/internal/datagram/management"
)

type Administration struct {
	AdminName                 string
	SigningCertificate        []byte
	SignatureAlgorithm        piv.Algorithm
	AuthenticationCertificate []byte
	AuthenticationAlgorithm   piv.Algorithm
	EncryptionKey             []byte
}

func (*Administration) Type() management.Datagram {
	return management.AdminWarrant
}
