package key

import (
	"crypto/ecdsa"
	"crypto/elliptic"
)

func GenerateP384KeyPair() (*ecdsa.PublicKey, *ecdsa.PrivateKey, error) {
	return ecdsaGenerateKeyPair(elliptic.P384())
}

func SerializeP384PublicKey(key *ecdsa.PublicKey) []byte {
	return elliptic.MarshalCompressed(elliptic.P384(), key.X, key.Y)
}

func DeserializeP384PublicKey(keyBytes []byte) (*ecdsa.PublicKey, error) {
	return ecdsaDeserializePublicKey(elliptic.P384(), keyBytes)
}
