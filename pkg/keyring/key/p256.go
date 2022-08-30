package key

import (
	"crypto/ecdsa"
	"crypto/elliptic"
)

func GenerateP256KeyPair() (*ecdsa.PublicKey, *ecdsa.PrivateKey, error) {
	return ecdsaGenerateKeyPair(elliptic.P256())
}

func SerializeP256PublicKey(key *ecdsa.PublicKey) []byte {
	return elliptic.MarshalCompressed(elliptic.P256(), key.X, key.Y)
}

func DeserializeP256PublicKey(keyBytes []byte) (*ecdsa.PublicKey, error) {
	return ecdsaDeserializePublicKey(elliptic.P256(), keyBytes)
}
