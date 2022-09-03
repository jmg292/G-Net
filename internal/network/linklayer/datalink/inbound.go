package datalink

import (
	"github.com/jmg292/G-Net/internal/datagram/management/edict"
	gnet "github.com/jmg292/G-Net/pkg/gneterrs"
	"github.com/suborbital/grav/grav"
)

type Inbound struct {
	name  string
	rules []*edict.NetworkRoute
}

func (i *Inbound) Name() string {
	return i.name
}

func (i *Inbound) Configure(routes []*edict.NetworkRoute) error {
	return gnet.ErrorNotYetImplemented
}

func (i *Inbound) ModifyConfiguration(route []*edict.NetworkRoute) (err error) {
	return gnet.ErrorNotYetImplemented
}

func (i *Inbound) Handle(packet []byte) (err error) {
	return gnet.ErrorNotYetImplemented
}

func (i *Inbound) Receive(msg grav.Message) error {
	return gnet.ErrorNotYetImplemented
}
