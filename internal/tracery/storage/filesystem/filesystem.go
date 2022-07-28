package filesystem

import (
	"encoding/hex"
	"fmt"
	"os"
	"path"

	"github.com/jmg292/G-Net/pkg/gnet"
)

type blockStorageDirectory struct {
	path string
}

func NewBlockStorageDirectory(storagePath string) *blockStorageDirectory {
	return &blockStorageDirectory{path: storagePath}
}

func (storage *blockStorageDirectory) lockFile() string {
	return path.Join(storage.path, "storage.lock")
}

func (storage *blockStorageDirectory) blockPath(blockId []byte) string {
	blockHexId := hex.EncodeToString(blockId)
	return path.Join(storage.path, fmt.Sprintf("%s.wumbo", blockHexId))
}

func (*blockStorageDirectory) error(err error) error {
	if os.IsNotExist(err) {
		err = fmt.Errorf(string(gnet.ErrorBlockNotFound))
	} else if os.IsExist(err) {
		err = fmt.Errorf(string(gnet.ErrorBlockExists))
	}
	return err
}

func (storage *blockStorageDirectory) lock() error {
	_, err := os.Create(storage.lockFile())
	if os.IsExist(err) {
		return fmt.Errorf(string(gnet.ErrorStorageLocked))
	}
	return err
}

func (storage *blockStorageDirectory) unlock() error {
	err := os.Remove(storage.lockFile())
	if os.IsNotExist(err) {
		err = nil
	}
	return err
}
