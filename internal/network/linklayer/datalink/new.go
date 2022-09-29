package datalink

import (
	"github.com/jmg292/G-Net/internal/datagram/management/edict"
	"github.com/jmg292/G-Net/internal/network/linklayer/packet"
	"github.com/jmg292/G-Net/internal/network/linklayer/rules"
)

type ConnectionDirection uint8

const (
	ConnectionInbound = iota
	ConnectionOutbound
)

func configureEngine(routes []*edict.NetworkRoute, direction ConnectionDirection) (engine *rules.Engine, err error) {
	engine = &rules.Engine{}
	for _, route := range routes {
		if direction == ConnectionInbound {
			if err = engine.Configure(route.From); err != nil {
				break
			}
		} else {
			if err = engine.Configure(route.To); err != nil {
				break
			}
		}
	}
	return
}

func New(routes []*edict.NetworkRoute, direction ConnectionDirection, callback func(*packet.DataFrame, ConnectionDirection) error) (link Connection, err error) {
	if engine, e := configureEngine(routes, direction); e != nil {
		err = e
	} else {
		if direction == ConnectionInbound {
			link = &Inbound{rulesengine: engine, callback: callback}
		} else {
			link = &Outbound{rulesengine: engine, callback: callback}
		}
	}
	return
}
