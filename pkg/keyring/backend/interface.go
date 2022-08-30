package backend

import (
	"crypto"
	"crypto/x509"

	"github.com/jmg292/G-Net/pkg/keyring"
)

type Storage interface {
	// Identification
	Name() (string, error)

	// Management
	Open() error
	Unlock([]byte) error
	Lock() error
	Close() error

	// Key lifecycle
	CreateKey(keyring.KeySlot, keyring.SupportedKeyType) error

	// Key retrieval
	GetPrivateKey(keyring.KeySlot) (crypto.PrivateKey, error)
	GetPublicKey(keyring.KeySlot) (crypto.PublicKey, error)

	// Certificate Management
	GetCertificate(keyring.KeySlot) (*x509.Certificate, error)
	SetCertificate(keyring.KeySlot, *x509.Certificate) error

	// Attestation functions
	AttestationCertificate() (*x509.Certificate, error)
	Attest(keyring.KeySlot) (*x509.Certificate, error)
}
