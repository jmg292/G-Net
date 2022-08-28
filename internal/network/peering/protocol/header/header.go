package header

import (
	"crypto/rand"
	"time"

	"github.com/jmg292/G-Net/internal/datagram"
	"github.com/jmg292/G-Net/internal/network/peering/protocol"
)

type header struct {
	CreationTime  time.Time
	SenderId      []byte
	EncryptionKey []byte
	Nonce         []byte
	Type          datagram.ContentType
}

func New(senderId []byte) *header {
	h := header{
		CreationTime:  time.Now(),
		SenderId:      senderId,
		EncryptionKey: make([]byte, protocol.KeySize),
		Nonce:         make([]byte, protocol.NonceSize),
	}
	rand.Read(h.EncryptionKey)
	rand.Read(h.Nonce)
	return &h
}
