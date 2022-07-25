package tracery

import (
	"github.com/jmg292/G-Net/management"
	"github.com/jmg292/G-Net/tracery/wumbo"
)

type Tracery interface {
	Open() error
	CreateNew(management.Fountainhead) error
	GetCurrentState() []byte
	GetRootBlock() wumbo.Block
	GetCurrentBlock() wumbo.Block
	GetBlockById([]byte) wumbo.Block
	AppendBlock(wumbo.Block) error
}
