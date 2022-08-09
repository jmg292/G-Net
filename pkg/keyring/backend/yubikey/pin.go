package yubikey

import (
	"github.com/awnumar/memguard"
	"github.com/jmg292/G-Net/pkg/gnet"
)

func (y *YubikeyStorageBackend) getPin() (pinbuffer *memguard.LockedBuffer, err error) {
	if y.pin == nil {
		err = gnet.ErrorKeystoreLocked
	} else {
		pinbuffer, err = y.pin.Open()
	}
	return
}

func (y *YubikeyStorageBackend) clearPin() (err error) {
	if y.pin == nil {
		err = gnet.ErrorKeystoreLocked
	} else {
		y.pin = nil
		memguard.Purge()
	}
	return
}
