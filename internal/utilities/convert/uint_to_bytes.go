package convert

func UInt64ToBytes(value uint64) []byte {
	returnValue := make([]byte, 8)
	for i := uint(0); i < 8; i++ {
		returnValue[i] = byte((value >> (i * 8)) & 0xFF)
	}
	return returnValue
}

func UInt32ToBytes(value uint32) []byte {
	returnValue := make([]byte, 4)
	for i := uint32(0); i < 4; i++ {
		returnValue[i] = byte((value >> (i * 4)) & 0xFF)
	}
	return returnValue
}

func UInt16ToBytes(value uint16) []byte {
	returnValue := make([]byte, 2)
	for i := uint16(0); i < 2; i++ {
		returnValue[i] = byte((value >> (i * 2)) & 0xFF)
	}
	return returnValue
}
