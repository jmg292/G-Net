package keyslot

import (
	"fmt"

	"github.com/jmg292/G-Net/pkg/gnet"
)

func (slot *keySlot) getKey() []byte {
	return slot[keyOffset:maxKeySize]
}

func (slot *keySlot) getKeyWithOverhead() []byte {
	return slot[keyOffset:]
}

func (slot *keySlot) setKey(key []byte) (err error) {
	if len(key) > maxKeySize {
		err = fmt.Errorf(string(gnet.ErrorInvalidContentLength))
	} else {
		copy(slot[keyOffset:maxKeySize], key)
	}
	return
}
