package wireguard

import (
	gnet "github.com/jmg292/G-Net/pkg/gneterrs"
	"golang.zx2c4.com/wireguard/tun"
	"golang.zx2c4.com/wireguard/windows/tunnel/winipcfg"
)

func (t *Tunnel) Name() (name string, err error) {
	if guid, e := winipcfg.LUID(t.device.(*tun.NativeTun).LUID()).GUID(); e != nil {
		err = e
	} else {
		name = guid.String()
	}
	return
}

func (t *Tunnel) Up() error {
	return gnet.ErrorNotYetImplemented
}

func (t *Tunnel) Down() error {
	return gnet.ErrorNotYetImplemented
}

func (t *Tunnel) Register(callback func([]byte) error) error {
	return gnet.ErrorNotYetImplemented
}

func (t *Tunnel) Send(data, address []byte) error {
	return gnet.ErrorNotYetImplemented
}
