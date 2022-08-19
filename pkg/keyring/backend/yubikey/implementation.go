package yubikey

import (
	"crypto"
	"crypto/x509"

	"github.com/awnumar/memguard"
	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/keyring"
)

func (y *Yubikey) Name() (name string, err error) {
	name, err = getYubikeyName()
	return
}

func (y *Yubikey) Unlock(pin []byte) (err error) {
	y.pinMutex.Lock()
	if y.pin == nil {
		y.pin = memguard.NewEnclave(pin)
	}
	y.pinMutex.Unlock()
	// Validate PIN by using it
	_, err = y.getManagementKey()
	return
}

func (y *Yubikey) Lock() (err error) {
	y.pinMutex.Lock()
	defer y.pinMutex.Unlock()
	if y.pin != nil {
		y.pin = nil
	} else {
		err = gnet.ErrorKeystoreLocked
	}
	return nil
}

func (y *Yubikey) Close() error {
	return gnet.ErrorNotYetImplemented
}

func (y *Yubikey) CreateKey(keyslot keyring.KeySlot, keytype keyring.SupportedKeyType) (err error) {
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

func (y *Yubikey) GetPrivateKey(keyslot keyring.KeySlot) (key crypto.PrivateKey, err error) {
	if keyslot == keyring.ManagementKeySlot {
		key, err = y.getManagementKey()
	} else if slot, e := convertKeyslotToPivSlot(keyslot); e == nil {
		key, err = y.getPrivateKey(slot)
	} else {
		err = e
	}
	return
}

func (y *Yubikey) GetPublicKey(keyslot keyring.KeySlot) (key crypto.PublicKey, err error) {
	if keyslot == keyring.ManagementKeySlot {
		err = gnet.ErrorExportNotAllowed
	} else if slot, e := convertKeyslotToPivSlot(keyslot); e == nil {
		key, err = y.getPublicKey(slot)
	} else {
		err = e
	}
	return
}

func (y *Yubikey) GetCertificate(keyslot keyring.KeySlot) (cert *x509.Certificate, err error) {
	if slot, e := convertKeyslotToPivSlot(keyslot); e != nil {
		err = e
	} else if handle, e := y.getYubikeyHandle(); e != nil {
		err = e
	} else {
		defer y.releaseYubikeyHandle()
		if cert, err = handle.Certificate(slot); err != nil && err == piv.ErrNotFound {
			err = gnet.ErrorCertificateNotFound
		}
	}
	return
}

func (y *Yubikey) SetCertificate(keyslot keyring.KeySlot, cert *x509.Certificate) (err error) {
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

func (y *Yubikey) AttestationCertificate() (cert *x509.Certificate, err error) {
	if handle, e := y.getYubikeyHandle(); e != nil {
		err = e
	} else {
		defer y.releaseYubikeyHandle()
		cert, err = handle.AttestationCertificate()
	}
	return
}

func (y *Yubikey) Attest(keyslot keyring.KeySlot) (cert *x509.Certificate, err error) {
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
