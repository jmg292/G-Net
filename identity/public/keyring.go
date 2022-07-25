package public

import (
	"crypto/x509"
)

type userCertificates struct {
	Name                      string
	SigningCertificate        *x509.Certificate
	AuthenticationCertificate *x509.Certificate
}

func ParseCertificates(name string, signingCertBytes []byte, authCertBytes []byte) (*userCertificates, error) {
	var err error
	keyring := userCertificates{Name: name}
	if keyring.SigningCertificate, err = x509.ParseCertificate(signingCertBytes); err != nil {
		return nil, err
	}
	if keyring.AuthenticationCertificate, err = x509.ParseCertificate(authCertBytes); err != nil {
		return nil, err
	}
	return &keyring, nil
}
