package slot

import (
	"github.com/jmg292/G-Net/pkg/gnet"
)

func (slot *KeySlot) getKeyWithOverhead() []byte {
	return slot[keyOffset:]
}

func (slot *KeySlot) GetKey() []byte {
	return slot[keyOffset:maxKeySize]
}

func (slot *KeySlot) SetKey(key []byte) (err error) {
	if len(key) > maxKeySize {
		err = gnet.ErrorInvalidContentLength
	} else {
		copy(slot[keyOffset:maxKeySize], key)
	}
	return
}
