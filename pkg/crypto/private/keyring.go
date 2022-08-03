package private

import (
	"github.com/jmg292/G-Net/pkg/crypto/private/keyring"
)

type KeyRing struct {
	unlockKey []byte
	storage   keyring.KeyStorage
}

func (keyring *KeyRing) Open() error {
	if err := keyring.storage.Unlock(keyring.unlockKey); err != nil {
		return err
	}
	if err := keyring.storage.Validate(); err != nil {
		return err
	}
	return nil
}
