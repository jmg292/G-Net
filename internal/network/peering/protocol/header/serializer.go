package header

import (
	"github.com/jmg292/G-Net/internal/utilities/convert"
)

func (h *header) ToBytes() []byte {
	headerBytes := convert.UInt64ToBytes(uint64(h.CreationTime.UnixMilli()))
	headerBytes = append(headerBytes, h.SenderId...)
	headerBytes = append(headerBytes, h.EncryptionKey...)
	headerBytes = append(headerBytes, h.Nonce...)
	return append(headerBytes, []byte(string(h.Type))...)
}
