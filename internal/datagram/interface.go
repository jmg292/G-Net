package datagram

type Opaque interface {
	Type() Type
	Marshal() []byte
	Unmarshal(DatagramContent) error
}

type DatagramContent interface {
	Type() ContentType
}
