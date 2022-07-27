package publication

import "github.com/jmg292/G-Net/internal/datagrams/management"

type NetworkInformation struct {
	PublicIPv4  []string
	PublicIPv6  []string
	PrivateIPv4 []string
	PrivateIPv6 []string
}

func (*NetworkInformation) Type() management.Datagram {
	return management.NetworkInformation
}
