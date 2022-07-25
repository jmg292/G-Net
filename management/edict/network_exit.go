package edict

import "gnet/management"

type NetworkExit struct {
	PeerFingerprint []byte
}

func (*NetworkExit) PacketType() management.PacketType {
	return management.NetworkExitEdict
}
