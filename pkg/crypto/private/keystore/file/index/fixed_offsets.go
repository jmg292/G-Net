package index

func (*index) PreambleOffset() int {
	return int(preamble)
}

func (*index) PreambleSize() int {
	return int(versionNumber - preamble)
}

func (*index) VersionNumberOffset() int {
	return int(versionNumber)
}

func (*index) VersionNumberSize() int {
	return int(indexBytes - versionNumber)
}

func (*index) IndexOffset() int {
	return int(indexBytes)
}

func (*index) IndexSize() int {
	return Size
}

func (*index) SaltOffset() int {
	return int(salt)
}

func (*index) SaltSize() int {
	return int(adminSlot - salt)
}

func (*index) AdminSlotOffset() int {
	return int(adminSlot)
}

func (*index) AdminSlotSize() int {
	return int(keySlotBase - adminSlot)
}
