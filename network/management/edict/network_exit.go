package edict

import "github.com/jmg292/G-Net/network/management"

type NetworkExit struct {
	PeerFingerprint []byte
}

func (*NetworkExit) PacketType() management.PacketType {
	return management.NetworkExitEdict
}
