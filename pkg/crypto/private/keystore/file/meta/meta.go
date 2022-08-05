package meta

type sectionFlag byte

const (
	IndexSection sectionFlag = 1 << iota
	SaltSection
	AdminSlotSection
	SigningKeySlotSection
	AuthKeySlotSection
	DeviceKeySlotSection
	EncryptionKeySlotSection
	CertificateStoreSection
)

type Meta struct {
	path             string
	unlockedSections byte
	modifiedSections byte
}

func New(path string) *Meta {
	return &Meta{path: path}
}

func (m *Meta) Path() string {
	return m.path
}
