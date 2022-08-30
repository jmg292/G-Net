package extensions

import (
	"encoding/asn1"

	gnet "github.com/jmg292/G-Net/pkg/gneterrs"
	"github.com/jmg292/G-Net/pkg/keyring"
)

func GetSlotExtensionOID(certificateSlot keyring.KeySlot) (oid asn1.ObjectIdentifier, err error) {
	switch certificateSlot {
	case keyring.SigningKeySlot:
		oid = OIDAuthenticate
	case keyring.AuthenticationKeySlot:
		oid = OIDVerifyAuthentication
	case keyring.EncryptionKeySlot:
		oid = OIDEncrypt
	default:
		err = gnet.ErrorInvalidKeySlot
	}
	return
}
