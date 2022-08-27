package key

import (
	"crypto"

	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/keyring"
)

func GenerateKeyPair(keyType keyring.SupportedKeyType) (public crypto.PublicKey, private crypto.PrivateKey, err error) {
	switch keyType {
	case keyring.EC256Key:
		public, private, err = GenerateP256KeyPair()
	case keyring.EC384Key:
		public, private, err = GenerateP384KeyPair()
	case keyring.X25519Key:
		public, private = GenerateX25519KeyPair()
	default:
		err = gnet.ErrorUnsupportedAlgorithm
	}
	return
}

func PrivateToBytes(keyType keyring.SupportedKeyType, key crypto.PrivateKey) (keyBytes []byte, err error) {
	switch keyType {
	case keyring.EC256Key:
		keyBytes, err = ecdsaPrivateToBytes(key)
	case keyring.EC384Key:
		keyBytes, err = ecdsaPrivateToBytes(key)
	case keyring.X25519Key:
		if x25519key, ok := key.(*X25519PrivateKey); !ok {
			err = gnet.ErrorInvalidPrivateKey
		} else {
			keyBytes = x25519key.Bytes()
		}
	default:
		err = gnet.ErrorUnsupportedAlgorithm
	}
	return
}

func BytesToPrivate(keyType keyring.SupportedKeyType, keyBytes []byte) (key crypto.PrivateKey, err error) {
	switch keyType {
	case keyring.EC256Key:
		key, err = DeserializeEcdsaKey(keyBytes)
	case keyring.EC384Key:
		key, err = DeserializeEcdsaKey(keyBytes)
	case keyring.X25519Key:
		key, err = DeserializeX25519Key(keyBytes)
	default:
		err = gnet.ErrorUnsupportedAlgorithm
	}
	return
}

func PublicToBytes(keyType keyring.SupportedKeyType, key crypto.PublicKey) (keyBytes []byte, err error) {
	switch keyType {
	case keyring.EC256Key:
		keyBytes, err = ecdsaPublicToBytes(key)
	case keyring.EC384Key:
		keyBytes, err = ecdsaPublicToBytes(key)
	case keyring.X25519Key:
		if x25519key, ok := key.(*X25519PublicKey); !ok {
			err = gnet.ErrorInvalidPublicKey
		} else {
			keyBytes = x25519key.Bytes()
		}
	default:
		err = gnet.ErrorUnsupportedAlgorithm
	}
	return
}

func BytesToPublic(keyType keyring.SupportedKeyType, keyBytes []byte) (key crypto.PublicKey, err error) {
	switch keyType {
	case keyring.EC256Key:
		key, err = DeserializeP256PublicKey(keyBytes)
	case keyring.EC384Key:
		key, err = DeserializeP384PublicKey(keyBytes)
	case keyring.X25519Key:
		key, err = DeserializeX25519PublicKey(keyBytes)
	default:
		err = gnet.ErrorUnsupportedAlgorithm
	}
	return
}
