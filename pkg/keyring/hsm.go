package keyring

type DeviceInfo interface {
	DeviceVendor() string
	DeviceIdentifier() []byte
	DeviceVersion() string
}

type HardwareKeyRing interface {
	DeviceInfo
	Keystore
	PublicKeyRing
	Private
}
