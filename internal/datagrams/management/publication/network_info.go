package publication

import "github.com/jmg292/G-Net/internal/datagrams/management"

type NetworkInfo struct {
	PublicIPv4  []string
	PublicIPv6  []string
	PrivateIPv4 []string
	PrivateIPv6 []string
}

func (*NetworkInfo) Type() management.Datagram {
	return management.NetworkInformation
}
