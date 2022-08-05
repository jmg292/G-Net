package keyslot

import (
	"bytes"
	"crypto"
	"fmt"

	gcrypt "github.com/jmg292/G-Net/pkg/crypto"
	"github.com/jmg292/G-Net/pkg/gnet"
	"golang.org/x/crypto/chacha20poly1305"
)

const Size int = 100
const (
	keyTypeOffset = 0
	keyTypeSize   = 2
	nonceOffset   = keyTypeSize
	nonceSize     = gcrypt.NonceSize
	keyOffset     = nonceOffset + nonceSize
	maxKeySize    = Size - keyOffset - chacha20poly1305.Overhead
)

type keySlot [Size]byte

var empty keySlot

func NewEmpty() *keySlot {
	var slot keySlot
	return &slot
}

func New(keyType gcrypt.SupportedKeyType, key crypto.PrivateKey, managementKey []byte) (slot *keySlot, err error) {
	slot = NewEmpty()
	err = fmt.Errorf(string(gnet.ErrorNotYetImplemented))
	return
}

func (slot *keySlot) IsEmpty() bool {
	return bytes.Equal(NewEmpty()[:], slot[:])
}
