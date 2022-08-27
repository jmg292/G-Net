package keyring

import "golang.org/x/crypto/chacha20poly1305"

const (
	NonceSize     int = chacha20poly1305.NonceSizeX
	KdfIterations int = 5000
)

type SupportedKeyType uint16

const (
	NullKey SupportedKeyType = iota
	EC256Key
	EC384Key
	X25519Key
	ManagementKey
)

type KeySlot uint8

const (
	SigningKeySlot KeySlot = iota
	AuthenticationKeySlot
	EncryptionKeySlot
	DeviceKeySlot
	ManagementKeySlot
)

type KeySizeBytes int

const (
	EC256KeySize      KeySizeBytes = 32
	EC384KeySize      KeySizeBytes = 48
	X25519KeySize     KeySizeBytes = 32
	SymmetricKeySize  KeySizeBytes = 32
	ManagementKeySize KeySizeBytes = 24
)

type IdentityType uint8

const (
	CertificateAuthority IdentityType = 1 << iota
	Administrator
	User
	Device
)
