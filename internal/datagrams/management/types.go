package management

import (
	"github.com/jmg292/G-Net/internal/datagrams"
)

type Datagram datagrams.Type

const (
	NetworkRoot Datagram = iota
	AdminWarrant
	DeviceWarrant
	NetworkInformation
	PeeringInformation
	RoutingDeclaration
	ExitNodeDeclaration
)
