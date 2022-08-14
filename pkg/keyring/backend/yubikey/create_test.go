package yubikey_test

import (
	"github.com/jmg292/G-Net/pkg/keyring"
)

type KeyGenTestParams struct {
	Slot          keyring.KeySlot
	Type          keyring.SupportedKeyType
	ExpectSuccess bool
}

var keyGenTestParams []*KeyGenTestParams = []*KeyGenTestParams{
	&KeyGenTestParams{Slot: keyring.SigningKeySlot, Type: keyring.NullKey, ExpectSuccess: false},
	&KeyGenTestParams{Slot: keyring.SigningKeySlot, Type: keyring.EC256Key, ExpectSuccess: true},
	&KeyGenTestParams{Slot: keyring.SigningKeySlot, Type: keyring.EC384Key, ExpectSuccess: true},
	&KeyGenTestParams{Slot: keyring.SigningKeySlot, Type: keyring.X25519Key, ExpectSuccess: false},
	&KeyGenTestParams{Slot: keyring.SigningKeySlot, Type: keyring.ManagementKey, ExpectSuccess: false},
	&KeyGenTestParams{Slot: keyring.AuthenticationKeySlot, Type: keyring.NullKey, ExpectSuccess: false},
	&KeyGenTestParams{Slot: keyring.AuthenticationKeySlot, Type: keyring.EC256Key, ExpectSuccess: true},
	&KeyGenTestParams{Slot: keyring.AuthenticationKeySlot, Type: keyring.EC384Key, ExpectSuccess: true},
	&KeyGenTestParams{Slot: keyring.AuthenticationKeySlot, Type: keyring.X25519Key, ExpectSuccess: false},
	&KeyGenTestParams{Slot: keyring.AuthenticationKeySlot, Type: keyring.ManagementKey, ExpectSuccess: false},
	&KeyGenTestParams{Slot: keyring.EncryptionKeySlot, Type: keyring.NullKey, ExpectSuccess: false},
	&KeyGenTestParams{Slot: keyring.EncryptionKeySlot, Type: keyring.EC256Key, ExpectSuccess: true},
	&KeyGenTestParams{Slot: keyring.EncryptionKeySlot, Type: keyring.EC384Key, ExpectSuccess: true},
	&KeyGenTestParams{Slot: keyring.EncryptionKeySlot, Type: keyring.X25519Key, ExpectSuccess: false},
	&KeyGenTestParams{Slot: keyring.EncryptionKeySlot, Type: keyring.ManagementKey, ExpectSuccess: false},
	&KeyGenTestParams{Slot: keyring.DeviceKeySlot, Type: keyring.NullKey, ExpectSuccess: false},
	&KeyGenTestParams{Slot: keyring.DeviceKeySlot, Type: keyring.EC256Key, ExpectSuccess: true},
	&KeyGenTestParams{Slot: keyring.DeviceKeySlot, Type: keyring.EC384Key, ExpectSuccess: true},
	&KeyGenTestParams{Slot: keyring.DeviceKeySlot, Type: keyring.X25519Key, ExpectSuccess: false},
	&KeyGenTestParams{Slot: keyring.DeviceKeySlot, Type: keyring.ManagementKey, ExpectSuccess: false},
	&KeyGenTestParams{Slot: keyring.ManagementKeySlot, Type: keyring.NullKey, ExpectSuccess: false},
	&KeyGenTestParams{Slot: keyring.ManagementKeySlot, Type: keyring.EC256Key, ExpectSuccess: false},
	&KeyGenTestParams{Slot: keyring.ManagementKeySlot, Type: keyring.EC384Key, ExpectSuccess: false},
	&KeyGenTestParams{Slot: keyring.ManagementKeySlot, Type: keyring.X25519Key, ExpectSuccess: false},
	&KeyGenTestParams{Slot: keyring.ManagementKeySlot, Type: keyring.ManagementKey, ExpectSuccess: true},
}
