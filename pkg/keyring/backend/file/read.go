package file

import (
	"io"
	"os"

	"github.com/jmg292/G-Net/pkg/keyring"
	"github.com/jmg292/G-Net/pkg/keyring/backend/file/meta"
	"github.com/jmg292/G-Net/pkg/keyring/key/slot"
)

func (f *fileKeyStore) fillBytesFromOffset(b []byte, offset int, handle *os.File) (err error) {
	handle.Seek(0, 0)
	reader := io.NewSectionReader(handle, int64(offset), int64(len(b)))
	_, err = reader.Read(b)
	return
}

func (f *fileKeyStore) readPreamble(handle *os.File) (p []byte, err error) {
	p = make([]byte, f.index.PreambleSize())
	err = f.fillBytesFromOffset(p, f.index.PreambleOffset(), handle)
	return
}

func (f *fileKeyStore) readVersionNumber(handle *os.File) (versionNumber byte, err error) {
	v := make([]byte, f.index.VersionNumberSize())
	if err = f.fillBytesFromOffset(v, f.index.VersionNumberOffset(), handle); err == nil {
		versionNumber = v[0]
	}
	return
}

func (f *fileKeyStore) readSection(section meta.Section, handle *os.File) (content []byte, err error) {
	var offset, size int
	switch section {
	case meta.IndexSection:
		offset = f.index.IndexOffset()
		size = f.index.IndexSize()
	case meta.SaltSection:
		offset = f.index.SaltOffset()
		size = f.index.SaltSize()
	case meta.AdminSlotSection:
		offset = f.index.AdminSlotOffset()
		size = f.index.AdminSlotSize()
	case meta.SigningKeySlotSection:
		offset = f.index.GetKeySlotOffset(keyring.SigningKeySlot)
		size = slot.Size
	case meta.AuthKeySlotSection:
		offset = f.index.GetKeySlotOffset(keyring.AuthenticationKeySlot)
		size = slot.Size
	case meta.DeviceKeySlotSection:
		offset = f.index.GetKeySlotOffset(keyring.DeviceKeySlot)
		size = slot.Size
	case meta.EncryptionKeySlotSection:
		offset = f.index.GetKeySlotOffset(keyring.EncryptionKeySlot)
		size = slot.Size
	case meta.CertificateStoreSection:
		offset = f.index.CertificateStoreOffset()
		size = f.index.CertificateStoreSize()
	}
	content = make([]byte, size)
	err = f.fillBytesFromOffset(content, offset, handle)
	return
}
