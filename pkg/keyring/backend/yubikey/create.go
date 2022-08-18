package yubikey

import (
	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/keyring"
)

func (y *Yubikey) createPivKey(slot piv.Slot, alg piv.Algorithm) (err error) {
	if key, e := y.getPrivateKey(slot); key != nil {
		err = gnet.ErrorKeyAlreadyExists
	} else if e != nil && e != gnet.ErrorKeyNotFound {
		err = e
	} else if mk, e := y.GetPrivateKey(keyring.ManagementKeySlot); e != nil {
		err = e
	} else if managementKey, ok := mk.([24]byte); !ok {
		err = gnet.ErrorInvalidManagementKey
	} else if handle, e := y.getYubikeyHandle(); e != nil {
		err = e
	} else {
		defer y.releaseYubikeyHandle()
		_, err = handle.GenerateKey(managementKey, slot, piv.Key{
			PINPolicy:   piv.PINPolicyAlways,
			TouchPolicy: piv.TouchPolicyAlways,
			Algorithm:   alg,
		})
	}
	return
}
