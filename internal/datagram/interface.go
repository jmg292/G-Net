package datagram

type Opaque interface {
	Type() string
	Data() []byte
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

type Content interface {
	UnmarshalData(any) error
}
