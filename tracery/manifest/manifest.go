package manifest

type Manifest interface {
	GetAdminWarrantBlockIndex([]byte) (int, error)
	GetDeviceWarrantBlockIndex([]byte) (int, error)
}
