package public

import (
	"fmt"

	"github.com/jmg292/G-Net/utilities/gnet"
	"golang.org/x/crypto/sha3"
)

func (v *userCertificates) Fingerprint() []byte {
	digest := sha3.New256()
	digest.Sum(v.SigningCertificate.Raw)
	return digest.Sum(v.AuthenticationCertificate.Raw)
}

func (v *userCertificates) Verify(data []byte, signature []byte) error {
	return fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (v *userCertificates) Encrypt(data []byte)
