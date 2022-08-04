package crypto

import "golang.org/x/crypto/chacha20poly1305"

const (
	NonceSize     int = chacha20poly1305.NonceSizeX
	KdfIterations int = 5000
)

type SupportedKeyType uint16

const (
	EC256Key SupportedKeyType = iota
	EC384Key
	X25519Key
)

type KeySlot uint8

const (
	SigningKeySlot KeySlot = iota
	AuthenticationKeySlot
	EncryptionKeySlot
	DeviceKeySlot
)

type KeySizeBytes int

const (
	EC256KeySize     KeySizeBytes = 32
	EC384KeySize     KeySizeBytes = 48
	X25519KeySize    KeySizeBytes = 32
	SymmetricKeySize KeySizeBytes = 32
)
