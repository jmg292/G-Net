package keyslot

import (
	"bytes"
	"crypto/rand"
)

func (slot *KeySlot) getNonce() []byte {
	return slot[nonceOffset : nonceSize+nonceOffset]
}

func (slot *KeySlot) generateNonce() {
	nonce := slot.getNonce()
	rand.Read(nonce)
}

func (slot *KeySlot) nonceIsEmpty() bool {
	return bytes.Equal(slot.getNonce(), empty.getNonce())
}
