package yubikey

import (
	"crypto"
	"crypto/x509"

	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/keyring"
)

func (y *yubikeyStorageBackend) Name() string {
	return y.name
}

func (y *yubikeyStorageBackend) Open() (err error) {
	if y.name, err = y.getCardName(); err == nil {
		if y.handle, err = piv.Open(y.name); err == nil {
			if y.handle == nil {
				err = gnet.ErrorUnableToOpenKeystore
			}
		}
	}
	return
}

func (y *yubikeyStorageBackend) Unlock(pin []byte) (err error) {
	if err = y.assertOpen(); err == nil {
		y.metadata, err = y.handle.Metadata(string(pin))
	}
	return
}

func (y *yubikeyStorageBackend) Lock() (err error) {
	if y.metadata != nil {
		y.metadata = nil
	}
	return
}

func (y *yubikeyStorageBackend) Close() (err error) {
	y.Lock()
	if err = y.assertOpen(); err == nil {
		err = y.handle.Close()
		y.handle = nil
	}
	return
}

func (y *yubikeyStorageBackend) CreateKey(keytype keyring.SupportedKeyType, keyslot keyring.KeySlot) (err error) {
	if keyslot == keyring.EncryptionKeySlot && keytype != keyring.X25519Key {
		err = gnet.ErrorUnsupportedAlgorithmForKeySlot
	} else if keyslot != keyring.EncryptionKeySlot && keytype == keyring.X25519Key {
		err = gnet.ErrorUnsupportedAlgorithmForKeySlot
	} else if keyslot == keyring.EncryptionKeySlot && keytype == keyring.X25519Key {
		if pubkey, e := y.generateEncryptionKey(); e != nil {
			err = e
		} else {
			err = y.storeEncryptionKey(pubkey)
		}
	} else if slot, e := convertToPivSlot(keyslot); e != nil {
		err = e
	} else if alg, e := convertToPivAlg(keytype); e != nil {
		err = e
	} else {
		err = y.generateKey(slot, alg)
	}
	return
}

func (y *yubikeyStorageBackend) GetPrivateKey(keyslot keyring.KeySlot) (key crypto.PrivateKey, err error) {
	if err = y.assertOpenAndUnlocked(); err == nil {
		if slot, e := convertToPivSlot(keyslot); e != nil {
			err = e
		} else if public, e := y.GetPublicKey(keyslot); e != nil {
			err = e
		} else {
			key, err = y.handle.PrivateKey(slot, public, piv.KeyAuth{PIN: string(y.metadata.ManagementKey[:])})
		}
	}
	return
}

func (*yubikeyStorageBackend) GetPrivateBytes(_ keyring.KeySlot) ([]byte, error) {
	return nil, gnet.ErrorExportNotAllowed
}

func (y *yubikeyStorageBackend) GetPublicKey(keyslot keyring.KeySlot) (key crypto.PublicKey, err error) {
	if keyslot == keyring.EncryptionKeySlot {
		err = gnet.ErrorNotYetImplemented
	} else if cert, e := y.Attest(keyslot); e != nil {
		err = e
	} else {
		key = cert.PublicKey
	}
	return
}

func (y *yubikeyStorageBackend) GetPublicBytes(keyslot keyring.KeySlot) (keyBytes []byte, err error) {
	if _, e := y.GetPublicKey(keyslot); e != nil {
		err = e
	} else {
		err = gnet.ErrorNotYetImplemented
	}
	return
}

func (*yubikeyStorageBackend) PutPrivateKey(_ crypto.PrivateKey, _ keyring.KeySlot) error {
	return gnet.ErrorImportNotAllowed
}

func (*yubikeyStorageBackend) PutPrivateBytes(_ []byte, _ keyring.KeySlot, _ bool) error {
	return gnet.ErrorImportNotAllowed
}

func (*yubikeyStorageBackend) PutPublicKey(_ crypto.PublicKey, _ keyring.KeySlot, _ bool) error {
	return gnet.ErrorImportNotAllowed
}

func (*yubikeyStorageBackend) PutPublicBytes(_ []byte, _ keyring.KeySlot, _ bool) error {
	return gnet.ErrorImportNotAllowed
}

func (y *yubikeyStorageBackend) Attest(keyslot keyring.KeySlot) (cert *x509.Certificate, err error) {
	if err = y.assertOpen(); err == nil {
		if slot, e := convertToPivSlot(keyslot); e != nil {
			err = e
		} else {
			cert, err = y.handle.Attest(slot)
		}
	}
	return
}

func (y *yubikeyStorageBackend) AttestationCertificate() (cert *x509.Certificate, err error) {
	if err = y.assertOpen(); err == nil {
		cert, err = y.handle.AttestationCertificate()
	}
	return
}
