package yubikey

import (
	"sync"

	"github.com/go-piv/piv-go/piv"
)

type yubikeyStorageBackend struct {
	name     string
	metadata *piv.Metadata
	handle   *piv.YubiKey
	mutex    sync.Mutex
}

func New() *yubikeyStorageBackend {
	return &yubikeyStorageBackend{}
}

func (y *yubikeyStorageBackend) Reset() (err error) {
	if handle, e := y.getHandle(); e != nil {
		err = e
	} else {
		defer y.releaseHandle()
		err = handle.Reset()
	}
	return
}
