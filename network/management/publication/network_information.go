package publication

import "github.com/jmg292/G-Net/network/management"

type NetworkInformation struct {
	PublicIPv4  []string
	PublicIPv6  []string
	PrivateIPv4 []string
	PrivateIPv6 []string
}

func (*NetworkInformation) PacketType() management.PacketType {
	return management.NetworkInfoPublication
}
