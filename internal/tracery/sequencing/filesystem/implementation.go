package filesystem

import (
	"fmt"
	"os"

	"github.com/jmg292/G-Net/internal/tracery/sequencing/filesystem/entry"
	"github.com/jmg292/G-Net/pkg/gnet"
)

func (f *sequenceMapFile) Open() error {
	handle, err := os.OpenFile(f.path, os.O_CREATE|os.O_RDONLY, 0600)
	if err != nil {
		return err
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
	handle, err := os.Open(f.path)
	if err != nil {
		return 0, err
	}
	defer handle.Close()
	return f.getBlockCount(handle)
}

func (f *sequenceMapFile) PutBlockId(blockId []byte) (uint64, error) {
	if _, ok := f.blockIdIndexMap[f.blockIdToString(blockId)]; ok {
		return 0, fmt.Errorf(string(gnet.ErrorBlockExists))
	}
	handle, err := os.OpenFile(f.path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return 0, err
	}
	defer handle.Close()
	currentCount, err := f.getBlockCount(handle)
	if err != nil {
		return 0, err
	}
	entry := entry.FileEntry{
		BlockId:    blockId,
		BlockIndex: currentCount + 1,
	}
	if _, err := handle.Write(entry.ToBytes()); err != nil {
		return 0, err
	}
	return entry.BlockIndex, nil
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

func (f *sequenceMapFile) GetIndexFromBlockId(blockId []byte) (uint64, error) {
	if index, ok := f.blockIdIndexMap[f.blockIdToString(blockId)]; ok {
		return index, nil
	}
	return 0, fmt.Errorf(string(gnet.ErrorBlockNotFound))
}
