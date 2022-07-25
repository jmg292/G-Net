package message

import (
	"gnet/control/protocol/payload"
	"time"
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
