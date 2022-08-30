package datagram

type Datagram interface {
	Type() Type
}

type DatagramContent interface {
	Type() ContentType
}
