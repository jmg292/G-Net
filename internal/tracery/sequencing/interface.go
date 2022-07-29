package sequencing

type SequenceMap interface {
	Open() error
	Close() error
	BlockCount() (uint64, error)
	PutBlockId([]byte) (uint64, error)
	GetBlockIdFromIndex(uint64) ([]byte, error)
	GetIndexFromBlockId([]byte) (uint64, error)
}
