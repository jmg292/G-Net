package yubikey

import (
	"crypto/subtle"

	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/pkg/gnet"
)

func (y *yubikeyStorageBackend) generateKey(slot piv.Slot, alg piv.Algorithm) (err error) {
	keyTemplate := piv.Key{
		Algorithm:   alg,
		PINPolicy:   piv.PINPolicyAlways,
		TouchPolicy: piv.TouchPolicyAlways,
	}
	if err = y.assertOpenAndUnlocked(); err == nil {
		_, err = y.handle.GenerateKey(*y.metadata.ManagementKey, slot, keyTemplate)
	}
	return
}

func (y *yubikeyStorageBackend) generateEncryptionKey() (publicBytes x25519PublicBytes, err error) {
	publicBytes.GenerateSalt()
	if privateKey, e := y.deriveX25519PrivateKey(publicBytes.Salt()); e != nil {
		err = e
	} else {
		subtle.ConstantTimeCopy(1, publicBytes.Key(), privateKey.PublicBytes())
	}
	return
}

func (y *yubikeyStorageBackend) storeEncryptionKey(public x25519PublicBytes) (err error) {
	keyPolicy := piv.Key{
		PINPolicy:   piv.PINPolicyAlways,
		TouchPolicy: piv.TouchPolicyAlways,
	}
	if err = y.assertOpenAndUnlocked(); err != nil {
		// This branch intentionally left blank :O
	} else if slot1, ok := piv.RetiredKeyManagementSlot(0x95); !ok {
		err = gnet.ErrorInvalidKeySlot
	} else if slot2, ok := piv.RetiredKeyManagementSlot(0x94); !ok {
		err = gnet.ErrorInvalidKeySlot
	} else if err = y.handle.SetPrivateKeyInsecure(*y.metadata.ManagementKey, slot1, public[:24], keyPolicy); err == nil {
		err = y.handle.SetPrivateKeyInsecure(*y.metadata.ManagementKey, slot2, public[24:], keyPolicy)
	}
	return
}
