package datalink

import (
	"github.com/jmg292/G-Net/internal/datagram/management/edict"
	"github.com/jmg292/G-Net/internal/network/linklayer/packet"
	"github.com/jmg292/G-Net/internal/network/linklayer/rules"
	gnet "github.com/jmg292/G-Net/pkg/gneterrs"
)

// Outbound is a unidirectional outbound Connection for Adapter
type Outbound struct {
	name        string
	rulesengine *rules.Engine
	callback    func(*packet.DataFrame, ConnectionDirection) error
}

// Name implements Connection for Outbound
func (o Outbound) Name() string {
	return o.name
}

// Direction implements Connection for Outbound
func (Outbound) Direction() ConnectionDirection {
	return ConnectionOutbound
}

// ModifyConfiguration implements Connection for Outbound
func (o Outbound) ModifyConfiguration(route *edict.NetworkRoute) error {
	return o.rulesengine.Configure(route.To)
}

// Handle receives a *packet.DataFrame from the authentication.Service and
// uses its rules.Engine to determine if the *packet.DataFrame is allowed
// to enter the service mesh.  If it is, Handle calls back to linklayer.Adapter
// to transmit the *packet.DataFrame on the service mesh. Otherwise,
// Outbound.Handle returns gnet.ErrorDropped.  Any errors are returned unmodified
// to the caller.
func (o Outbound) Handle(packet *packet.DataFrame) (err error) {
	if evaluation, e := o.rulesengine.Evaluate(packet); e != nil {
		err = e
	} else if evaluation&rules.AllowedOut == 0 {
		err = gnet.ErrorDropped
	} else {
		err = o.callback(packet, ConnectionOutbound)
	}
	return
}
