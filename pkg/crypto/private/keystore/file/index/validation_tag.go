package index

func (i *index) ValidationTagOffset() int {
	return i.DeviceCertificateOffset() + i.DeviceCertificateSize()
}

func (i *index) ValidationTagSize(fileSize int) int {
	return fileSize - i.ValidationTagOffset()
}
