package certificate

import (
	"crypto"
	"crypto/x509"

	"github.com/jmg292/G-Net/pkg/keyring"
)

type Storage interface {
	GetCertificate(keyring.KeySlot) (*x509.Certificate, error)
	SetCertificate(keyring.KeySlot) error
}

type Signer interface {
	GetPrivateKey(keyring.KeySlot) (crypto.PrivateKey, error)
}

type Requester interface {
	GetPublicKey(keyring.KeySlot) (*x509.CertificateRequest, error)
}
