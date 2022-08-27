package yubikeyring

import (
	"sync"

	"github.com/awnumar/memguard"
	"github.com/go-piv/piv-go/piv"
)

var instanceMutex = &sync.Mutex{}

var (
	instance *Backend
)

type Backend struct {
	handle      *piv.YubiKey
	handleMutex *sync.Mutex
	pin         *memguard.Enclave
	pinMutex    *sync.Mutex
}

func (y *Backend) Reset() (err error) {
	if handle, e := y.getYubikeyHandle(); e != nil {
		err = e
	} else {
		defer y.releaseYubikeyHandle()
		err = handle.Reset()
	}
	return
}
