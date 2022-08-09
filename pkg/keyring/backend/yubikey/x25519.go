package yubikey

import (
	"crypto/rand"
	"crypto/subtle"

	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/keyring/kdf"
	"github.com/jmg292/G-Net/pkg/keyring/key"
)

type x25519PublicBytes [48]byte

func (p *x25519PublicBytes) Salt() []byte {
	return p[:16]
}

func (p *x25519PublicBytes) Key() []byte {
	return p[16:]
}

func (p *x25519PublicBytes) GenerateSalt() {
	rand.Read(p.Salt())
}

func newX25519PublicBytes(b []byte) (publicbytes x25519PublicBytes, err error) {
	if len(b) < len(publicbytes) {
		err = gnet.ErrorInvalidContentLength
	} else {
		subtle.ConstantTimeCopy(1, b, publicbytes[:])
	}
	return
}

func (y *YubikeyStorageBackend) deriveX25519PrivateKey(salt []byte) (private *key.X25519PrivateKey, err error) {
	if err = y.assertOpenAndUnlocked(); err == nil {
		private = key.NewX25519PrivateKey(kdf.DeriveKey(y.metadata.ManagementKey[:], salt))
	}
	return
}

func (y *YubikeyStorageBackend) getX25519KeySlots() (slot1 piv.Slot, slot2 piv.Slot, err error) {
	if slot0x95, ok := piv.RetiredKeyManagementSlot(0x95); !ok {
		err = gnet.ErrorInvalidKeySlot
	} else if slot0x94, ok := piv.RetiredKeyManagementSlot(0x94); !ok {
		err = gnet.ErrorInvalidKeySlot
	} else {
		slot1, slot2 = slot0x94, slot0x95
	}
	return
}
