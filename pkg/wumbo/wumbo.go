package wumbo

import (
	"github.com/jmg292/G-Net/internal/datagram"
	"github.com/jmg292/G-Net/pkg/pki/private"
	"github.com/jmg292/G-Net/pkg/wumbo/header"
)

func New(precedingBlockId []byte, data any, issuer any) (*Block, error) {
	contentType := data.(datagram.Datagram).Type()
	blockContent, err := MarshalContent(data)
	if err != nil {
		return nil, err
	}
	newBlock := Block{
		Header:  header.New(precedingBlockId, contentType, len(blockContent), issuer),
		Content: blockContent,
	}
	if newBlock.Signature, err = issuer.(private.KeyRing).Sign(newBlock.Digest()); err != nil {
		return nil, err
	}
	return &newBlock, nil
}
