package linklayer

import (
	"github.com/jmg292/G-Net/internal/datagram"
	"github.com/jmg292/G-Net/internal/datagram/management/edict"
)

type Link interface {
	Name() string
	Configure([]*edict.NetworkRoute) error
	ModifyConfiguration(*edict.NetworkRoute) error
	Handle(*datagram.Sealed) error
}
