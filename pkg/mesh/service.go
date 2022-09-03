package mesh

import "github.com/jmg292/G-Net/internal/datagram"

type Service interface {
	Handle(datagram.Sealed) error
	Register(func(datagram.Opaque) error) error
	Send(datagram.Opaque) error
}
