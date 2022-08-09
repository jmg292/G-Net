package yubikey

import (
	"crypto"

	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/pkg/gnet"
)

func (y *YubikeyStorageBackend) getPivPrivateKey(slot piv.Slot, publickey crypto.PublicKey) (privatekey crypto.PrivateKey, err error) {
	if pin, e := y.getPin(); e != nil {
		err = e
	} else {
		defer pin.Destroy()
		if handle, e := y.getHandle(); e != nil {
			err = e
		} else {
			defer y.releaseHandle()
			if privatekey, err = handle.PrivateKey(slot, publickey, piv.KeyAuth{PIN: pin.String()}); err == nil && privatekey == nil {
				err = gnet.ErrorKeyNotFound
			}
		}
	}
	return
}

func (y *YubikeyStorageBackend) getX25519PublicBytes() (public *x25519PublicBytes, err error) {
	if slot1, slot2, e := y.getX25519KeySlots(); e != nil {
		err = e
	} else if firsthalf, e := y.getPivPrivateKey(slot1, nil); e != nil {
		err = e
	} else if secondhalf, e := y.getPivPrivateKey(slot2, nil); e != nil {
		err = e
	} else if part1, ok := firsthalf.([24]byte); !ok {
		err = gnet.ErrorInvalidPublicKey
	} else if part2, ok := secondhalf.([24]byte); !ok {
		err = gnet.ErrorInvalidPublicKey
	} else if publicbytes, e := newX25519PublicBytes(append(part1[:], part2[:]...)); e != nil {
		err = e
	} else {
		public = &publicbytes
	}
	return
}

func (y *YubikeyStorageBackend) getX25519PrivateKey() (private crypto.PrivateKey, err error) {
	if y.encryptionKey != nil {
		private = y.encryptionKey
	} else if publicbytes, e := y.getX25519PublicBytes(); e != nil {
		err = e
	} else if y.encryptionKey, err = y.deriveX25519PrivateKey(publicbytes.Salt()); err == nil {
		private = y.encryptionKey
	}
	return
}

func (y *YubikeyStorageBackend) getX25519PublicKey() (public crypto.PublicKey, err error) {
	if y.encryptionKey == nil {
		if _, err = y.getX25519PrivateKey(); err != nil {
			return
		}
	}
	public = y.encryptionKey.PublicKey()
	return
}
