package authentication

import (
	"github.com/jmg292/G-Net/internal/datagram"
	"github.com/jmg292/G-Net/internal/network/linklayer/packet"
	gnet "github.com/jmg292/G-Net/pkg/gneterrs"
	"github.com/jmg292/G-Net/pkg/identity/certificate"
	"github.com/jmg292/G-Net/pkg/keyring"
)

// authentication.Service
type Service struct {
	Keyring keyring.HardwareKeyRing
}

func (auth *Service) HandleFrame(msg *packet.DataFrame) (err error) {

}

func (auth *Service) Register(callback func(datagram.Opaque) error) error {
	return gnet.ErrorNotYetImplemented
}

func (auth *Service) Handle(msg datagram.Sealed) error {
	return gnet.ErrorNotYetImplemented
}

func (auth *Service) Send(msg datagram.Opaque, svcDomain string, peer certificate.Identity) error {
	return gnet.ErrorNotYetImplemented
}
