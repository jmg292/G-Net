package yubikey

import (
	"sync"

	"github.com/awnumar/memguard"
	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/pkg/keyring/key"
)

type YubikeyStorageBackend struct {
	name          string
	pin           *memguard.Enclave
	metadata      *piv.Metadata
	handle        *piv.YubiKey
	encryptionKey *key.X25519PrivateKey
	mutex         sync.Mutex
}

func New() *YubikeyStorageBackend {
	// Safely terminate in case of an interrupt signal
	memguard.CatchInterrupt()
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
