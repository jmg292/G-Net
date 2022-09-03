package datagram

type Opaque interface {
	Type() string
	Data() []byte
}

type Content interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
	UnmarshalData(any) error
}
