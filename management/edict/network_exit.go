package edict

import "github.com/jmg292/G-Net/management"

type NetworkExit struct {
	PeerFingerprint []byte
}

func (*NetworkExit) PacketType() management.PacketType {
	return management.NetworkExitEdict
}
