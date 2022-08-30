package certificate

import (
	"crypto/x509"

	gnet "github.com/jmg292/G-Net/pkg/gneterrs"
	"github.com/jmg292/G-Net/pkg/identity/certificate/extensions"
)

type ProvableOrigin interface {
	Certificate() *x509.Certificate
	ProofOfOrigin() (*x509.Certificate, error)
}

func GetProofOfOrigin(certificate Extensible) (proofOfOrigin *x509.Certificate, err error) {
	if extProofOfOrigin, e := findExtensionByOID(certificate, extensions.OIDProofOfOrigin); e != nil {
		err = e
	} else if extProofOfOrigin.Value == nil {
		err = gnet.ErrorCertificateNotFound
	} else {
		proofOfOrigin, err = x509.ParseCertificate(extProofOfOrigin.Value)
	}
	return
}
