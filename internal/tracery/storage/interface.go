package storage

import "github.com/jmg292/G-Net/pkg/wumbo"

type BlockStorage interface {
	Open() error
	Close() error
	PutBlock(wumbo.Block) error
	GetBlock([]byte) (wumbo.Block, error)
}
