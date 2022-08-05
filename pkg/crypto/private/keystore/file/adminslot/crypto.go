package adminslot

import (
	"github.com/jmg292/G-Net/pkg/crypto/kdf"
	"golang.org/x/crypto/chacha20poly1305"
)

func (slot *AdminSlot) Lock(pin []byte, salt []byte) (err error) {
	if cipher, err := chacha20poly1305.NewX(kdf.DeriveKey(pin, salt)); err == nil {
		cipher.Seal(slot[:], salt[:cipher.NonceSize()], slot[:cipher.Overhead()], nil)
	}
	return
}

func (slot *AdminSlot) Unlock(pin []byte, salt []byte) (err error) {
	if cipher, err := chacha20poly1305.NewX(kdf.DeriveKey(pin, salt)); err == nil {
		if _, err = cipher.Open(slot[:cipher.Overhead()], salt[:cipher.NonceSize()], slot[:], nil); err != nil {
			return err
		}
	}
	return
}
