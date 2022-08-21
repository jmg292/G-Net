package request

import (
	"crypto/x509"

	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/keyring/certificate"
)

func NewCertificate(requester certificate.Requester, template *x509.Certificate) (*x509.Certificate, error) {
	return nil, gnet.ErrorNotYetImplemented
}

func Sign(signer certificate.Signer, request *x509.CertificateRequest) (*x509.Certificate, error) {
	return nil, gnet.ErrorNotYetImplemented
}
