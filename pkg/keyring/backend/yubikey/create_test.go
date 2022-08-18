package yubikey_test

import (
	"github.com/jmg292/G-Net/pkg/keyring"
	"github.com/jmg292/G-Net/pkg/keyring/backend/yubikey"
)

type KeyGenTestParams struct {
	Slot          keyring.KeySlot
	Type          keyring.SupportedKeyType
	ExpectSuccess bool
}

var keyGenTestParams []*KeyGenTestParams = []*KeyGenTestParams{
	{Slot: keyring.SigningKeySlot, Type: keyring.NullKey, ExpectSuccess: false},
	{Slot: keyring.SigningKeySlot, Type: keyring.EC256Key, ExpectSuccess: true},
	{Slot: keyring.SigningKeySlot, Type: keyring.EC384Key, ExpectSuccess: true},
	{Slot: keyring.SigningKeySlot, Type: keyring.X25519Key, ExpectSuccess: false},
	{Slot: keyring.SigningKeySlot, Type: keyring.ManagementKey, ExpectSuccess: false},
	{Slot: keyring.AuthenticationKeySlot, Type: keyring.NullKey, ExpectSuccess: false},
	{Slot: keyring.AuthenticationKeySlot, Type: keyring.EC256Key, ExpectSuccess: true},
	{Slot: keyring.AuthenticationKeySlot, Type: keyring.EC384Key, ExpectSuccess: true},
	{Slot: keyring.AuthenticationKeySlot, Type: keyring.X25519Key, ExpectSuccess: false},
	{Slot: keyring.AuthenticationKeySlot, Type: keyring.ManagementKey, ExpectSuccess: false},
	{Slot: keyring.EncryptionKeySlot, Type: keyring.NullKey, ExpectSuccess: false},
	{Slot: keyring.EncryptionKeySlot, Type: keyring.EC256Key, ExpectSuccess: true},
	{Slot: keyring.EncryptionKeySlot, Type: keyring.EC384Key, ExpectSuccess: true},
	{Slot: keyring.EncryptionKeySlot, Type: keyring.X25519Key, ExpectSuccess: false},
	{Slot: keyring.EncryptionKeySlot, Type: keyring.ManagementKey, ExpectSuccess: false},
	{Slot: keyring.DeviceKeySlot, Type: keyring.NullKey, ExpectSuccess: false},
	{Slot: keyring.DeviceKeySlot, Type: keyring.EC256Key, ExpectSuccess: true},
	{Slot: keyring.DeviceKeySlot, Type: keyring.EC384Key, ExpectSuccess: true},
	{Slot: keyring.DeviceKeySlot, Type: keyring.X25519Key, ExpectSuccess: false},
	{Slot: keyring.DeviceKeySlot, Type: keyring.ManagementKey, ExpectSuccess: false},
	{Slot: keyring.ManagementKeySlot, Type: keyring.NullKey, ExpectSuccess: false},
	{Slot: keyring.ManagementKeySlot, Type: keyring.EC256Key, ExpectSuccess: false},
	{Slot: keyring.ManagementKeySlot, Type: keyring.EC384Key, ExpectSuccess: false},
	{Slot: keyring.ManagementKeySlot, Type: keyring.X25519Key, ExpectSuccess: false},
	{Slot: keyring.ManagementKeySlot, Type: keyring.ManagementKey, ExpectSuccess: true},
}

func generatePrivateKeys(yk *yubikey.Yubikey) (err error) {
	if err = yk.Reset(); err == nil {
		if err = yk.CreateKey(keyring.ManagementKeySlot, keyring.ManagementKey); err == nil {
			for i := keyring.SigningKeySlot; i < keyring.ManagementKeySlot; i++ {
				if err = yk.CreateKey(i, keyring.EC384Key); err != nil {
					break
				}
			}
		}
	}
	return
}
