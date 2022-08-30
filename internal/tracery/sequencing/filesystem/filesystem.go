package filesystem

import (
	"encoding/hex"
	"os"

	"github.com/jmg292/G-Net/internal/tracery/sequencing/filesystem/entry"
	gnet "github.com/jmg292/G-Net/pkg/gneterrs"
)

type sequenceMapFile struct {
	path            string
	blockCount      uint64
	blockIdIndexMap map[string]uint64
}

func New(path string) *sequenceMapFile {
	return &sequenceMapFile{path: path}
}

func (*sequenceMapFile) blockIdToString(blockId []byte) string {
	return hex.EncodeToString(blockId)
}

func (*sequenceMapFile) validateManifestFile(handle *os.File) error {
	stat, err := handle.Stat()
	if err != nil {
		return err
	}
	if stat.Size()%entry.Size != 0 {
		return gnet.ErrorManifestInvalidSize
	}
	return nil
}

func (f *sequenceMapFile) getBlockCount(handle *os.File) (uint64, error) {
	if f.blockCount == 0 {
		stat, err := handle.Stat()
		if err != nil {
			return 0, err
		}
		f.blockCount = uint64(stat.Size() / entry.Size)
	}
	return f.blockCount, nil
}

func (f *sequenceMapFile) getEntryAtIndex(handle *os.File, blockIndex uint64) (*entry.FileEntry, error) {
	blockCount, err := f.getBlockCount(handle)
	if err != nil {
		return nil, err
	}
	if blockIndex > blockCount {
		return nil, gnet.ErrorBlockIndexOutOfRange
	}
	entryBytes := make([]byte, entry.Size)
	handle.Seek(int64(blockIndex)*entry.Size, 0)
	if _, err = handle.Read(entryBytes); err != nil {
		return nil, err
	}
	return entry.New(entryBytes)
}

func (f *sequenceMapFile) loadBlockIdIndexMap(handle *os.File) error {
	totalBlockCount, err := f.getBlockCount(handle)
	if err != nil {
		return err
	}
	f.blockIdIndexMap = make(map[string]uint64, totalBlockCount)
	for i := uint64(0); i < totalBlockCount; i++ {
		mapEntry, err := f.getEntryAtIndex(handle, i)
		if err != nil {
			return err
		}
		f.blockIdIndexMap[mapEntry.BlockIdString()] = i
	}
	return nil
}
