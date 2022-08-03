package file

import (
	"fmt"

	gcrypt "github.com/jmg292/G-Net/pkg/crypto"
	"github.com/jmg292/G-Net/pkg/gnet"
)

func (f *fileKeyStore) getStorableKey(keyType gcrypt.SupportedKeyType) (keyBytes []byte, err error) {
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) toBytes() (keyStore []byte, err error) {
	keyStore = append(f.managementKey)
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}
