package wumbo

import (
	"github.com/jmg292/G-Net/internal/datagrams"
	"github.com/jmg292/G-Net/pkg/identity/private"
	"github.com/jmg292/G-Net/pkg/wumbo/block"
	"github.com/jmg292/G-Net/pkg/wumbo/header"
)

type Header *header.Header
type Block *block.Block

func New(precedingBlockId []byte, data any, issuer any) (Block, error) {
	contentType := data.(datagrams.Datagram).Type()
	blockContent, err := MarshalContent(data)
	if err != nil {
		return nil, err
	}
	issuedBlock := block.Block{
		Header:  header.New(precedingBlockId, contentType, len(blockContent), issuer),
		Content: blockContent,
	}
	if issuedBlock.Signature, err = issuer.(private.KeyRing).Sign(issuedBlock.Digest()); err != nil {
		return nil, err
	}
	return &issuedBlock, nil
}
