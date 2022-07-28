package block

func (block *Block) AsBytes() []byte {
	blockBytes := append(block.Header.ToBytes(), block.Content...)
	return append(blockBytes, block.Signature...)
}

func ReadFromBytes(blockData []byte) (*Block, error) {
	return nil, nil
}
