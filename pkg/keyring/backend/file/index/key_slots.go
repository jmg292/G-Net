package index

import (
	"github.com/jmg292/G-Net/pkg/keyring"
	"github.com/jmg292/G-Net/pkg/keyring/backend/file/keyslot"
)

func (*Index) GetKeySlotOffset(keySlot keyring.KeySlot) int {
	relativeOffset := int(keySlot) * keyslot.Size
	return int(keySlotBase) + relativeOffset
}
