package linklayer

import (
	"github.com/jmg292/G-Net/internal/datagram"
	"github.com/jmg292/G-Net/internal/datagram/management/edict"
	gnet "github.com/jmg292/G-Net/pkg/gneterrs"
)

// output.Link
type Link struct {
	name  string
	rules []*edict.NetworkRoute
}

func (o *Link) Name() string {
	return o.name
}

func (o *Link) Configure(routes *[]edict.NetworkRoute) error {
	return gnet.ErrorNotYetImplemented
}

func (o *Link) ModifyConfiguration(route *edict.NetworkRoute) error {
	return gnet.ErrorNotYetImplemented
}

func (o *Link) Handle(packet *datagram.Sealed) error {
	return gnet.ErrorNotYetImplemented
}
