package yubikey

import (
	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/pkg/gnet"
)

func (y *Yubikey) getYubikeyHandle() (handle *piv.YubiKey, err error) {
	if y.handle == nil {
		err = gnet.ErrorKeystoreHandleClosed
	} else {
		y.mutex.Lock()
		handle = y.handle
	}
	return
}

func (y *Yubikey) releaseYubikeyHandle() {
	y.mutex.Unlock()
}
