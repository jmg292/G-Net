package datalink

import (
	"github.com/jmg292/G-Net/internal/datagram"
	"github.com/jmg292/G-Net/internal/datagram/management/edict"
	gnet "github.com/jmg292/G-Net/pkg/gneterrs"
)

type Outbound struct {
	name  string
	rules []*edict.NetworkRoute
}

func (o *Outbound) Name() string {
	return o.name
}

func (o *Outbound) Configure(routes *[]edict.NetworkRoute) error {
	return gnet.ErrorNotYetImplemented
}

func (o *Outbound) ModifyConfiguration(route *edict.NetworkRoute) error {
	return gnet.ErrorNotYetImplemented
}

func (o *Outbound) Handle(packet *datagram.Sealed) error {
	return gnet.ErrorNotYetImplemented
}
