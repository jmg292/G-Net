package software

import (
	"crypto"

	"github.com/jmg292/G-Net/pkg/identity"
)

// Exists to facilitate user identity backups
type SoftwareKeyRing struct {
	signingKey            *crypto.PrivateKey
	signingKeyType        identity.SupportedKeyType
	encryptionKey         *crypto.PrivateKey
	encryptionKeyType     identity.SupportedKeyType
	authenticationKey     *crypto.PrivateKey
	authenticationKeyType identity.SupportedKeyType
}
