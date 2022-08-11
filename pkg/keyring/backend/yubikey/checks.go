package yubikey

import (
	"crypto/subtle"

	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/keyring"
)

func (y *YubikeyStorageBackend) assertOpen() (err error) {
	if y.handle == nil {
		err = gnet.ErrorInvalidHandle
	}
	return
}

func (y *YubikeyStorageBackend) assertOpenAndUnlocked() (err error) {
	if err = y.assertOpen(); err == nil {
		if y.metadata == nil {
			err = gnet.ErrorKeystoreLocked
		}
	}
	return
}

func (y *YubikeyStorageBackend) assertDefaultManagementKey() (defaultKey bool, err error) {
	if _, managementKey, e := y.getHandleAndManagementKey(); e != nil {
		err = e
	} else {
		defer y.releaseHandle()
		defaultKey = subtle.ConstantTimeCompare(managementKey[:], piv.DefaultManagementKey[:]) == 1
	}
	return
}

func (y *YubikeyStorageBackend) assertKeySlotIsEmpty(slot keyring.KeySlot) (empty bool, err error) {
	if slot != keyring.EncryptionKeySlot {
		_, err = y.Attest(slot)
	} else {
		_, err = y.getX25519PrivateKey()
	}
	if err != nil && (err == piv.ErrNotFound || err == gnet.ErrorKeyNotFound) {
		err = nil
		empty = true
	}
	return
}
