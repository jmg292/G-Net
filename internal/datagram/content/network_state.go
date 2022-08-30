package content

import (
	"time"

	"github.com/jmg292/G-Net/internal/datagram"
)

type NetworkState struct {
	CurrentStateDigest []byte
	CurrentBlockId     []byte
	PrecedingBlockID   []byte
	LastModifiedTime   time.Time
	CurrentTime        time.Time
}

func (*NetworkState) Type() datagram.ContentType {
	return NetworkStateContent
}
