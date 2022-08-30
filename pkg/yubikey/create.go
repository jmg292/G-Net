package yubikey

import (
	"crypto/rand"

	"github.com/go-piv/piv-go/piv"
	gnet "github.com/jmg292/G-Net/pkg/gneterrs"
	"github.com/jmg292/G-Net/pkg/keyring"
)

func (y *Backend) createPivKey(slot piv.Slot, alg piv.Algorithm) (err error) {
	if key, e := y.getPrivateKey(slot); key != nil {
		err = gnet.ErrorKeyAlreadyExists
	} else if e != nil && e != gnet.ErrorKeyNotFound {
		err = e
	} else if mk, e := y.GetPrivateKey(keyring.ManagementKeySlot); e != nil {
		err = e
	} else if managementKey, ok := mk.(*[24]byte); !ok {
		err = gnet.ErrorInvalidManagementKey
	} else if handle, e := y.getYubikeyHandle(); e != nil {
		err = e
	} else {
		defer y.releaseYubikeyHandle()
		_, err = handle.GenerateKey(*managementKey, slot, piv.Key{
			PINPolicy:   piv.PINPolicyAlways,
			TouchPolicy: piv.TouchPolicyAlways,
			Algorithm:   alg,
		})
	}
	return
}

func (y *Backend) createManagementKey() (err error) {
	var newKey [keyring.ManagementKeySize]byte
	if bytesRead, e := rand.Read(newKey[:]); e != nil {
		err = e
	} else if bytesRead != int(keyring.ManagementKeySize) {
		err = gnet.ErrorKeyGenFailed
	} else if managmentKey, e := y.GetPrivateKey(keyring.ManagementKeySlot); e != nil {
		err = e
	} else if managmentKey, ok := managmentKey.(*[24]byte); !ok {
		err = gnet.ErrorInvalidManagementKey
	} else if handle, e := y.getYubikeyHandle(); e != nil {
		err = e
	} else {
		defer y.releaseYubikeyHandle()
		if err = handle.SetManagementKey(*managmentKey, newKey); err == nil {
			err = handle.SetMetadata(newKey, &piv.Metadata{ManagementKey: &newKey})
		}
	}
	return
}
