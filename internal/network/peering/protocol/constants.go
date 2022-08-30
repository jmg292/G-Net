package protocol

import "golang.org/x/crypto/chacha20poly1305"

const (
	KeySize        = chacha20poly1305.KeySize
	NonceSize      = chacha20poly1305.NonceSizeX
	DefaultTimeout = 300
)
