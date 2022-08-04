package file

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"io"

	"github.com/cloudflare/circl/dh/x25519"
	gcrypt "github.com/jmg292/G-Net/pkg/crypto"
	"github.com/jmg292/G-Net/pkg/gnet"
)

func (*fileKeyStore) generateKey(keyType gcrypt.SupportedKeyType) (newKey crypto.PrivateKey, err error) {
	switch keyType {
	case gcrypt.EC256Key:
		newKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case gcrypt.EC384Key:
		newKey, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case gcrypt.X25519Key:
		var privateKey x25519.Key
		io.ReadFull(rand.Reader, privateKey[:])
		newKey = privateKey
	default:
		err = fmt.Errorf(string(gnet.ErrorUnsupportedAlgorithm))
	}
	return &newKey, err
}

func (keyring *fileKeyStore) slotKey(key *crypto.PrivateKey, keyType gcrypt.SupportedKeyType, slot gcrypt.KeySlot) error {
	// X25519 keys only supported for encryption
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
		var emptyKey x25519.Key
		if bytes.Compare(keyring.encryptionKey[:], emptyKey[:]) != 0 {
			return fmt.Errorf(string(gnet.ErrorKeyAlreadyExists))
		}
		keyring.encryptionKey = key
		keyring.encryptionKeyType = keyType
	case gcrypt.AuthenticationKeySlot:
		if keyring.authenticationKey != nil {
			return fmt.Errorf(string(gnet.ErrorKeyAlreadyExists))
		}
		keyring.authenticationKey = key
		keyring.authenticationKeyType = keyType
	default:
		return fmt.Errorf(string(gnet.ErrorInvalidKeySlot))
	}
	return nil
}
