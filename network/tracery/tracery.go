package tracery

import (
	"github.com/jmg292/G-Net/network/management"
	"github.com/jmg292/G-Net/network/tracery/manifest"
	"github.com/jmg292/G-Net/network/tracery/sequencing"
	"github.com/jmg292/G-Net/network/tracery/wumbo"
)

type Tracery interface {
	Open() error
	Manifest() manifest.Manifest
	SequenceMap() sequencing.SequenceMap
	CreateNew(management.Fountainhead) error
	AppendBlock(*wumbo.Block) error
	GetCurrentState() []byte
	GetCurrentIndex() uint64
	GetRootBlock() *wumbo.Block
	GetCurrentBlock() *wumbo.Block
	GetBlockById([]byte) (*wumbo.Block, error)
	GetBlockByIndex(uint64) (*wumbo.Block, error)
}
