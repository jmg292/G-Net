package yubikey

import (
	"github.com/go-piv/piv-go/piv"
)

func (y *YubikeyStorageBackend) storeEncryptionKey(public x25519PublicBytes) (err error) {
	keyPolicy := piv.Key{
		PINPolicy:   piv.PINPolicyAlways,
		TouchPolicy: piv.TouchPolicyAlways,
	}
	if handle, managementKey, e := y.getHandleAndManagementKey(); e != nil {
		err = e
	} else {
		defer y.releaseHandle()
		if slot1, slot2, e := y.getX25519KeySlots(); e != nil {
			err = e
		} else if err = handle.SetPrivateKeyInsecure(*managementKey, slot1, public[:24], keyPolicy); err == nil {
			err = handle.SetPrivateKeyInsecure(*managementKey, slot2, public[24:], keyPolicy)
		}
	}
	return
}
