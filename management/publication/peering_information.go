package publication

import "github.com/jmg292/G-Net/management"

type PeeringInformation struct {
	ServiceAddress string
	PortNumber     int
	Certificate    []byte
}

func (*PeeringInformation) PacketType() management.PacketType {
	return management.PeeringInfoPublication
}
