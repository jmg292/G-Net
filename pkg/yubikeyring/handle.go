package yubikeyring

import (
	"github.com/go-piv/piv-go/piv"
	gnet "github.com/jmg292/G-Net/pkg/gneterrs"
)

func (y *Backend) getYubikeyHandle() (handle *piv.YubiKey, err error) {
	if y.handle == nil {
		err = gnet.ErrorKeystoreHandleClosed
	} else {
		y.handleMutex.Lock()
		handle = y.handle
	}
	return
}

func (y *Backend) releaseYubikeyHandle() {
	y.handleMutex.Unlock()
}
