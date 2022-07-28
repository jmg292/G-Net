package content

import "github.com/jmg292/G-Net/internal/datagram"

type BlockObject struct {
	EncodedContent string
}

func (*BlockObject) Type() datagram.ContentType {
	return BlockObjectContent
}
