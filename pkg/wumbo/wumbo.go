package wumbo

import (
	"github.com/jmg292/G-Net/internal/datagrams"
	"github.com/jmg292/G-Net/pkg/identity/private"
	"github.com/jmg292/G-Net/pkg/wumbo/header"
)

func New(precedingBlockId []byte, data any, issuer any) (*Block, error) {
	contentType := data.(datagrams.Datagram).Type()
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
