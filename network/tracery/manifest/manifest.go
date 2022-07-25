package manifest

type Manifest interface {
	GetAdminWarrantBlockIndex([]byte) (uint64, error)
	GetDeviceWarrantBlockIndex([]byte) (uint64, error)
}
