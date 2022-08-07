package private

import (
	"crypto"
	"fmt"
	"io"

	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/keyring/backend"
)

type KeyRing struct {
	keystore backend.Storage
}

// Implement crypto.PrivateKey for private.KeyRing
func (keyring *KeyRing) Equal(key crypto.PrivateKey) bool {
	return false
}

// Implement crypto.PrivateKey for private.KeyRing
func (keyring *KeyRing) Public() crypto.PublicKey {
	return nil
}

// Implement crypto.Signer for private.KeyRing
func (keyring *KeyRing) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) (signature []byte, err error) {
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}
