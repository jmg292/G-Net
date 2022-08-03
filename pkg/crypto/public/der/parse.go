package der

import "crypto/x509"

func ParseCertificates(name string, signingCertBytes []byte, authCertBytes []byte) (*keyRing, error) {
	var err error
	keyring := keyRing{name: name}
	if keyring.signingCert, err = x509.ParseCertificate(signingCertBytes); err != nil {
		return nil, err
	}
	if keyring.authenticationCert, err = x509.ParseCertificate(authCertBytes); err != nil {
		return nil, err
	}
	return &keyring, nil
}
