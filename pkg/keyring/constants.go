package keyring

type KeyType uint16

const (
	NilKey KeyType = iota
	EC256Key
	EC384Key
	Ed25519Key
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
