package yubikey

import (
	"sync"

	"github.com/awnumar/memguard"
	"github.com/go-piv/piv-go/piv"
)

var instanceMutex = &sync.Mutex{}

var (
	instance *Yubikey
)

type Yubikey struct {
	handle      *piv.YubiKey
	handleMutex *sync.Mutex
	pin         *memguard.Enclave
	pinMutex    *sync.Mutex
}

func (y *Yubikey) Reset() (err error) {
	if handle, e := y.getYubikeyHandle(); e != nil {
		err = e
	} else {
		err = handle.Reset()
	}
	return
}
