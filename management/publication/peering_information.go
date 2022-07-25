package publication

import "gnet/management"

type PeeringInformation struct {
	ServiceAddress string
	PortNumber     int
	Certificate    []byte
}

func (*PeeringInformation) PacketType() management.PacketType {
	return management.PeeringInfoPublication
}
