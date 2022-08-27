package keyring

import (
	"crypto"
	"io"

	"github.com/jmg292/G-Net/pkg/gnet"
)

type PrivateKeyRing struct {
	keyring Backend
}

// Implement crypto.PrivateKey for private.KeyRing
func (keyring *PrivateKeyRing) Equal(key crypto.PrivateKey) bool {
	return false
}

// Implement crypto.PrivateKey for private.KeyRing
func (keyring *PrivateKeyRing) Public() crypto.PublicKey {
	return nil
}

// Implement crypto.Signer for private.KeyRing
func (keyring *PrivateKeyRing) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) (signature []byte, err error) {
	return nil, gnet.ErrorNotYetImplemented
}
