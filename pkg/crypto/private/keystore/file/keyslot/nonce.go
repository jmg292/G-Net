package keyslot

import (
	"bytes"
	"crypto/rand"
)

func (slot *keySlot) getNonce() []byte {
	return slot[nonceOffset : nonceSize+nonceOffset]
}

func (slot *keySlot) generateNonce() []byte {
	nonce := slot.getNonce()
	rand.Read(nonce)
	return nonce
}

func (slot *keySlot) nonceIsEmpty() bool {
	return bytes.Equal(slot.getNonce(), empty.getNonce())
}
