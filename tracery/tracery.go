package tracery

import (
	"gnet/management"
	"gnet/tracery/wumbo"
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
