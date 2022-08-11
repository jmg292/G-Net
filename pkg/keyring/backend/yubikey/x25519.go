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

func (p *x25519PublicBytes) AsManagementKeys() (key1 [24]byte, key2 [24]byte) {
	copy(key1[:], p[:24])
	copy(key2[:], p[24:])
	return
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
	var managementkey *[24]byte
	if err = y.assertOpenAndUnlocked(); err == nil {
		if y.metadata == nil || y.metadata.ManagementKey == nil {
			managementkey = &piv.DefaultManagementKey
		} else {
			managementkey = y.metadata.ManagementKey
		}
		private = key.NewX25519PrivateKey(kdf.DeriveKey(managementkey[:], salt))
	}
	return
}

func (y *YubikeyStorageBackend) getX25519KeySlots() (slot1 piv.Slot, slot2 piv.Slot, err error) {
	if slot0x82, ok := piv.RetiredKeyManagementSlot(0x82); !ok {
		err = gnet.ErrorInvalidKeySlot
	} else if slot0x83, ok := piv.RetiredKeyManagementSlot(0x83); !ok {
		err = gnet.ErrorInvalidKeySlot
	} else {
		slot1, slot2 = slot0x82, slot0x83
	}
	return
}
