package yubikey

import (
	"sync"

	"github.com/go-piv/piv-go/piv"
)

var instanceMutex = &sync.Mutex{}

var (
	instance *Yubikey
)

type Yubikey struct {
	handle *piv.YubiKey
	mutex  *sync.Mutex
}
