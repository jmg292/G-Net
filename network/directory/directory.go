package directory

import (
	"github.com/jmg292/G-Net/network/tracery"
)

type directory struct {
	tracery tracery.Tracery
}

func NewNetworkDirectory(v any) *directory {
	return &directory{tracery: v.(tracery.Tracery)}
}
