package datalink

import (
	"fmt"

	"github.com/jmg292/G-Net/internal/datagram/management/edict"
	"github.com/jmg292/G-Net/internal/host"
	"github.com/jmg292/G-Net/internal/network/linklayer/packet"
	"github.com/jmg292/G-Net/internal/network/linklayer/rules"
	gnet "github.com/jmg292/G-Net/pkg/gneterrs"
	"github.com/suborbital/grav/grav"
)

// Inbound is a unidirectional inbound Connection for Adapter
type Inbound struct {
	name        string
	rulesengine *rules.Engine
	callback    func(*packet.DataFrame, ConnectionDirection) error
}

// Name implements Connection for Inbound
func (i Inbound) Name() string {
	return i.name
}

// Direction implements Connection for Inbound
func (Inbound) Direction() ConnectionDirection {
	return ConnectionInbound
}

// ModifyConfiguration implements Connection for Inbound
func (i Inbound) ModifyConfiguration(route *edict.NetworkRoute) (err error) {
	return i.rulesengine.Configure(route.From)
}

// Handle receives a *packet.Dataframe through Inbound.Receive from the
// service mesh.  It then uses its rules.Engine to determine if the
// *packet.DataFrame is allowed to pass to the authentication.Service
// for further evaluation.  Any errors returned are modified by Inbound.Receive
// before being returned to the sender.
func (i Inbound) Handle(packet *packet.DataFrame) (err error) {
	if evaluation, e := i.rulesengine.Evaluate(packet); e != nil {
		err = gnet.ErrorEvalFailed
		host.GetHost().Log(fmt.Sprintf(gnet.ErrorEvalFailed.Error(), "Link IN", e), host.Warn)
	} else if evaluation&rules.AllowedIn != 0 {
		if err = i.callback(packet, i.Direction()); err != nil {
			host.GetHost().Log(fmt.Sprintf(gnet.ErrorHandleFailed.Error(), "Link IN", e), host.Error)
		}
	}
	return nil
}

// Receive is registered as a grav.Pod message handling callback.  It receives
// a new grav.Messages from the service mesh and attempts to cast that
// grav.Message to a *packet.DataFrame. If the cast is successful, Receive
// passes the *packet.DataFrame to Inbound.Handle for further processing.
// Any errors returned are logged, but replaced with a generic error that
// wraps ErrDropped prior to returning that error to the sender.
func (i Inbound) Receive(msg grav.Message) (err error) {
	if dataframe, ok := msg.(*packet.DataFrame); !ok {
		host.GetHost().Log(gnet.ErrorInvalidDataframe.Error(), host.Warn)
		err = gnet.ErrorDropped
	} else if e := i.Handle(dataframe); e != nil {
		err = gnet.ErrorDropped
	}
	return
}
