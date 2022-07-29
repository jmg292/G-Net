package filesystem

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/jmg292/G-Net/pkg/gnet"
)

const entrySize int64 = 40

type sequenceMapFile struct {
	path            string
	blockIdIndexMap map[string]uint64
}

func (*sequenceMapFile) blockIdToString(blockId []byte) string {
	return hex.EncodeToString(blockId)
}

func (*sequenceMapFile) stringToBlockId(blockId string) ([]byte, error) {
	return hex.DecodeString(blockId)
}

func (*sequenceMapFile) validateManifestFile(handle *os.File) error {
	stat, err := handle.Stat()
	if err != nil {
		return err
	}
	if stat.Size()%entrySize != 0 {
		return fmt.Errorf(string(gnet.ErrorManifestInvalidSize))
	}
	return nil
}

func (*sequenceMapFile) getBlockCount(handle *os.File) (uint64, error) {
	stat, err := handle.Stat()
	if err != nil {
		return 0, err
	}
	return uint64(stat.Size() / entrySize), nil
}

func (f *sequenceMapFile) loadBlockIdIndexMap(handle *os.File) error {
	reader := bufio.NewReader(handle)
	entryBytes := make([]byte, entrySize)
	readCount, err := reader.Read(entryBytes)
	for readCount {

	}
}
