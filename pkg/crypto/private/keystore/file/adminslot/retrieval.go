package adminslot

func (slot *AdminSlot) getKey(t keyType) []byte {
	keyOffset := int(t) * keySize
	return slot[keyOffset : keyOffset+keySize]
}

func (slot *AdminSlot) ManagementKey() []byte {
	return slot.getKey(managementKey)
}

func (slot *AdminSlot) KdfSalt() []byte {
	return slot.getKey(kdfSalt)
}
