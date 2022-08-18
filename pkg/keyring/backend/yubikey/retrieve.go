package yubikey

import (
	"crypto"
	"crypto/x509"

	"github.com/awnumar/memguard"
	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/pkg/gnet"
)

func (y *Yubikey) getPin() (pin *memguard.LockedBuffer, err error) {
	y.pinMutex.Lock()
	defer y.pinMutex.Unlock()
	if y.pin == nil {
		err = gnet.ErrorKeystoreLocked
	} else {
		pin, err = y.pin.Open()
	}
	return
}

func (y *Yubikey) getManagementKey() (managementKey *[24]byte, err error) {
	if pin, e := y.getPin(); e != nil {
		err = e
	} else if handle, e := y.getYubikeyHandle(); e != nil {
		err = e
	} else if metadata, e := handle.Metadata(pin.String()); e != nil {
		err = e
	} else if metadata == nil {
		err = gnet.ErrorKeyNotFound
	} else {
		managementKey = metadata.ManagementKey
	}
	return
}

func (y *Yubikey) getPublicKey(slot piv.Slot) (key crypto.PublicKey, err error) {
	if handle, e := y.getYubikeyHandle(); e != nil {
		err = e
	} else {
		defer y.releaseYubikeyHandle()
		if key, err = handle.Attest(slot); err != nil {
			if err == piv.ErrNotFound {
				err = gnet.ErrorKeyNotFound
			}
		} else if key == nil {
			err = gnet.ErrorKeyNotFound
		}
		key = key.(x509.Certificate).PublicKey
	}
	return
}

func (y *Yubikey) getPrivateKey(slot piv.Slot) (key crypto.PrivateKey, err error) {
	if pubkey, e := y.getPublicKey(slot); e != nil {
		err = e
	} else if handle, e := y.getYubikeyHandle(); e != nil {
		err = e
	} else {
		defer y.releaseYubikeyHandle()
		if pin, e := y.pin.Open(); e != nil {
			err = e
		} else {
			defer pin.Wipe()
			key, err = handle.PrivateKey(slot, pubkey, piv.KeyAuth{PIN: pin.String()})
		}
	}
	return
}
