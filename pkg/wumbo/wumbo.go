package wumbo

import (
	"crypto/rand"

	"github.com/jmg292/G-Net/internal/datagram"
	"github.com/jmg292/G-Net/pkg/keyring"
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
	if newBlock.Signature, err = issuer.(keyring.PrivateKeyRing).Sign(rand.Reader, newBlock.Digest(), nil); err != nil {
		return nil, err
	}
	return &newBlock, nil
}
