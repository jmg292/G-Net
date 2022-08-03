package software

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"

	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/pki"
)

func (*SoftwareKeyRing) generateKey(keyType pki.SupportedKeyType) (newKey crypto.PrivateKey, err error) {
	switch keyType {
	case pki.EC256Key:
		newKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case pki.EC384Key:
		newKey, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case pki.Ed25519Key:
		_, newKey, err = ed25519.GenerateKey(rand.Reader)
	default:
		err = fmt.Errorf(string(gnet.ErrorUnsupportedAlgorithm))
	}
	return &newKey, err
}

func (keyring *SoftwareKeyRing) slotKey(key *crypto.PrivateKey, keyType pki.SupportedKeyType, slot pki.KeySlot) error {
	// ed25519 keys only supported for encryption
	if keyType != pki.Ed25519Key && slot == pki.EncryptionKeySlot {
		return fmt.Errorf(string(gnet.ErrorUnsupportedAlgorithmForKeySlot))
	} else if keyType == pki.Ed25519Key && slot != pki.EncryptionKeySlot {
		return fmt.Errorf(string(gnet.ErrorUnsupportedAlgorithmForKeySlot))
	}
	switch slot {
	case pki.SigningKeySlot:
		if keyring.signingKey != nil {
			return fmt.Errorf(string(gnet.ErrorKeyAlreadyExists))
		}
		keyring.signingKey = key
		keyring.signingKeyType = keyType
	case pki.EncryptionKeySlot:
		if keyring.encryptionKey != nil {
			return fmt.Errorf(string(gnet.ErrorKeyAlreadyExists))
		}
		keyring.encryptionKey = key
		keyring.encryptionKeyType = keyType
	case pki.AuthenticationKeySlot:
		if keyring.encryptionKey != nil {
			return fmt.Errorf(string(gnet.ErrorKeyAlreadyExists))
		}
		keyring.authenticationKey = key
		keyring.authenticationKeyType = keyType
	default:
		return fmt.Errorf(string(gnet.ErrorInvalidKeySlot))
	}
	return nil
}
