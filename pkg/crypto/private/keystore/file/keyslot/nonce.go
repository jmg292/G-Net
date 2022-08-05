package keyslot

import (
	"bytes"
	"crypto/rand"
)

func (slot *keySlot) getNonce() []byte {
	return slot[nonceOffset : nonceSize+nonceOffset]
}

func (slot *keySlot) generateNonce() {
	nonce := slot.getNonce()
	rand.Read(nonce)
}

func (slot *keySlot) nonceIsEmpty() bool {
	return bytes.Equal(slot.getNonce(), empty.getNonce())
}
