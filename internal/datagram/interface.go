package datagram

type Opaque interface {
	Type() Type
}

type DatagramContent interface {
	Type() ContentType
}
