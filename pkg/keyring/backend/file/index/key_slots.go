package index

import (
	"github.com/jmg292/G-Net/pkg/keyring"
	"github.com/jmg292/G-Net/pkg/keyring/key/slot"
)

func (*Index) GetKeySlotOffset(keySlot keyring.KeySlot) int {
	relativeOffset := int(keySlot) * slot.Size
	return int(keySlotBase) + relativeOffset
}
