package payload

type ContentType string

const (
	NetworkState ContentType = "NetworkState"
	BlockObject              = "BlockObject"
)

type Content interface {
	Type() ContentType
}
