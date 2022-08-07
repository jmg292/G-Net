package keyslot

import (
	"bytes"
	"crypto"
	"fmt"

	gcrypt "github.com/jmg292/G-Net/pkg/crypto"
	"github.com/jmg292/G-Net/pkg/gnet"
	"golang.org/x/crypto/chacha20poly1305"
)

const Size int = 128
const (
	keyTypeOffset = 0
	keyTypeSize   = 2
	nonceOffset   = keyTypeSize
	nonceSize     = gcrypt.NonceSize
	keyOffset     = nonceOffset + nonceSize
	maxKeySize    = Size - keyOffset - chacha20poly1305.Overhead
)

type KeySlot [Size]byte

var empty KeySlot

func NewEmpty() *KeySlot {
	var slot KeySlot
	return &slot
}

func New(keyType gcrypt.SupportedKeyType, key crypto.PrivateKey, managementKey []byte) (slot *KeySlot, err error) {
	slot = NewEmpty()
	err = fmt.Errorf(string(gnet.ErrorNotYetImplemented))
	return
}

func (slot *KeySlot) IsEmpty() bool {
	return bytes.Equal(NewEmpty()[:], slot[:])
}
