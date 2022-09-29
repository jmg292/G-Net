package mesh

import (
	"github.com/jmg292/G-Net/internal/datagram"
	"github.com/jmg292/G-Net/pkg/identity/certificate"
)

type Service interface {
	Handle(datagram.Sealed) error
	Register(func(datagram.Opaque) error) error
	Send(datagram.Opaque, string, certificate.Identity) error
}
