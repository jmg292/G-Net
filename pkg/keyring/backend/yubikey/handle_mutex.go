package yubikey

import "github.com/go-piv/piv-go/piv"

func (y *YubikeyStorageBackend) getHandle() (handle *piv.YubiKey, err error) {
	if err = y.assertOpen(); err == nil {
		y.mutex.Lock()
		handle = y.handle
	}
	return
}

func (y *YubikeyStorageBackend) releaseHandle() {
	y.mutex.Unlock()
}

func (y *YubikeyStorageBackend) getHandleAndManagementKey() (handle *piv.YubiKey, managementKey *[24]byte, err error) {
	if err = y.assertOpenAndUnlocked(); err == nil {
		if handle, err = y.getHandle(); err != nil {
			handle = nil
			y.releaseHandle()
		} else {
			managementKey = y.metadata.ManagementKey
		}
	}
	return
}
