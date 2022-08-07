package public

import (
	"crypto"

	"github.com/jmg292/G-Net/pkg/keyring/certificates"
	"github.com/jmg292/G-Net/pkg/keyring/key"
)

type KeyRing struct {
	encryptionKey *key.X25519PublicKey
	certstore     certificates.CertificateStore
}

// Implement crypto.PublicKey for public.KeyRing
func (keyring *KeyRing) Equal(key crypto.PublicKey) bool {
	return false
}
