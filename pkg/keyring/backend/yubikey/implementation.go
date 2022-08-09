package yubikey

import (
	"crypto"
	"crypto/x509"

	"github.com/awnumar/memguard"
	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/keyring"
)

func (y *YubikeyStorageBackend) Name() string {
	return y.name
}

func (y *YubikeyStorageBackend) Open() (err error) {
	if y.name, err = y.getCardName(); err == nil {
		if y.handle, err = piv.Open(y.name); err == nil {
			if y.handle == nil {
				err = gnet.ErrorUnableToOpenKeystore
			}
		}
	}
	return
}

func (y *YubikeyStorageBackend) Unlock(pin *memguard.Enclave) (err error) {
	y.pin = pin
	if handle, e := y.getHandle(); e != nil {
		err = e
	} else {
		defer y.releaseHandle()
		if pin, e := y.getPin(); e != nil {
			err = e
		} else {
			y.metadata, err = handle.Metadata(pin.String())
			if err != nil {
				y.Lock()
			}
		}
	}
	return
}

func (y *YubikeyStorageBackend) Lock() (err error) {
	err = y.clearPin()
	if y.metadata != nil {
		y.metadata = nil
	}
	return
}

func (y *YubikeyStorageBackend) Close() (err error) {
	y.Lock()
	if _, e := y.getHandle(); e != nil {
		err = e
	} else {
		defer y.releaseHandle()
		err = y.handle.Close()
		y.handle = nil
	}
	return
}

func (y *YubikeyStorageBackend) CreateKey(keytype keyring.SupportedKeyType, keyslot keyring.KeySlot) (err error) {
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
	} else if keyslot == keyring.ManagementKeySlot && keytype != keyring.ManagementKey {
		err = gnet.ErrorUnsupportedAlgorithmForKeySlot
	} else if keyslot != keyring.ManagementKeySlot && keytype == keyring.ManagementKey {
		err = gnet.ErrorUnsupportedAlgorithmForKeySlot
	} else if keyslot == keyring.ManagementKeySlot && keytype == keyring.ManagementKey {
		err = y.generateManagementKey()
	} else if slot, e := convertToPivSlot(keyslot); e != nil {
		err = e
	} else if alg, e := convertToPivAlg(keytype); e != nil {
		err = e
	} else {
		err = y.generateKey(slot, alg)
	}
	return
}

func (y *YubikeyStorageBackend) GetPrivateKey(keyslot keyring.KeySlot) (key crypto.PrivateKey, err error) {
	if slot, e := convertToPivSlot(keyslot); e == nil {
		if publickey, e := y.GetPublicKey(keyslot); e != nil {
			err = e
		} else {
			key, err = y.getPivPrivateKey(slot, publickey)
		}
	} else if e == gnet.ErrorInvalidKeySlot && keyslot == keyring.EncryptionKeySlot {

	} else {
		err = e
	}
	return
}

func (*YubikeyStorageBackend) GetPrivateBytes(_ keyring.KeySlot) ([]byte, error) {
	return nil, gnet.ErrorExportNotAllowed
}

func (y *YubikeyStorageBackend) GetPublicKey(keyslot keyring.KeySlot) (key crypto.PublicKey, err error) {
	if keyslot == keyring.EncryptionKeySlot {
		err = gnet.ErrorNotYetImplemented
	} else if cert, e := y.Attest(keyslot); e != nil && e != piv.ErrNotFound {
		err = e
	} else if e != nil && e == piv.ErrNotFound {
		err = gnet.ErrorKeyNotFound
	} else {
		key = cert.PublicKey
	}
	return
}

func (y *YubikeyStorageBackend) GetPublicBytes(keyslot keyring.KeySlot) (keyBytes []byte, err error) {
	if _, e := y.GetPublicKey(keyslot); e != nil {
		err = e
	} else {
		err = gnet.ErrorNotYetImplemented
	}
	return
}

func (*YubikeyStorageBackend) PutPrivateKey(_ crypto.PrivateKey, _ keyring.KeySlot) error {
	return gnet.ErrorImportNotAllowed
}

func (*YubikeyStorageBackend) PutPrivateBytes(_ []byte, _ keyring.KeySlot, _ bool) error {
	return gnet.ErrorImportNotAllowed
}

func (*YubikeyStorageBackend) PutPublicKey(_ crypto.PublicKey, _ keyring.KeySlot, _ bool) error {
	return gnet.ErrorImportNotAllowed
}

func (*YubikeyStorageBackend) PutPublicBytes(_ []byte, _ keyring.KeySlot, _ bool) error {
	return gnet.ErrorImportNotAllowed
}

func (y *YubikeyStorageBackend) Attest(keyslot keyring.KeySlot) (cert *x509.Certificate, err error) {
	if handle, e := y.getHandle(); e != nil {
		err = e
	} else {
		defer y.releaseHandle()
		if slot, e := convertToPivSlot(keyslot); e != nil {
			err = e
		} else {
			cert, err = handle.Attest(slot)
		}
	}
	return
}

func (y *YubikeyStorageBackend) AttestationCertificate() (cert *x509.Certificate, err error) {
	if handle, e := y.getHandle(); e != nil {
		err = e
	} else {
		defer y.releaseHandle()
		cert, err = handle.AttestationCertificate()
	}
	return
}
