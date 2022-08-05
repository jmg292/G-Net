package file

import (
	"fmt"

	"github.com/jmg292/G-Net/pkg/gnet"
)

// Exists to facilitate user identity backups
type fileKeyStore struct {
}

func New(path string) (*fileKeyStore, error) {
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}
