package file

import (
	"crypto/rand"
	"os"

	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/keyring/backend/file/adminslot"
)

func (f *fileKeyStore) createFile() (handle *os.File, err error) {
	if _, e := os.Stat(f.metadata.Path()); !os.IsNotExist(e) {
		err = gnet.ErrorFileAlreadyExists
	} else {
		handle, err = os.OpenFile(f.metadata.Path(), os.O_CREATE|os.O_WRONLY, 0600)
	}
	return
}

func (f *fileKeyStore) populateMetadata() {
	f.preamble = []byte(preamble)
	f.versionNumber = 1
	f.adminSlot = adminslot.New()
	f.salt = make([]byte, f.index.SaltSize())
	rand.Read(f.salt)
}
