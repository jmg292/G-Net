package slot

import "golang.org/x/crypto/chacha20poly1305"

func (slot *KeySlot) Lock(kek []byte) (err error) {
	if slot.nonceIsEmpty() {
		slot.generateNonce()
	}
	if cipher, err := chacha20poly1305.NewX(kek); err == nil {
		cipher.Seal(slot.getKeyWithOverhead(), slot.getNonce(), slot.GetKey(), slot.getKeyTypeBytes())
	}
	return
}

func (slot *KeySlot) Unlock(kek []byte) (err error) {
	if cipher, err := chacha20poly1305.NewX(kek); err == nil {
		_, err = cipher.Open(slot.GetKey(), slot.getNonce(), slot.getKeyWithOverhead(), slot.getKeyTypeBytes())
		return err
	}
	return
}
