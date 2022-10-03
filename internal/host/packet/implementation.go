package packet

import gnet "github.com/jmg292/G-Net/pkg/gneterrs"

// Name implements network.Interface for packet.Filter
func (*Filter) Name() (string, error) {
	return "WinDivert Packet Filter", nil
}

// Up implements network.Interface for packet.Filter
func (f *Filter) Up() error {
	return gnet.ErrorNotYetImplemented
}

// Down implements network.Interface for packet.Filter
func (f *Filter) Down() error {
	return gnet.ErrorNotYetImplemented
}

// Register implements network.Interface for packet.Filter
func (f *Filter) Register(callback func([]byte) error) error {
	return gnet.ErrorNotYetImplemented
}

// Send implements network.Interface for packet.Filter
func (f *Filter) Send(data, address []byte) error {
	return gnet.ErrorNotYetImplemented
}
