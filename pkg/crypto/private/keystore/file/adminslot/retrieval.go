package adminslot

func (slot *adminSlot) getKey(t keyType) []byte {
	keyOffset := int(t) * keySize
	return slot[keyOffset : keyOffset+keySize]
}

func (slot *adminSlot) ManagementKey() []byte {
	return slot.getKey(managementKey)
}

func (slot *adminSlot) KdfSalt() []byte {
	return slot.getKey(kdfSalt)
}
