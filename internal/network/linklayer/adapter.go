package linklayer

import (
	"github.com/jmg292/G-Net/internal/datagram/management/edict"
	"github.com/jmg292/G-Net/internal/network/linklayer/authentication"
	"github.com/jmg292/G-Net/internal/network/linklayer/datalink"
	"github.com/jmg292/G-Net/internal/network/linklayer/packet"
	"github.com/jmg292/G-Net/pkg/keyring"
	"github.com/suborbital/grav/grav"
)

// Adapter is binding layer between linklayer/authentication.Service,
// linklayer/datalink.Connections (in and out), and a grav.Pod,
// which is used to connect a service to the service mesh
type Adapter struct {
	in      datalink.Connection
	out     datalink.Connection
	auth    *authentication.Service
	pod     *grav.Pod
	timeout int
}

// NewAdapter returns a new *Adapter with inbound and outbound datalinks
// connected to the provided pod, each with rules engines configured using
// the provided routes, and an authentication.Service configured using the
// provided hsm.  The timeout parameter determines the amount of time an adapter
// will wait for a synchronous reply to an outbound message.
func NewAdapter(pod *grav.Pod, routes []*edict.NetworkRoute, hsm keyring.HardwareKeyRing, timeout int) (adapter *Adapter, err error) {
	adapter = &Adapter{
		auth:    &authentication.Service{Keyring: hsm},
		pod:     pod,
		timeout: timeout,
	}
	if adapter.in, err = datalink.New(routes, datalink.ConnectionInbound, adapter.handle); err != nil {
		adapter = nil
	} else if adapter.out, err = datalink.New(routes, datalink.ConnectionOutbound, adapter.handle); err != nil {
		adapter = nil
	}
	return
}

// Handle provides a callback layer to allow a datalink.Outbound to transmit
// packet.DataFrames to the rest of the service mesh using a grav.Pod. It's
// also used by a datalink.Inbound to relay packet.DataFrames to the adapter's
// attached authentication.Service for further processing before passing the
// datagram through the service interface.
func (a *Adapter) handle(packet *packet.DataFrame, direction datalink.ConnectionDirection) (err error) {
	if direction == datalink.ConnectionInbound {
		err = a.auth.HandleFrame(packet)
	} else {
		err = a.pod.Send(packet).WaitUntil(grav.Timeout(a.timeout), a.in.(datalink.Inbound).Receive)
	}
	return
}
