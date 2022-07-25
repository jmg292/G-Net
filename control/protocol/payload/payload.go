package payload

type ContentType uint16

const (
	ContentType_NetworkState ContentType = iota
)

type Content interface {
	Type() ContentType
}
