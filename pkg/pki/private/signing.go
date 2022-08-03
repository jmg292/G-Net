package private

import (
	"crypto"
	"crypto/rand"
	"fmt"

	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/pki"
	"golang.org/x/crypto/sha3"
)

func (keyring *KeyRing) signDataWithKey(data []byte, keyType pki.KeySlot) ([]byte, error) {
	signingKey, err := keyring.storage.GetPrivateKey(keyType)
	if err != nil {
		return nil, err
	}
	if signer, ok := signingKey.(crypto.Signer); ok {
		digest := sha3.New256()
		return signer.Sign(rand.Reader, digest.Sum(data), nil)
	}
	return nil, fmt.Errorf(string(gnet.ErrorInvalidSigningKey))
}

func (keyring *KeyRing) Sign(data []byte) ([]byte, error) {
	return keyring.signDataWithKey(data, pki.SigningKeySlot)
}

func (keyring *KeyRing) Authenticate(data []byte) ([]byte, error) {
	return keyring.signDataWithKey(data, pki.AuthenticationKeySlot)
}
