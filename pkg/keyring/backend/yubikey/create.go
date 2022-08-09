package yubikey

import (
	"crypto/subtle"

	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/pkg/gnet"
)

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

func (y *YubikeyStorageBackend) storeEncryptionKey(public x25519PublicBytes) (err error) {
	keyPolicy := piv.Key{
		PINPolicy:   piv.PINPolicyAlways,
		TouchPolicy: piv.TouchPolicyAlways,
	}
	if handle, managementKey, e := y.getHandleAndManagementKey(); e != nil {
		err = e
	} else {
		defer y.releaseHandle()
		if slot1, ok := piv.RetiredKeyManagementSlot(0x95); !ok {
			err = gnet.ErrorInvalidKeySlot
		} else if slot2, ok := piv.RetiredKeyManagementSlot(0x94); !ok {
			err = gnet.ErrorInvalidKeySlot
		} else if err = handle.SetPrivateKeyInsecure(*managementKey, slot1, public[:24], keyPolicy); err == nil {
			err = handle.SetPrivateKeyInsecure(*managementKey, slot2, public[24:], keyPolicy)
		}
	}
	return
}
