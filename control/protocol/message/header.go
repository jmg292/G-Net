package message

import (
	"time"

	"github.com/jmg292/G-Net/control/protocol/payload"
)

type header struct {
	CreationTime  time.Time
	SenderId      []byte
	EncryptionKey []byte
	Nonce         []byte
	ContentType   payload.ContentType
}

func CreateHeader(senderId []byte, contentType payload.ContentType) *header {
	return &header{
		CreationTime:  time.Now(),
		SenderId:      senderId,
		ContentType:   contentType,
		EncryptionKey: make([]byte, KeySize),
		Nonce:         make([]byte, NonceSize),
	}
}
