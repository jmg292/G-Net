package keyslot

import (
	"crypto"
	"fmt"

	gcrypt "github.com/jmg292/G-Net/pkg/crypto"
	"github.com/jmg292/G-Net/pkg/gnet"
)

const Size int = 100

type keySlot [Size]byte

func Empty() *keySlot {
	var slot keySlot
	return &slot
}

func New(keyType gcrypt.SupportedKeyType, key crypto.PrivateKey, managementKey []byte) (slot *keySlot, err error) {
	slot = Empty()
	err = fmt.Errorf(string(gnet.ErrorNotYetImplemented))
	return
}
