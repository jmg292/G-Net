package management

type PacketType uint16

const (
	NetworkCreation PacketType = iota
	AdminWarrant
	DeviceWarrant
	NetworkInfoPublication
	PeeringInfoPublication
	NetworkRouteEdict
	NetworkExitEdict
)
