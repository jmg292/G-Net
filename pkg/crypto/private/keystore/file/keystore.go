package file

import (
	"crypto"
	"fmt"

	"github.com/cloudflare/circl/dh/x25519"
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
	encryptionKey         x25519.Key
	encryptionKeyType     gcrypt.SupportedKeyType
	authenticationKey     crypto.PrivateKey
	authenticationKeyType gcrypt.SupportedKeyType
	deviceKey             crypto.PrivateKey
}

func New(path string) (*fileKeyStore, error) {
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}
