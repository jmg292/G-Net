package keyring

import (
	"crypto"
	"crypto/x509"
)

type Backend interface {
	// Identification
	KeyInfo() (string, error)

	// Management
	Open() error
	Unlock([]byte) error
	Lock() error
	Close() error

	// Key lifecycle
	CreateKey(KeySlot, SupportedKeyType) error

	// Key retrieval
	GetPrivateKey(KeySlot) (crypto.PrivateKey, error)
	GetPublicKey(KeySlot) (crypto.PublicKey, error)

	// Certificate Management
	GetCertificate(KeySlot) (*x509.Certificate, error)
	SetCertificate(KeySlot, *x509.Certificate) error

	// Attestation functions
	AttestationCertificate() (*x509.Certificate, error)
	Attest(KeySlot) (*x509.Certificate, error)
}
