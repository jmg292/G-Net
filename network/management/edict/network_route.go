package edict

import "github.com/jmg292/G-Net/network/management"

type NetworkRoute struct {
	From   string
	To     string
	Weight int
}

func (*NetworkRoute) PacketType() management.PacketType {
	return management.NetworkRouteEdict
}
