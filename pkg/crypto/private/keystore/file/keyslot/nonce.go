package keyslot

import (
	"crypto/rand"

	gcrypt "github.com/jmg292/G-Net/pkg/crypto"
)

func (slot *keySlot) getNonce() []byte {
	return slot[2 : gcrypt.NonceSize+2]
}

func (slot *keySlot) generateNonce() []byte {
	nonce := slot.getNonce()
	rand.Read(nonce)
	return nonce
}
