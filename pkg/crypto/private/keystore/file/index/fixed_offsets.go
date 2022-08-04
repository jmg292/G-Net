package index

func (*Index) PreambleOffset() int {
	return int(preamble)
}

func (*Index) PreambleSize() int {
	return int(versionNumber - preamble)
}

func (*Index) VersionNumberOffset() int {
	return int(versionNumber)
}

func (*Index) VersionNumberSize() int {
	return int(indexBytes - versionNumber)
}

func (*Index) IndexOffset() int {
	return int(indexBytes)
}

func (*Index) IndexSize() int {
	return size
}

func (*Index) SaltOffset() int {
	return int(salt)
}

func (*Index) SaltSize() int {
	return int(managementKeys - salt)
}

func (*Index) ManagementKeyOffset() int {
	return int(managementKeys)
}

func (*Index) ManagementKeySize() int {
	return int(keySlotBase - managementKeys)
}
