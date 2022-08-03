package keyring

import (
	"crypto"
	"crypto/x509"

	"github.com/jmg292/G-Net/pkg/pki"
	"github.com/jmg292/G-Net/pkg/pki/public"
)

type KeyStorage interface {
	// Identification
	Name() string
	GetKeyId(pki.KeySlot) []byte

	// Management
	Unlock([]byte) error
	Validate() error
	ManagementKey() ([]byte, error)
	KeyEncryptionKey() ([]byte, error)
	Lock() error

	// Key lifecycle
	DestroyKey(pki.KeySlot) error
	CreateKey(pki.SupportedKeyType) ([]byte, error)

	// Public keys
	GetPublicKey(pki.KeySlot) (crypto.PublicKey, error)
	PublicKeyRing() (public.KeyRing, error)

	// Key retrieval
	GetPrivateKey(pki.KeySlot) (crypto.PrivateKey, error)
	GetPrivateKeyBytes(pki.KeySlot) ([]byte, error)

	// Key storage
	PutPrivateKey(crypto.PrivateKey, pki.KeySlot, bool) error
	PutPrivateKeyBytes([]byte, pki.KeySlot, bool) error

	// Attestation functions
	AttestationCertificate() (*x509.Certificate, error)
	Attest(pki.KeySlot) (*x509.Certificate, error)

	// Cryptographic functions
	Sign([]byte) ([]byte, error)
	Authenticate([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
}
