package yubikey

import (
	"crypto/rand"
	"crypto/subtle"
	"io"

	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/pkg/gnet"
)

func (y *YubikeyStorageBackend) generateManagementKey() (err error) {
	if usingDefaultKey, e := y.assertDefaultManagementKey(); e != nil {
		err = e
	} else if !usingDefaultKey {
		err = gnet.ErrorKeyAlreadyExists
	} else if handle, managementKey, e := y.getHandleAndManagementKey(); e != nil {
		err = e
	} else {
		defer y.releaseHandle()
		var newKey [24]byte
		if _, err = io.ReadFull(rand.Reader, newKey[:]); err == nil {
			err = handle.SetManagementKey(*managementKey, newKey)
		}
	}
	return
}

func (y *YubikeyStorageBackend) generateKey(slot piv.Slot, alg piv.Algorithm) (err error) {
	keyTemplate := piv.Key{
		Algorithm:   alg,
		PINPolicy:   piv.PINPolicyAlways,
		TouchPolicy: piv.TouchPolicyAlways,
	}
	if handle, managementKey, e := y.getHandleAndManagementKey(); e != nil {
		err = e
	} else {
		defer y.releaseHandle()
		_, err = handle.GenerateKey(*managementKey, slot, keyTemplate)
	}
	return
}

func (y *YubikeyStorageBackend) generateEncryptionKey() (publicBytes x25519PublicBytes, err error) {
	publicBytes.GenerateSalt()
	if privateKey, e := y.deriveX25519PrivateKey(publicBytes.Salt()); e != nil {
		err = e
	} else {
		subtle.ConstantTimeCopy(1, publicBytes.Key(), privateKey.PublicBytes())
	}
	return
}
