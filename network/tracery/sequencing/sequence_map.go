package sequencing

type SequenceMap interface {
	GetIndexFromId([]byte) (uint64, error)
	GetIdFromIndex(uint64) ([]byte, error)
}
