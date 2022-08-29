package keyring

import "errors"

var (
	ErrInvalidSigningKey    = errors.New("invalid signing key")
	ErrInvalidDecryptionKey = errors.New("invalid decryption key")
)
