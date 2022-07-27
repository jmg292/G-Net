package edict

import "github.com/jmg292/G-Net/internal/datagrams/management"

type NetworkExit struct {
	PeerFingerprint []byte
}

func (*NetworkExit) PacketType() management.Datagram {
	return management.ExitNodeDeclaration
}
