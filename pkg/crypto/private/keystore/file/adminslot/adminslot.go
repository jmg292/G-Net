package management_key

import (
	"crypto/rand"

	"golang.org/x/crypto/chacha20poly1305"
)

const (
	keySize = 24
	Size    = (keySize * 2) + chacha20poly1305.Overhead
)

type adminSlot [Size]byte

func New() *adminSlot {
	var slot adminSlot
	rand.Read(slot[:])
	return &slot
}
