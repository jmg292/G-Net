package publication

import "github.com/jmg292/G-Net/internal/datagrams/management"

type PeeringInfo struct {
	ServiceAddress string
	PortNumber     int
	Certificate    []byte
}

func (*PeeringInfo) Type() management.Datagram {
	return management.PeeringInformation
}
