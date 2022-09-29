package rules

import (
	"github.com/VirusTotal/gyp/ast"
	"github.com/jmg292/G-Net/internal/network/linklayer/packet"
	gnet "github.com/jmg292/G-Net/pkg/gneterrs"
)

type Result uint8

const (
	NotAllowed Result = 1 << iota // 00000001
	AllowedIn                     // 00000010
	AllowedOut                    // 00000100
)

// 00000110
const Allowed Result = AllowedIn | AllowedOut

// rules.Engine
type Engine struct {
	ruleset ast.RuleSet
}

func (e *Engine) Configure(rules string) (err error) {
	err = gnet.ErrorNotYetImplemented
	return
}

func (e *Engine) Evaluate(frame *packet.DataFrame) (eval Result, err error) {
	if frame == nil || frame.Sender == "" || frame.Recipient == "" {
		err = gnet.ErrorInvalidDatagram
	} else {
		eval = NotAllowed
	}
	return
}
