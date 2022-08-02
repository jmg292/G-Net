package software

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"

	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/identity"
)

func (*SoftwareKeyRing) generateKey(keyType identity.SupportedKeyType) (newKey crypto.PrivateKey, err error) {
	switch keyType {
	case identity.EC256Key:
		newKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case identity.EC384Key:
		newKey, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case identity.Ed25519Key:
		_, newKey, err = ed25519.GenerateKey(rand.Reader)
	default:
		err = fmt.Errorf(string(gnet.ErrorUnsupportedAlgorithm))
	}
	return &newKey, err
}

func (keyring *SoftwareKeyRing) slotKey(key *crypto.PrivateKey, keyType identity.SupportedKeyType, slot identity.KeySlot) error {
	// ed25519 keys only supported for encryption
	if keyType != identity.Ed25519Key && slot == identity.EncryptionKeySlot {
		return fmt.Errorf(string(gnet.ErrorUnsupportedAlgorithmForKeySlot))
	} else if keyType == identity.Ed25519Key && slot != identity.EncryptionKeySlot {
		return fmt.Errorf(string(gnet.ErrorUnsupportedAlgorithmForKeySlot))
	}
	switch slot {
	case identity.SigningKeySlot:
		if keyring.signingKey != nil {
			return fmt.Errorf(string(gnet.ErrorKeyAlreadyExists))
		}
		keyring.signingKey = key
		keyring.signingKeyType = keyType
	case identity.EncryptionKeySlot:
		if keyring.encryptionKey != nil {
			return fmt.Errorf(string(gnet.ErrorKeyAlreadyExists))
		}
		keyring.encryptionKey = key
		keyring.encryptionKeyType = keyType
	case identity.AuthenticationKeySlot:
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
