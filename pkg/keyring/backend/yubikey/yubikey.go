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
