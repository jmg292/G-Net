package yubikey

import (
	"strings"

	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/pkg/gnet"
)

func (y *YubikeyStorageBackend) getCardName() (name string, err error) {
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
