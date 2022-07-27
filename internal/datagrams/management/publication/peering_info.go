package publication

import "github.com/jmg292/G-Net/internal/datagrams/management"

type PeeringInformation struct {
	ServiceAddress string
	PortNumber     int
	Certificate    []byte
}

func (*PeeringInformation) Type() management.Datagram {
	return management.PeeringInformation
}
