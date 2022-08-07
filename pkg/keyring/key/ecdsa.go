package key

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/jmg292/G-Net/internal/utilities/convert"
	"github.com/jmg292/G-Net/pkg/gnet"
)

func SerializeEcdsaKey(key *ecdsa.PrivateKey) []byte {
	serializedKey := append(convert.UInt16ToBytes(uint16(key.Params().BitSize/8)), key.D.Bytes()...)
	return append(serializedKey, elliptic.MarshalCompressed(key.Curve, key.X, key.Y)...)
}

func DeserializeEcdsaKey(keyBytes []byte) (key *ecdsa.PrivateKey, err error) {
	var publicKey *ecdsa.PublicKey
	keySize := convert.BytesToUInt16(keyBytes[:2])
	privateBytes := keyBytes[2 : keySize+2]
	switch keySize {
	case 32:
		publicKey, err = DeserializeP256PublicKey(keyBytes[keySize+2:])
	case 48:
		publicKey, err = DeserializeP384PublicKey(keyBytes[keySize+2:])
	default:
		err = gnet.ErrorUnsupportedAlgorithm
	}
	if err == nil {
		key = ecdsaDeserializePrivateKey(privateBytes, publicKey)
	}
	return
}

func ecdsaGenerateKeyPair(curve elliptic.Curve) (*ecdsa.PublicKey, *ecdsa.PrivateKey, error) {
	var private *ecdsa.PrivateKey
	if key, err := ecdsa.GenerateKey(curve, rand.Reader); err != nil {
		return nil, nil, err
	} else {
		private = key
	}
	return private.Public().(*ecdsa.PublicKey), private, nil
}

func ecdsaDeserializePublicKey(curve elliptic.Curve, keyBytes []byte) (*ecdsa.PublicKey, error) {
	publicKey := ecdsa.PublicKey{
		Curve: curve,
	}
	publicKey.X, publicKey.Y = elliptic.UnmarshalCompressed(curve, keyBytes)
	if publicKey.X == nil {
		return nil, gnet.ErrorInvalidPublicKey
	}
	return &publicKey, nil
}

func ecdsaDeserializePrivateKey(keyBytes []byte, public *ecdsa.PublicKey) *ecdsa.PrivateKey {
	key := ecdsa.PrivateKey{
		PublicKey: *public,
	}
	key.D.SetBytes(keyBytes)
	return &key
}

func ecdsaPrivateToBytes(key crypto.PrivateKey) (keyBytes []byte, err error) {
	if ecdsaKey, ok := key.(*ecdsa.PrivateKey); !ok {
		err = gnet.ErrorInvalidPrivateKey
	} else {
		keyBytes = SerializeEcdsaKey(ecdsaKey)
	}
	return
}

func ecdsaPublicToBytes(key crypto.PublicKey) (keyBytes []byte, err error) {
	if ecdsaKey, ok := key.(*ecdsa.PublicKey); !ok {
		err = gnet.ErrorInvalidPublicKey
	} else {
		switch ecdsaKey.Params().Name {
		case "P-256":
			keyBytes = SerializeP256PublicKey(ecdsaKey)
		case "P-384":
			keyBytes = SerializeP384PublicKey(ecdsaKey)
		default:
			err = gnet.ErrorUnsupportedAlgorithm
		}
	}
	return
}
