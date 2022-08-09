package yubikey_test

import (
	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/pkg/keyring"
)

var invalidKeyGeneration map[keyring.KeySlot][]keyring.SupportedKeyType = map[keyring.KeySlot][]keyring.SupportedKeyType{
	keyring.SigningKeySlot:        {keyring.NullKey, keyring.X25519Key, keyring.ManagementKey},
	keyring.AuthenticationKeySlot: {keyring.NullKey, keyring.X25519Key, keyring.ManagementKey},
	keyring.EncryptionKeySlot:     {keyring.NullKey, keyring.EC256Key, keyring.EC384Key, keyring.ManagementKey},
	keyring.DeviceKeySlot:         {keyring.NullKey, keyring.X25519Key, keyring.ManagementKey},
	keyring.ManagementKeySlot:     {keyring.NullKey, keyring.EC256Key, keyring.EC384Key, keyring.X25519Key},
}

var validKeyGeneration map[keyring.KeySlot][]keyring.SupportedKeyType = map[keyring.KeySlot][]keyring.SupportedKeyType{
	keyring.SigningKeySlot:        {keyring.EC256Key, keyring.EC384Key},
	keyring.AuthenticationKeySlot: {keyring.EC256Key, keyring.EC384Key},
	keyring.EncryptionKeySlot:     {keyring.X25519Key},
	keyring.DeviceKeySlot:         {keyring.EC256Key, keyring.EC384Key},
	keyring.ManagementKeySlot:     {keyring.ManagementKey},
}

func (r *ResetAuthorized) batchGenerateKey(keyslotmap map[keyring.KeySlot][]keyring.SupportedKeyType, failOnError bool) {
	for slot, keys := range keyslotmap {
		for _, key := range keys {
			r.reset()
			r.Test.Logf("Generating key. (KeySlot: %d, KeyType: %d)", slot, key)
			if err := r.yk.CreateKey(key, slot); err != nil && failOnError {
				r.Test.Errorf("--- Failed.  Error: %s", err)
			} else if err == nil && !failOnError {
				r.Test.Errorf("--- Failed.  Key generation succeeded.")
			} else if err != nil && !failOnError {
				r.Test.Logf("--- Passed.  Error: %s", err)
			} else {
				r.Test.Logf("--- Passed.")
			}
		}
	}
}

func (r *ResetAuthorized) TestInvalidKeyGeneration() {
	r.openAndUnlock(piv.DefaultPIN)
	r.Test.Logf("Batch testing invalid key slot/key type generation combinations")
	r.batchGenerateKey(invalidKeyGeneration, false)
	r.yk.Close()
}

func (r *ResetAuthorized) TestValidKeyGeneration() {
	r.openAndUnlock(piv.DefaultPIN)
	r.Test.Logf("Batch testing valid key slot/key type generation combinations")
	r.batchGenerateKey(validKeyGeneration, true)
	r.yk.Close()
}

func (r *ResetAuthorized) TestKeyAlreadyExists() {
	r.openAndUnlock(piv.DefaultPIN)
	r.Test.Logf("Testing key generation when key already exists")
	for keyslot, validtypes := range validKeyGeneration {
		r.reset()
		r.Test.Logf("Generating new key. (KeySlot: %d, KeyType: %d)", keyslot, validtypes[0])
		if err := r.yk.CreateKey(validtypes[0], keyslot); err != nil {
			r.Test.Errorf("Failed to generate key. (KeySlot: %d, KeyType: %d, Error: %s)", keyslot, validtypes[0], err)
		} else if err := r.yk.CreateKey(validtypes[0], keyslot); err == nil {
			r.Test.Errorf("Failed to stop key generation. (KeySlot: %d, KeyType: %d)", keyslot, validtypes[0])
		} else {
			r.Test.Logf("--- Passed. (KeySlot: %d, KeyType: %d)", keyslot, validtypes[0])
		}
	}
	r.yk.Close()
}
