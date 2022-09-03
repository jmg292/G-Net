package linklayer

import (
	"github.com/jmg292/G-Net/internal/datagram/management/edict"
	gnet "github.com/jmg292/G-Net/pkg/gneterrs"
	"github.com/suborbital/grav/grav"
)

type Link struct {
	name  string
	rules []*edict.NetworkRoute
}

func (i *Link) Name() string {
	return i.name
}

func (i *Link) Configure(routes []*edict.NetworkRoute) error {
	return gnet.ErrorNotYetImplemented
}

func (i *Link) ModifyConfiguration(route []*edict.NetworkRoute) (err error) {
	return gnet.ErrorNotYetImplemented
}

func (i *Link) Handle(packet *edict.NetworkRoute) (err error) {
	return gnet.ErrorNotYetImplemented
}

func (i *Link) Receive(msg grav.Message) error {
	return gnet.ErrorNotYetImplemented
}
