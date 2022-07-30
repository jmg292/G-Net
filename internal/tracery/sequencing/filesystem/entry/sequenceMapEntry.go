package entry

import (
	"encoding/hex"
	"fmt"

	"github.com/jmg292/G-Net/internal/utilities/convert"
	"github.com/jmg292/G-Net/pkg/gnet"
)

const Size int64 = 40

type FileEntry struct {
	BlockId    []byte
	BlockIndex uint64
}

func (f *FileEntry) ToBytes() []byte {
	return append(convert.UInt64ToBytes(f.BlockIndex), f.BlockId...)
}

func (f *FileEntry) BlockIdString() string {
	return hex.EncodeToString(f.BlockId)
}

func New(entryBytes []byte) (*FileEntry, error) {
	if len(entryBytes) != int(Size) {
		return nil, fmt.Errorf(string(gnet.ErrorMalformedEntry))
	}
	return &FileEntry{
		BlockIndex: convert.BytesToUInt64(entryBytes[:8]),
		BlockId:    entryBytes[8:],
	}, nil
}
