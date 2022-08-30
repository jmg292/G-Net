package header

import (
	"time"

	"github.com/jmg292/G-Net/internal/datagram"
	"github.com/jmg292/G-Net/pkg/keyring"
)

const ByteCount int = 82

type Header struct {
	PrecedingBlockDigest []byte
	IssuerFingerprint    []byte
	CreationTime         time.Time
	BlockType            datagram.Type
	ContentLength        uint32
	SignatureLength      uint32
}

func New(precedingBlockId []byte, blockType datagram.Type, contentLength int, issuer any) *Header {
	Header := Header{
		PrecedingBlockDigest: precedingBlockId,
		IssuerFingerprint:    issuer.(keyring.PublicKeyRing).Identifier(),
		CreationTime:         time.Now(),
		BlockType:            blockType,
		ContentLength:        uint32(contentLength),
	}
	return &Header
}
