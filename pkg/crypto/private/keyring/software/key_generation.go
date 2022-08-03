package software

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"

	gcrypt "github.com/jmg292/G-Net/pkg/crypto"
	"github.com/jmg292/G-Net/pkg/gnet"
)

func (*SoftwareKeyRing) generateKey(keyType gcrypt.SupportedKeyType) (newKey crypto.PrivateKey, err error) {
	switch keyType {
	case gcrypt.EC256Key:
		newKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case gcrypt.EC384Key:
		newKey, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case gcrypt.X25519Key:
		_, newKey, err = ed25519.GenerateKey(rand.Reader)
	default:
		err = fmt.Errorf(string(gnet.ErrorUnsupportedAlgorithm))
	}
	return &newKey, err
}

func (keyring *SoftwareKeyRing) slotKey(key *crypto.PrivateKey, keyType gcrypt.SupportedKeyType, slot gcrypt.KeySlot) error {
	// ed25519 keys only supported for encryption
	if keyType != gcrypt.X25519Key && slot == gcrypt.EncryptionKeySlot {
		return fmt.Errorf(string(gnet.ErrorUnsupportedAlgorithmForKeySlot))
	} else if keyType == gcrypt.X25519Key && slot != gcrypt.EncryptionKeySlot {
		return fmt.Errorf(string(gnet.ErrorUnsupportedAlgorithmForKeySlot))
	}
	switch slot {
	case gcrypt.SigningKeySlot:
		if keyring.signingKey != nil {
			return fmt.Errorf(string(gnet.ErrorKeyAlreadyExists))
		}
		keyring.signingKey = key
		keyring.signingKeyType = keyType
	case gcrypt.EncryptionKeySlot:
		if keyring.encryptionKey != nil {
			return fmt.Errorf(string(gnet.ErrorKeyAlreadyExists))
		}
		keyring.encryptionKey = key
		keyring.encryptionKeyType = keyType
	case gcrypt.AuthenticationKeySlot:
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
