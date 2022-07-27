package edict

import "github.com/jmg292/G-Net/internal/datagrams/management"

type NetworkRoute struct {
	From   string
	To     string
	Weight int
}

func (*NetworkRoute) PacketType() management.Datagram {
	return management.RoutingDeclaration
}
