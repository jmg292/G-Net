package file

import (
	"crypto"
	"fmt"

	gcrypt "github.com/jmg292/G-Net/pkg/crypto"
	"github.com/jmg292/G-Net/pkg/gnet"
)

const managementKeySize = 24

// Exists to facilitate user identity backups
type fileKeyStore struct {
	path                  string
	managementKey         []byte
	keyEncryptionKeySalt  []byte
	signingKey            crypto.PrivateKey
	signingKeyType        gcrypt.SupportedKeyType
	encryptionKey         crypto.PrivateKey
	encryptionKeyType     gcrypt.SupportedKeyType
	authenticationKey     crypto.PrivateKey
	authenticationKeyType gcrypt.SupportedKeyType
	deviceKey             crypto.PrivateKey
}

func New(path string) (*fileKeyStore, error) {
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}
