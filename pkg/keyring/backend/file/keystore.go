package file

import (
	"fmt"

	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/keyring/backend/file/adminslot"
	"github.com/jmg292/G-Net/pkg/keyring/backend/file/index"
	"github.com/jmg292/G-Net/pkg/keyring/backend/file/keyslot"
	"github.com/jmg292/G-Net/pkg/keyring/backend/file/meta"
)

// Exists to facilitate user identity backups
type fileKeyStore struct {
	metadata          meta.Meta
	preamble          []byte
	versionNumber     byte
	index             index.Index
	salt              []byte
	adminSlot         adminslot.AdminSlot
	signingKeySlot    keyslot.KeySlot
	authKeySlot       keyslot.KeySlot
	encryptionKeySlot keyslot.KeySlot
	deviceKeySlot     keyslot.KeySlot
	certificateStore  []byte
	validationTag     []byte
}

func New(path string) (*fileKeyStore, error) {
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}
