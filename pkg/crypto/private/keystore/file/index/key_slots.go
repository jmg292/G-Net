package index

import "github.com/jmg292/G-Net/pkg/crypto"

const keySlotSize = 100

func (*index) getSlotOffset(keySlot crypto.KeySlot) int {
	relativeOffset := int(keySlot) * keySlotSize
	return int(keySlotBase) + relativeOffset
}

func (i *index) SigningKeySlotOffset() int {
	return i.getSlotOffset(crypto.SigningKeySlot)
}

func (*index) SigningKeySlotSize() int {
	return keySlotSize
}

func (i *index) AuthenticationKeySlotOffset() int {
	return i.getSlotOffset(crypto.AuthenticationKeySlot)
}

func (*index) AuthenticationKeySlotSize() int {
	return keySlotSize
}

func (i *index) EncryptionKeySlotOffset() int {
	return i.getSlotOffset(crypto.EncryptionKeySlot)
}

func (*index) EncryptionKeySlotSize() int {
	return keySlotSize
}

func (i *index) DeviceKeySlotOffset() int {
	return i.getSlotOffset(crypto.DeviceKeySlot)
}

func (*index) DeviceKeySlotSize() int {
	return keySlotSize
}
