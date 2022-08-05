package keyslot

import (
	"fmt"

	"github.com/jmg292/G-Net/pkg/gnet"
)

func (slot *keySlot) getKeyWithOverhead() []byte {
	return slot[keyOffset:]
}

func (slot *keySlot) GetKey() []byte {
	return slot[keyOffset:maxKeySize]
}

func (slot *keySlot) SetKey(key []byte) (err error) {
	if len(key) > maxKeySize {
		err = fmt.Errorf(string(gnet.ErrorInvalidContentLength))
	} else {
		copy(slot[keyOffset:maxKeySize], key)
	}
	return
}
