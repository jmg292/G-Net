package payload

import "time"

type NetworkStateContent struct {
	CurrentStateDigest []byte
	CurrentBlockId     []byte
	PrecedingBlockID   []byte
	LastModifiedTime   time.Time
	CurrentTime        time.Time
}

func (*NetworkStateContent) Type() ContentType {
	return NetworkState
}
