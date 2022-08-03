package file

import (
	"fmt"
	"os"

	"github.com/jmg292/G-Net/pkg/gnet"
)

func (f *fileKeyStore) saveFile(handle *os.File) (err error) {
	if keyStoreBytes, err := f.toBytes(); err == nil {
		handle.Seek(0, 0)
		handle.Write(keyStoreBytes)
	}
	return
}

func (f *fileKeyStore) createFile() (handle *os.File, err error) {
	err = fmt.Errorf(string(gnet.ErrorNotYetImplemented))
	return
}
