package identity

type SupportedKeyType uint16

const (
	EC256Key SupportedKeyType = iota
	EC384Key
	Ed25519Key
)

type KeySlot uint8

const (
	SigningKeySlot KeySlot = iota
	EncryptionKeySlot
	AuthenticationKeySlot
)
