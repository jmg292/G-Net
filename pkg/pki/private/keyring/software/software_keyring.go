package software

import (
	"crypto"

	"github.com/jmg292/G-Net/pkg/pki"
)

// Exists to facilitate user pki backups
type SoftwareKeyRing struct {
	signingKey            *crypto.PrivateKey
	signingKeyType        pki.SupportedKeyType
	encryptionKey         *crypto.PrivateKey
	encryptionKeyType     pki.SupportedKeyType
	authenticationKey     *crypto.PrivateKey
	authenticationKeyType pki.SupportedKeyType
}
