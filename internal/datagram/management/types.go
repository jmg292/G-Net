package management

import (
	"github.com/jmg292/G-Net/internal/datagram"
)

type Datagram datagram.Type

const (
	NetworkRoot Datagram = iota
	AdminWarrant
	DeviceWarrant
	NetworkInformation
	PeeringInformation
	RoutingDeclaration
	ExitNodeDeclaration
)
