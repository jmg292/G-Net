package signet

import (
	"crypto"
	"crypto/rand"

	"github.com/aead/ecdh"
	"golang.org/x/crypto/chacha20poly1305"
)

func generateSharedSecret(private crypto.PrivateKey, public crypto.PublicKey) ([]byte, error) {
	keyExchange := ecdh.X25519()
	if err := keyExchange.Check(public); err != nil {
		return nil, err
	}
	return keyExchange.ComputeSecret(private, public), nil
}

func Encrypt(data []byte, publicKey crypto.PublicKey) ([]byte, error) {
	var s *sealed
	var sharedSecret []byte
	keyExchange := ecdh.X25519()
	if ephemeralPrivateKey, ephemeralPublicKey, err := keyExchange.GenerateKey(rand.Reader); err != nil {
		return nil, err
	} else {
		if s, err = newSealedContent(data, ephemeralPublicKey); err != nil {
			return nil, err
		}
		if sharedSecret, err = generateSharedSecret(ephemeralPrivateKey, publicKey); err != nil {
			return nil, err
		}
	}
	if cipher, err := chacha20poly1305.NewX(sharedSecret); err != nil {
		cipher.Seal(s.content, s.nonce, data, nil)
	}
	return s.toBytes(), nil
}

func Decrypt(data []byte, privateKey crypto.PrivateKey) (plaintext []byte, err error) {
	if s, err := sealedFromBytes(data); err == nil {
		if sharedSecret, err := generateSharedSecret(privateKey, s.ephemeralPublicKey); err == nil {
			if cipher, err := chacha20poly1305.NewX(sharedSecret); err == nil {
				plaintext = make([]byte, len(s.content)-cipher.Overhead())
				if _, err = cipher.Open(plaintext, s.nonce, s.content, nil); err != nil {
					return nil, err
				}
			}
		}
	}
	return
}
