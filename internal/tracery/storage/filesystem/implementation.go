package filesystem

import (
	"bufio"
	"os"

	"github.com/jmg292/G-Net/pkg/wumbo"
)

func (storage *blockStorageDirectory) Open() error {
	return storage.lock()
}

func (storage *blockStorageDirectory) Close() error {
	return storage.unlock()
}

func (storage *blockStorageDirectory) PutBlock(block wumbo.Block) error {
	handle, err := os.Create(storage.blockPath(block.Digest()))
	if err != nil {
		return storage.error(err)
	}
	_, err = bufio.NewWriter(handle).Write(block.ToBytes())
	return err
}

func (storage *blockStorageDirectory) GetBlock(blockId []byte) (*wumbo.Block, error) {
	handle, err := os.Open(storage.blockPath(blockId))
	if err != nil {
		return nil, storage.error(err)
	}
	stat, err := handle.Stat()
	if err != nil {
		return nil, err
	}
	blockContent := make([]byte, stat.Size())
	if _, err = bufio.NewReader(handle).Read(blockContent); err != nil {
		return nil, err
	}
	return wumbo.NewFromBytes(blockContent)
}
