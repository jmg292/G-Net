package backend

import (
	"crypto"
	"crypto/x509"

	"github.com/jmg292/G-Net/pkg/keyring"
)

type Storage interface {
	// Identification
	Name() string

	// Management
	Open() error
	Unlock([]byte) error
	Lock() error
	Close() error

	// Key lifecycle
	CreateKey(keyring.SupportedKeyType, keyring.KeySlot) error

	// Key retrieval
	GetPrivateKey(keyring.KeySlot) (crypto.PrivateKey, error)
	GetPrivateBytes(keyring.KeySlot) ([]byte, error)
	GetPublicKey(keyring.KeySlot) (crypto.PublicKey, error)
	GetPublicBytes(keyring.KeySlot) (crypto.PublicKey, error)

	// Key storage
	PutPrivateKey(crypto.PrivateKey, keyring.KeySlot, bool) error
	PutPrivateBytes([]byte, keyring.KeySlot, bool) error
	PutPublicKey(crypto.PublicKey, keyring.KeySlot, bool) error
	PutPublicBytes([]byte, keyring.KeySlot, bool) error

	// Attestation functions
	AttestationCertificate() (*x509.Certificate, error)
	Attest(keyring.KeySlot) (*x509.Certificate, error)
}
