package software

import (
	"crypto"

	gcrypt "github.com/jmg292/G-Net/pkg/crypto"
)

// Exists to facilitate user crypto backups
type SoftwareKeyRing struct {
	signingKey            *crypto.PrivateKey
	signingKeyType        gcrypt.SupportedKeyType
	encryptionKey         *crypto.PrivateKey
	encryptionKeyType     gcrypt.SupportedKeyType
	authenticationKey     *crypto.PrivateKey
	authenticationKeyType gcrypt.SupportedKeyType
}
