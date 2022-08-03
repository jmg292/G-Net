package keystore

import (
	"crypto"
	"crypto/x509"

	gcrypt "github.com/jmg292/G-Net/pkg/crypto"
	"github.com/jmg292/G-Net/pkg/crypto/public"
)

type Storage interface {
	// Identification
	Name() string
	GetKeyId(gcrypt.KeySlot) []byte

	// Management
	Unlock([]byte) error
	Validate() error
	ManagementKey() ([]byte, error)
	KeyEncryptionKey() ([]byte, error)
	Lock() error

	// Key lifecycle
	DestroyKey(gcrypt.KeySlot) error
	CreateKey(gcrypt.SupportedKeyType) ([]byte, error)

	// Public keys
	GetPublicKey(gcrypt.KeySlot) (crypto.PublicKey, error)
	PublicKeyRing() (public.KeyRing, error)

	// Key retrieval
	GetPrivateKey(gcrypt.KeySlot) (crypto.PrivateKey, error)
	GetPrivateKeyBytes(gcrypt.KeySlot) ([]byte, error)

	// Key storage
	PutPrivateKey(crypto.PrivateKey, gcrypt.KeySlot, bool) error
	PutPrivateKeyBytes([]byte, gcrypt.KeySlot, bool) error

	// Attestation functions
	AttestationCertificate() (*x509.Certificate, error)
	Attest(gcrypt.KeySlot) (*x509.Certificate, error)
}
