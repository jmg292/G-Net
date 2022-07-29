package header

import (
	"fmt"
	"time"

	"github.com/jmg292/G-Net/internal/datagram"
	"github.com/jmg292/G-Net/internal/utilities/convert"
	"github.com/jmg292/G-Net/pkg/gnet"
)

func (h *header) ToBytes() []byte {
	headerBytes := convert.UInt64ToBytes(uint64(h.CreationTime.UnixMilli()))
	headerBytes = append(headerBytes, h.SenderId...)
	headerBytes = append(headerBytes, h.EncryptionKey...)
	headerBytes = append(headerBytes, h.Nonce...)
	return append(headerBytes, []byte(string(h.Type))...)
}

func FromBytes(data []byte) (*header, error) {
	if len(data) < 84 {
		return nil, fmt.Errorf(string(gnet.ErrorInvalidHeader))
	}
	return &header{
		CreationTime:  time.UnixMilli(int64(convert.BytesToUInt64(data[:8]))),
		SenderId:      data[8:40],
		EncryptionKey: data[40:72],
		Nonce:         data[72:84],
		Type:          datagram.ContentType(string(data[84:])),
	}, nil
}
