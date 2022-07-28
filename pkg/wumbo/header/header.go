package header

import (
	"time"

	"github.com/jmg292/G-Net/internal/datagrams"
	"github.com/jmg292/G-Net/pkg/identity/public"
)

const ByteCount int = 82

type Header struct {
	PrecedingBlockDigest []byte
	IssuerFingerprint    []byte
	CreationTime         time.Time
	BlockType            datagrams.Type
	ContentLength        uint32
	SignatureLength      uint32
}

func New(precedingBlockId []byte, blockType datagrams.Type, contentLength int, issuer any) *Header {
	Header := Header{
		PrecedingBlockDigest: precedingBlockId,
		IssuerFingerprint:    issuer.(public.KeyRing).Fingerprint(),
		CreationTime:         time.Now(),
		BlockType:            blockType,
		ContentLength:        uint32(contentLength),
	}
	return &Header
}
