package cipher

import "golang.org/x/crypto/chacha20poly1305"

const (
	NonceSize = chacha20poly1305.NonceSizeX
	Overhead  = chacha20poly1305.Overhead
	KeySize   = chacha20poly1305.KeySize
)
