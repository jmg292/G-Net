package der

import (
	"crypto/x509"
	"fmt"

	"github.com/jmg292/G-Net/utilities/gnet"
	"golang.org/x/crypto/sha3"
)

func (k *keyRing) Name() string {
	return k.name
}

func (k *keyRing) Fingerprint() []byte {
	digest := sha3.New256()
	digest.Sum(k.signingCert.Raw)
	return digest.Sum(k.authenticationCert.Raw)
}

func (k *keyRing) SigningCertificate() *x509.Certificate {
	return k.signingCert
}

func (k *keyRing) AuthenticationCertificate() *x509.Certificate {
	return k.authenticationCert
}

func (k *keyRing) VerifySignature(data []byte, signature []byte) error {
	return x509ValidateSignature(k.signingCert, data, signature)
}

func (k *keyRing) VerifyAuthentication(data []byte, signature []byte) error {
	return x509ValidateSignature(k.authenticationCert, data, signature)
}

func (k *keyRing) Encrypt(data []byte) ([]byte, error) {
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}
