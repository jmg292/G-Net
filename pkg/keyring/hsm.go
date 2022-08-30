package keyring

type DeviceInfo interface {
	Vendor() string
	Identifier() []byte
	Version() string
}

type HardwareKeyRing interface {
	DeviceInfo
	Keystore
	PublicKeyRing
	PrivateKeyRing
}
