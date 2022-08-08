package file

import (
	"github.com/jmg292/G-Net/pkg/keyring/backend/file/adminslot"
	"github.com/jmg292/G-Net/pkg/keyring/backend/file/index"
	"github.com/jmg292/G-Net/pkg/keyring/backend/file/meta"
	"github.com/jmg292/G-Net/pkg/keyring/certificates"
	"github.com/jmg292/G-Net/pkg/keyring/key/slot"
)

// Exists to facilitate user identity backups
type fileKeyStore struct {
	metadata          meta.Meta
	preamble          []byte
	versionNumber     byte
	index             index.Index
	salt              []byte
	adminSlot         adminslot.AdminSlot
	signingKeySlot    slot.KeySlot
	authKeySlot       slot.KeySlot
	encryptionKeySlot slot.KeySlot
	deviceKeySlot     slot.KeySlot
	certificateStore  certificates.CertificateStore
	validationTag     []byte
}

func New(path string, pin []byte) (*fileKeyStore, error) {
	f := fileKeyStore{
		metadata: meta.New(path),
	}
	f.populateMetadata()
	handle, err := f.createFile()
	if err != nil {
		return nil, err
	}
	defer handle.Close()

}
