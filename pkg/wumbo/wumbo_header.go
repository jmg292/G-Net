package wumbo

import (
	"fmt"
	"time"

	"github.com/jmg292/G-Net/internal/datagrams"
	"github.com/jmg292/G-Net/internal/utilities/convert"
	"github.com/jmg292/G-Net/pkg/identity/public"
)

const WumboHeaderByteCount int = 82

type WumboHeader struct {
	PrecedingBlockDigest []byte
	IssuerFingerprint    []byte
	CreationTime         time.Time
	BlockType            datagrams.Type
	ContentLength        uint32
	SignatureLength      uint32
}

func (header *WumboHeader) AsBytes() []byte {
	headerBytes := append(header.PrecedingBlockDigest, header.IssuerFingerprint...)
	headerBytes = append(headerBytes, convert.UInt64ToBytes(uint64(header.CreationTime.UnixMilli()))...)
	headerBytes = append(headerBytes, convert.UInt16ToBytes(uint16(header.BlockType))...)
	headerBytes = append(headerBytes, convert.UInt32ToBytes(header.ContentLength)...)
	headerBytes = append(headerBytes, convert.UInt32ToBytes(header.SignatureLength)...)
	return headerBytes
}

func ReadHeaderFromBlock(data []byte) (*WumboHeader, error) {
	if data == nil || len(data) < WumboHeaderByteCount {
		return nil, fmt.Errorf("invalid WUMBO data supplied")
	}
	return &WumboHeader{
		PrecedingBlockDigest: data[:32],
		IssuerFingerprint:    data[32:64],
		CreationTime:         time.UnixMilli(int64(convert.BytesToUInt64(data[64:72]))),
		BlockType:            datagrams.Type(convert.BytesToUInt16(data[72:74])),
		ContentLength:        convert.BytesToUInt32(data[74:78]),
		SignatureLength:      convert.BytesToUInt32(data[78:82]),
	}, nil
}

func NewWumboHeader(precedingBlockId []byte, blockType datagrams.Type, contentLength int, issuer any) *WumboHeader {
	header := WumboHeader{
		PrecedingBlockDigest: precedingBlockId,
		IssuerFingerprint:    issuer.(public.KeyRing).Fingerprint(),
		CreationTime:         time.Now(),
		BlockType:            blockType,
		ContentLength:        uint32(contentLength),
	}
	return &header
}
