package filesystem

import (
	"fmt"
	"os"

	"github.com/jmg292/G-Net/pkg/gnet"
)

func (f *sequenceMapFile) Open() error {
	var handle *os.File
	if _, err := os.Stat(f.path); os.IsNotExist(err) {
		if handle, err = os.Create(f.path); err != nil {
			return err
		}
	} else {
		if handle, err = os.Open(f.path); err != nil {
			return err
		}
	}
	defer handle.Close()
	if err := f.validateManifestFile(handle); err != nil {
		return err
	}
	return f.loadBlockIdIndexMap(handle)
}

func (f *sequenceMapFile) Close() error {
	for k := range f.blockIdIndexMap {
		delete(f.blockIdIndexMap, k)
	}
	f.blockIdIndexMap = nil
	f.blockCount = 0
	return nil
}

func (f *sequenceMapFile) BlockCount() (uint64, error) {
	var blockCount uint64 = 0
	handle, err := os.Open(f.path)
	if err != nil {
		return blockCount, err
	}
	defer handle.Close()
	return f.getBlockCount(handle)
}

func (f *sequenceMapFile) PutBlockId(blockId []byte) error {
	if _, ok := f.blockIdIndexMap[f.blockIdToString(blockId)]; ok {
		return fmt.Errorf(string(gnet.ErrorBlockExists))
	}
	return fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *sequenceMapFile) GetBlockIdFromIndex(index uint64) ([]byte, error) {
	handle, err := os.Open(f.path)
	if err != nil {
		return nil, err
	}
	defer handle.Close()
	entry, err := f.getEntryAtIndex(handle, index)
	if err != nil {
		return nil, err
	}
	return entry.BlockId, nil
}
