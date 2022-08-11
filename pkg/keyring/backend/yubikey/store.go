package yubikey

func (y *YubikeyStorageBackend) storeEncryptionKey(public x25519PublicBytes) (err error) {
	if handle, managementKey, e := y.getHandleAndManagementKey(); e != nil {
		err = e
	} else {
		defer y.releaseHandle()
		var keyCopy [24]byte
		copy(keyCopy[:], managementKey[:])
		key1, key2 := public.AsManagementKeys()
		if err = handle.SetManagementKey(*managementKey, key1); err == nil {
			if err = handle.SetManagementKey(key1, key2); err != nil {
				err = handle.SetManagementKey(key1, keyCopy)
			} else {
				err = handle.SetManagementKey(key2, keyCopy)
			}
		}
	}
	return
}
