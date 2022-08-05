package adminslot

func (slot *adminSlot) getKey(t keyType) []byte {
	key := make([]byte, keySize)
	keyOffset := int(t) * keySize
	copy(key, slot[keyOffset:keyOffset+keySize])
	return key
}

func (slot *adminSlot) ManagementKey() []byte {
	return slot.getKey(managementKey)
}

func (slot *adminSlot) KdfSalt() []byte {
	return slot.getKey(kdfSalt)
}
