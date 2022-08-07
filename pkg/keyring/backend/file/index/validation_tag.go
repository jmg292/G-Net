package index

func (i *Index) ValidationTagOffset() int {
	return i.DeviceCertificateOffset() + i.DeviceCertificateSize()
}

func (i *Index) ValidationTagSize(fileSize int) int {
	return fileSize - i.ValidationTagOffset()
}
