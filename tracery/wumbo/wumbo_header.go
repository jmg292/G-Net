package wumbo

import (
	"fmt"
	"gnet/identity"
	"gnet/management"
	"gnet/utilities/convert"
	"time"
)

const WumboHeaderByteCount int = 82

type WumboHeader struct {
	PrecedingBlockDigest []byte
	IssuerFingerprint    []byte
	CreationTime         time.Time
	BlockType            management.PacketType
	ContentLength        uint32
	SignatureLength      uint32
}

func createHeaderFromBlock(data []byte) (*WumboHeader, error) {
	if data == nil || len(data) < WumboHeaderByteCount {
		return nil, fmt.Errorf("invalid WUMBO data supplied")
	}
	return &WumboHeader{
		PrecedingBlockDigest: data[:32],
		IssuerFingerprint:    data[32:64],
		CreationTime:         time.UnixMilli(int64(convert.UInt64FromBinary(data[64:72]))),
		BlockType:            management.PacketType(convert.UInt16FromBinary(data[72:74])),
		ContentLength:        convert.UInt32FromBinary(data[74:78]),
		SignatureLength:      convert.UInt32FromBinary(data[78:82]),
	}, nil
}

func NewWumboHeader(precedingBlockId []byte, blockType management.PacketType, contentLength int, issuer any) *WumboHeader {
	header := WumboHeader{
		PrecedingBlockDigest: precedingBlockId,
		IssuerFingerprint:    issuer.(identity.Key).Fingerprint(),
		CreationTime:         time.Now(),
		BlockType:            blockType,
		ContentLength:        uint32(contentLength),
	}
	return &header
}

func (header *WumboHeader) AsBytes() []byte {
	headerBytes := append(header.PrecedingBlockDigest, header.IssuerFingerprint...)
	headerBytes = append(headerBytes, convert.UInt64ToBinary(uint64(header.CreationTime.UnixMilli()))...)
	headerBytes = append(headerBytes, convert.UInt16ToBinary(uint16(header.BlockType))...)
	headerBytes = append(headerBytes, convert.UInt32ToBinary(header.ContentLength)...)
	headerBytes = append(headerBytes, convert.UInt32ToBinary(header.SignatureLength)...)
	return headerBytes
}
