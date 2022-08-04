package file

import (
	"crypto/rand"

	gcrypt "github.com/jmg292/G-Net/pkg/crypto"
	"github.com/jmg292/G-Net/pkg/crypto/kdf"
	"golang.org/x/crypto/chacha20poly1305"
)

// File-based keystores won't have access to X25519 private keys for symmetric key generation
// Instead, they use a symmetric key derived from a user-provided PIN
func encrypt(data []byte, pin []byte) (ciphertext []byte, err error) {
	nonce := make([]byte, gcrypt.NonceSize)
	rand.Read(nonce)
	if cipher, err := chacha20poly1305.NewX(kdf.DeriveKey(pin, nonce)); err == nil {
		ciphertext = make([]byte, len(data)+len(nonce)+cipher.Overhead())
		ciphertext = append(ciphertext, nonce...)
		cipher.Seal(ciphertext[len(nonce):], nonce, data, nil)
	}
	return
}

func decrypt(data []byte, pin []byte) (plaintext []byte, err error) {
	if cipher, e := chacha20poly1305.NewX(kdf.DeriveKey(pin, data[:gcrypt.NonceSize])); e == nil {
		plaintext = make([]byte, len(data)-cipher.Overhead()-gcrypt.NonceSize)
		plaintext, err = cipher.Open(plaintext, data[:gcrypt.NonceSize], data[gcrypt.NonceSize:], nil)
	} else {
		err = e
	}
	return
}
