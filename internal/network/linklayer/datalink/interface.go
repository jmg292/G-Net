package datalink

import (
	"github.com/jmg292/G-Net/internal/datagram/management/edict"
	"github.com/jmg292/G-Net/internal/network/linklayer/packet"
)

type Connection interface {
	Name() string
	Direction() ConnectionDirection
	ModifyConfiguration(*edict.NetworkRoute) error
	Handle(*packet.DataFrame) error
}
