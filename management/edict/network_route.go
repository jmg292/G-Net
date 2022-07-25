package edict

import "gnet/management"

type NetworkRoute struct {
	From   string
	To     string
	Weight int
}

func (*NetworkRoute) PacketType() management.PacketType {
	return management.NetworkRouteEdict
}
