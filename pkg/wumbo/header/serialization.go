package header

import (
	"fmt"
	"time"

	"github.com/jmg292/G-Net/internal/datagrams"
	"github.com/jmg292/G-Net/internal/utilities/convert"
)

func (header *Header) ToBytes() []byte {
	headerBytes := append(header.PrecedingBlockDigest, header.IssuerFingerprint...)
	headerBytes = append(headerBytes, convert.UInt64ToBytes(uint64(header.CreationTime.UnixMilli()))...)
	headerBytes = append(headerBytes, convert.UInt16ToBytes(uint16(header.BlockType))...)
	headerBytes = append(headerBytes, convert.UInt32ToBytes(header.ContentLength)...)
	headerBytes = append(headerBytes, convert.UInt32ToBytes(header.SignatureLength)...)
	return headerBytes
}

func FromBytes(data []byte) (*Header, error) {
	if data == nil || len(data) < ByteCount {
		return nil, fmt.Errorf("invalid WUMBO data supplied")
	}
	return &Header{
		PrecedingBlockDigest: data[:32],
		IssuerFingerprint:    data[32:64],
		CreationTime:         time.UnixMilli(int64(convert.BytesToUInt64(data[64:72]))),
		BlockType:            datagrams.Type(convert.BytesToUInt16(data[72:74])),
		ContentLength:        convert.BytesToUInt32(data[74:78]),
		SignatureLength:      convert.BytesToUInt32(data[78:82]),
	}, nil
}
