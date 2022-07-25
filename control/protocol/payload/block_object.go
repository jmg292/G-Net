package payload

type BlockObjectContent struct {
	EncodedContent string
}

func (*BlockObjectContent) Type() ContentType {
	return BlockObject
}
