package index

import "github.com/jmg292/G-Net/pkg/crypto"

const keySlotSize = 100

func (*Index) getSlotOffset(keySlot crypto.KeySlot) int {
	relativeOffset := int(keySlot) * keySlotSize
	return int(keySlotBase) + relativeOffset
}

func (i *Index) SigningKeySlotOffset() int {
	return i.getSlotOffset(crypto.SigningKeySlot)
}

func (*Index) SigningKeySlotSize() int {
	return keySlotSize
}

func (i *Index) AuthenticationKeySlotOffset() int {
	return i.getSlotOffset(crypto.AuthenticationKeySlot)
}

func (*Index) AuthenticationKeySlotSize() int {
	return keySlotSize
}

func (i *Index) EncryptionKeySlotOffset() int {
	return i.getSlotOffset(crypto.EncryptionKeySlot)
}

func (*Index) EncryptionKeySlotSize() int {
	return keySlotSize
}

func (i *Index) DeviceKeySlotOffset() int {
	return i.getSlotOffset(crypto.DeviceKeySlot)
}

func (*Index) DeviceKeySlotSize() int {
	return keySlotSize
}
