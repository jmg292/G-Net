package payload

import "time"

type NetworkState struct {
	CurrentStateDigest []byte
	CurrentBlockId     []byte
	PrecedingBlockID   []byte
	LastModifiedTime   time.Time
	CurrentTime        time.Time
}

func (*NetworkState) Type() ContentType {
	return ContentType_NetworkState
}
