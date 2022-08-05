package adminslot

import (
	"bytes"
	"crypto/rand"

	"golang.org/x/crypto/chacha20poly1305"
)

const (
	keySize = 24
	Size    = (keySize * 2) + chacha20poly1305.Overhead
)

type keyType int

const (
	managementKey keyType = iota
	kdfSalt
)

type AdminSlot [Size]byte

var empty AdminSlot

func Empty() *AdminSlot {
	var slot AdminSlot
	return &slot
}

func New() *AdminSlot {
	var slot AdminSlot
	rand.Read(slot[:])
	return &slot
}

func (slot *AdminSlot) IsEmpty() bool {
	return bytes.Equal(slot[:], empty[:])
}
