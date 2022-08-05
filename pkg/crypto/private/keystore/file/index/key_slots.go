package index

import (
	"github.com/jmg292/G-Net/pkg/crypto"
	"github.com/jmg292/G-Net/pkg/crypto/private/keystore/file/keyslot"
)

func (*Index) GetKeySlotOffset(keySlot crypto.KeySlot) int {
	relativeOffset := int(keySlot) * keyslot.Size
	return int(keySlotBase) + relativeOffset
}
