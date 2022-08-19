package yubikey

import (
	"crypto"
	"crypto/x509"
	"errors"

	"github.com/awnumar/memguard"
	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/keyring"
)

func (y *Backend) Name() (name string, err error) {
	name, err = getYubikeyName()
	return
}

func (y *Backend) Unlock(pin []byte) (err error) {
	y.pinMutex.Lock()
	if y.pin == nil {
		y.pin = memguard.NewEnclave(pin)
	}
	y.pinMutex.Unlock()
	// Validate PIN by using it
	_, err = y.getManagementKey()
	return
}

func (y *Backend) Lock() (err error) {
	y.pinMutex.Lock()
	defer y.pinMutex.Unlock()
	if y.pin != nil {
		y.pin = nil
	} else {
		err = gnet.ErrorKeystoreLocked
	}
	return nil
}

func (y *Backend) Close() error {
	instanceMutex.Lock()
	defer instanceMutex.Unlock()
	y.Lock()
	// Intentionally left locked
	y.handleMutex.Lock()
	if y.handle != nil {
		y.handle.Close()
		y.handle = nil
	}
	instance = nil
	return nil
}

func (y *Backend) CreateKey(keyslot keyring.KeySlot, keytype keyring.SupportedKeyType) (err error) {
	if keyslot == keyring.ManagementKeySlot && keytype == keyring.ManagementKey {
		err = y.createManagementKey()
	} else if keyslot == keyring.ManagementKeySlot || keytype == keyring.ManagementKey {
		err = gnet.ErrorUnsupportedAlgorithmForKeySlot
	} else if slot, e := convertKeyslotToPivSlot(keyslot); e != nil {
		err = e
	} else if alg, e := convertKeytypeToPivAlg(keytype); e != nil {
		err = e
	} else {
		err = y.createPivKey(slot, alg)
	}
	return
}

func (y *Backend) GetPrivateKey(keyslot keyring.KeySlot) (key crypto.PrivateKey, err error) {
	if keyslot == keyring.ManagementKeySlot {
		key, err = y.getManagementKey()
	} else if slot, e := convertKeyslotToPivSlot(keyslot); e == nil {
		key, err = y.getPrivateKey(slot)
	} else {
		err = e
	}
	return
}

func (y *Backend) GetPublicKey(keyslot keyring.KeySlot) (key crypto.PublicKey, err error) {
	if keyslot == keyring.ManagementKeySlot {
		err = gnet.ErrorExportNotAllowed
	} else if slot, e := convertKeyslotToPivSlot(keyslot); e == nil {
		key, err = y.getPublicKey(slot)
	} else {
		err = e
	}
	return
}

func (y *Backend) GetCertificate(keyslot keyring.KeySlot) (cert *x509.Certificate, err error) {
	if slot, e := convertKeyslotToPivSlot(keyslot); e != nil {
		err = e
	} else if handle, e := y.getYubikeyHandle(); e != nil {
		err = e
	} else {
		defer y.releaseYubikeyHandle()
		if cert, err = handle.Certificate(slot); errors.Is(err, piv.ErrNotFound) {
			err = gnet.ErrorCertificateNotFound
		}
	}
	return
}

func (y *Backend) SetCertificate(keyslot keyring.KeySlot, cert *x509.Certificate) (err error) {
	if currentCert, e := y.GetCertificate(keyslot); e != nil && e != gnet.ErrorCertificateNotFound {
		err = e
	} else if e == nil || currentCert != nil {
		err = gnet.ErrorCertAlreadyExists
	} else if slot, e := convertKeyslotToPivSlot(keyslot); e != nil {
		err = e
	} else if key, e := y.GetPrivateKey(keyring.ManagementKeySlot); e != nil {
		err = e
	} else if managementKey, ok := key.([keyring.ManagementKeySize]byte); !ok {
		err = gnet.ErrorInvalidManagementKey
	} else if handle, e := y.getYubikeyHandle(); e != nil {
		err = e
	} else {
		defer y.releaseYubikeyHandle()
		err = handle.SetCertificate(managementKey, slot, cert)
	}
	return
}

func (y *Backend) AttestationCertificate() (cert *x509.Certificate, err error) {
	if handle, e := y.getYubikeyHandle(); e != nil {
		err = e
	} else {
		defer y.releaseYubikeyHandle()
		cert, err = handle.AttestationCertificate()
	}
	return
}

func (y *Backend) Attest(keyslot keyring.KeySlot) (cert *x509.Certificate, err error) {
	if slot, e := convertKeyslotToPivSlot(keyslot); e != nil {
		err = e
	} else if handle, e := y.getYubikeyHandle(); e != nil {
		err = e
	} else {
		defer y.releaseYubikeyHandle()
		cert, err = handle.Attest(slot)
	}
	return
}
