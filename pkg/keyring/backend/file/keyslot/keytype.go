package keyslot

import (
	"github.com/jmg292/G-Net/internal/utilities/convert"
	"github.com/jmg292/G-Net/pkg/keyring"
)

func (slot *KeySlot) getKeyTypeBytes() []byte {
	return slot[keyTypeOffset : keyTypeSize+keyTypeOffset]
}

func (slot *KeySlot) KeyType() keyring.SupportedKeyType {
	return keyring.SupportedKeyType(convert.BytesToUInt16(slot.getKeyTypeBytes()))
}

func (slot *KeySlot) SetKeyType(keyType keyring.SupportedKeyType) {
	copy(slot[:2], convert.UInt16ToBytes(uint16(keyType)))
}
