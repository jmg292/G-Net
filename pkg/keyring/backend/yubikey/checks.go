package yubikey

import (
	"crypto/subtle"

	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/pkg/keyring"
)

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
	if _, err = y.Attest(slot); err != nil && err == piv.ErrNotFound {
		empty = true
	}
	return
}
