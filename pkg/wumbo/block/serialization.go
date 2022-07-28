package block

import "github.com/jmg292/G-Net/pkg/wumbo/header"

func (block *Block) ToBytes() []byte {
	blockBytes := append(block.Header.ToBytes(), block.Content...)
	return append(blockBytes, block.Signature...)
}

func FromBytes(blockData []byte) (*Block, error) {
	blockHeader, err := header.FromBytes(blockData)
	if err != nil {
		return nil, err
	}
	return &Block{
		Header:    blockHeader,
		Content:   blockData[header.ByteCount:blockHeader.ContentLength],
		Signature: blockData[header.ByteCount+int(blockHeader.ContentLength):],
	}, nil
}
