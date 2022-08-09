package yubikey

import (
	"sync"

	"github.com/go-piv/piv-go/piv"
)

type YubikeyStorageBackend struct {
	name     string
	metadata *piv.Metadata
	handle   *piv.YubiKey
	mutex    sync.Mutex
}

func New() *YubikeyStorageBackend {
	return &YubikeyStorageBackend{}
}

func (y *YubikeyStorageBackend) Reset() (err error) {
	if handle, e := y.getHandle(); e != nil {
		err = e
	} else {
		defer y.releaseHandle()
		err = handle.Reset()
	}
	return
}
