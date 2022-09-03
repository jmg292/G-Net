package authentication

import (
	"github.com/jmg292/G-Net/internal/datagram"
	gnet "github.com/jmg292/G-Net/pkg/gneterrs"
	"github.com/jmg292/G-Net/pkg/keyring"
)

// authentication.Service
type Service struct {
	keyring *keyring.HardwareKeyRing
}

func (auth *Service) Register(callback func(datagram.Sealed) error) error {
	return gnet.ErrorNotYetImplemented
}

func (auth *Service) Inbound(msg datagram.Sealed) error {
	return gnet.ErrorNotYetImplemented
}

func (auth *Service) Outbound(msg datagram.Sealed) error {
	return gnet.ErrorNotYetImplemented
}
