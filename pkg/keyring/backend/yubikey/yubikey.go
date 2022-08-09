package yubikey

import (
	"sync"

	"github.com/awnumar/memguard"
	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/pkg/keyring/key"
)

type YubikeyStorageBackend struct {
	name          string
	pin           *memguard.Enclave
	metadata      *piv.Metadata
	handle        *piv.YubiKey
	encryptionKey *key.X25519PrivateKey
	mutex         sync.Mutex
}

func New() *YubikeyStorageBackend {
	// Safely terminate in case of an interrupt signal
	memguard.CatchInterrupt()
	return &YubikeyStorageBackend{}
}

func (y *YubikeyStorageBackend) Reset() (err error) {
	if handle, e := y.getHandle(); e != nil {
		err = e
	} else {
		defer y.releaseHandle()
		err = handle.Reset()
		y.Lock()
	}
	return
}

func (y *YubikeyStorageBackend) UpdatePIN(newPin *memguard.Enclave) (err error) {
	var oldPin *memguard.LockedBuffer
	if oldPin, err = y.getPin(); err != nil {
		oldPin = memguard.NewBufferFromBytes([]byte(piv.DefaultPIN))
	}
	defer oldPin.Destroy()
	if handle, e := y.getHandle(); e != nil {
		err = e
	} else {
		defer y.releaseHandle()
		if pin, e := newPin.Open(); e != nil {
			err = e
		} else {
			defer pin.Destroy()
			if err = handle.SetPIN(oldPin.String(), pin.String()); err != nil {
				y.pin = newPin
			}
		}
	}
	return
}
