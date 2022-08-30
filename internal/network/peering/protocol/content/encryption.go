package content

import (
	"golang.org/x/crypto/chacha20poly1305"
)

func encryptContent(key []byte, nonce []byte, plaintextPayload []byte) ([]byte, error) {
	cipher, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, len(plaintextPayload)+cipher.Overhead())
	return cipher.Seal(ciphertext, nonce, plaintextPayload, nil), nil
}

func decryptContent(key []byte, nonce []byte, encryptedPayload []byte) ([]byte, error) {
	cipher, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, err
	}
	plaintext := make([]byte, len(encryptedPayload)-cipher.Overhead())
	return cipher.Open(plaintext, nonce, encryptedPayload, nil)
}
