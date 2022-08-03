package public

import (
	"crypto"
	"crypto/x509"
)

type KeyRing interface {
	Name() string
	Fingerprint() []byte
	EnsureSecureKeyStorage() error
	EncryptionKey() crypto.PublicKey
	SigningCertificate() *x509.Certificate
	AuthenticationCertificate() *x509.Certificate
	Encrypt([]byte) ([]byte, error)
	VerifySignature([]byte, []byte) error
	VerifyAuthentication([]byte, []byte) error
}
