package keyslot

import (
	"github.com/jmg292/G-Net/internal/utilities/convert"
	gcrypt "github.com/jmg292/G-Net/pkg/crypto"
)

func (slot *KeySlot) getKeyTypeBytes() []byte {
	return slot[keyTypeOffset : keyTypeSize+keyTypeOffset]
}

func (slot *KeySlot) KeyType() gcrypt.SupportedKeyType {
	return gcrypt.SupportedKeyType(convert.BytesToUInt16(slot.getKeyTypeBytes()))
}

func (slot *KeySlot) SetKeyType(keyType gcrypt.SupportedKeyType) {
	copy(slot[:2], convert.UInt16ToBytes(uint16(keyType)))
}
