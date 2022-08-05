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

type adminSlot [Size]byte

var empty adminSlot

func Empty() *adminSlot {
	var slot adminSlot
	return &slot
}

func New() *adminSlot {
	var slot adminSlot
	rand.Read(slot[:])
	return &slot
}

func (slot *adminSlot) IsEmpty() bool {
	return bytes.Equal(slot[:], empty[:])
}
