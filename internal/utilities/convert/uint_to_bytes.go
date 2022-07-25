package convert

func UInt64ToBytes(value uint64) []byte {
	returnValue := make([]byte, 8)
	for i := uint(0); i < 8; i++ {
		returnValue[i] = byte((value >> (i * 8)) & 0xFF)
	}
	return returnValue
}
