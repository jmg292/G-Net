package convert

func BytesToUInt64(value []byte) uint64 {
	returnValue := uint64(0)
	for i := uint64(0); i < 8; i++ {
		returnValue |= uint64(value[i]) << (i * 8)
	}
	return returnValue
}

func BytesToUInt32(value []byte) uint32 {
	returnValue := uint32(0)
	for i := uint32(0); i < 4; i++ {
		returnValue |= uint32(value[i]) << (i * 4)
	}
	return returnValue
}

func BytesToUInt16(value []byte) uint16 {
	returnValue := uint16(0)
	for i := uint16(0); i < 2; i++ {
		returnValue |= uint16(value[i]) << (i * 2)
	}
	return returnValue
}
