package yubikey

import (
	"github.com/go-piv/piv-go/piv"
)

type yubikeyStorageBackend struct {
	name     string
	metadata *piv.Metadata
	handle   *piv.YubiKey
}
