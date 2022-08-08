package yubikey

import (
	"crypto/rand"

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

func (y *yubikeyStorageBackend) deriveX25519PrivateKey(salt []byte) (private key.X25519PrivateKey, err error) {
	if err = y.assertOpenAndUnlocked(); err == nil {
		private = *key.NewX25519PrivateKey(kdf.DeriveKey(y.metadata.ManagementKey[:], salt))
	}
	return
}
