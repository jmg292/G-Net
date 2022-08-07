package kdf

import (
	"fmt"

	"github.com/cloudflare/circl/dh/x25519"
	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/keyring"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/sha3"
)

func DeriveKey(keyMaterial []byte, salt []byte) []byte {
	return pbkdf2.Key(keyMaterial, salt, keyring.KdfIterations, int(keyring.SymmetricKeySize), sha3.New256)
}

func ExchangeKey(privateKey x25519.Key, publicKey x25519.Key) (key []byte, err error) {
	var secret x25519.Key
	if ok := x25519.Shared(&secret, &privateKey, &publicKey); !ok {
		key = DeriveKey(secret[:], publicKey[:])
	} else {
		err = fmt.Errorf(string(gnet.ErrorKeyExchangeFailed))
	}
	return
}
