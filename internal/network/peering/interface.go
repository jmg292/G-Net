package peering

import (
	"github.com/jmg292/G-Net/internal/datagram"
	"github.com/jmg292/G-Net/internal/datagram/management/publication"
)

type Connection interface {
	SendTo(datagram.Sealed, publication.PeeringInfo)
	RegisterHandler(func(datagram.Sealed))
}
