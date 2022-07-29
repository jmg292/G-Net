package filesystem

import "os"

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
