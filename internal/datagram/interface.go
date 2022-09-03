package datagram

type Opaque interface {
	Type() Type
	Data() []byte
}

type Content interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
	UnmarshalData(any) error
}
