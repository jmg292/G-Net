package yubikey

import (
	"strings"
	"sync"

	"github.com/go-piv/piv-go/piv"
)

func getYubikeyHandle() (handle *piv.YubiKey, err error) {
	if cards, e := piv.Cards(); e != nil {
		e = err
	} else {
		for _, name := range cards {
			if strings.Contains(strings.ToLower(name), "yubikey") {
				handle, err = piv.Open(name)
				break
			}
		}
	}
	return
}

func New() (backend *Yubikey, err error) {
	instanceMutex.Lock()
	defer instanceMutex.Unlock()

	if instance == nil {
		if handle, e := getYubikeyHandle(); e != nil {
			err = e
		} else {
			instance = &Yubikey{
				handle: handle,
				mutex:  &sync.Mutex{},
			}
		}
	}

	backend = instance

	return
}
