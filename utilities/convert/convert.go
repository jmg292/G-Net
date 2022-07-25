package convert

func UInt64ToBinary(value uint64) []byte {
	returnValue := make([]byte, 8)
	for i := uint(0); i < 8; i++ {
		returnValue[i] = byte((value >> (i * 8)) & 0xFF)
	}
	return returnValue
}

func UInt64FromBinary(value []byte) uint64 {
	returnValue := uint64(0)
	for i := uint64(0); i < 8; i++ {
		returnValue |= uint64(value[i]) << (i * 8)
	}
	return returnValue
}

func UInt32ToBinary(value uint32) []byte {
	returnValue := make([]byte, 4)
	for i := uint32(0); i < 4; i++ {
		returnValue[i] = byte((value >> (i * 4)) & 0xFF)
	}
	return returnValue
}

func UInt32FromBinary(value []byte) uint32 {
	returnValue := uint32(0)
	for i := uint32(0); i < 4; i++ {
		returnValue |= uint32(value[i]) << (i * 4)
	}
	return returnValue
}

func UInt16ToBinary(value uint16) []byte {
	returnValue := make([]byte, 2)
	for i := uint16(0); i < 2; i++ {
		returnValue[i] = byte((value >> (i * 2)) & 0xFF)
	}
	return returnValue
}

func UInt16FromBinary(value []byte) uint16 {
	returnValue := uint16(0)
	for i := uint16(0); i < 2; i++ {
		returnValue |= uint16(value[i]) << (i * 2)
	}
	return returnValue
}
