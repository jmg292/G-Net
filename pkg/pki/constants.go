package pki

type SupportedKeyType uint16

const (
	EC256Key SupportedKeyType = iota
	EC384Key
	X25519Key
)

type KeySlot uint8

const (
	SigningKeySlot KeySlot = iota
	EncryptionKeySlot
	AuthenticationKeySlot
)

type KeySizeBytes int

const (
	EC256KeySize  KeySizeBytes = 32
	EC384KeySize  KeySizeBytes = 48
	X25519KeySize KeySizeBytes = 32
)
