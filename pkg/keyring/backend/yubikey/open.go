package yubikey

import (
	"strings"

	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/pkg/gnet"
)

func (y *yubikeyStorageBackend) getCardName() (name string, err error) {
	if names, e := piv.Cards(); err != nil {
		err = e
	} else {
		for _, n := range names {
			if strings.Contains(strings.ToLower(n), "yubikey") {
				name = n
				break
			}
		}
	}
	if name == "" {
		err = gnet.ErrorKeystoreNotFound
	}
	return
}

func (y *yubikeyStorageBackend) assertOpen() (err error) {
	if y.handle == nil {
		err = gnet.ErrorInvalidHandle
	}
	return
}

func (y *yubikeyStorageBackend) assertOpenAndUnlocked() (err error) {
	if err = y.assertOpen(); err == nil {
		if y.metadata == nil {
			err = gnet.ErrorKeystoreLocked
		}
	}
	return
}
