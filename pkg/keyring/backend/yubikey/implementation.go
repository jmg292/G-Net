package yubikey

import (
	"crypto"
	"crypto/x509"

	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/keyring"
)

func (y *Yubikey) Name() (name string, err error) {
	name, err = getYubikeyName()
	return
}

func (y *Yubikey) Unlock(pin []byte) error {
	return gnet.ErrorNotYetImplemented
}

func (y *Yubikey) Lock() error {
	return gnet.ErrorNotYetImplemented
}

func (y *Yubikey) Close() error {
	return gnet.ErrorNotYetImplemented
}

func (y *Yubikey) CreateKey(keyslot keyring.KeySlot, keytype keyring.SupportedKeyType) (err error) {
	if slot, e := convertKeyslotToPivSlot(keyslot); e != nil {
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
	if slot, e := convertKeyslotToPivSlot(keyslot); e == nil {
		key, err = y.getPublicKey(slot)
	} else {
		err = e
	}
	return
}

func (y *Yubikey) GetCertificate(keyslot keyring.KeySlot) (*x509.Certificate, error) {
	return nil, gnet.ErrorNotYetImplemented
}

func (y *Yubikey) SetCertificate(keyslot keyring.KeySlot, cert *x509.CertPool) error {
	return gnet.ErrorNotYetImplemented
}

func (y *Yubikey) CreateCSR(keyslot keyring.KeySlot) (*x509.Certificate, error) {
	return nil, gnet.ErrorNotYetImplemented
}

func (y *Yubikey) SignCSR(keyslot keyring.KeySlot, csr *x509.CertificateRequest) error {
	return gnet.ErrorNotYetImplemented
}

func (y *Yubikey) AttestationCertificate() (*x509.Certificate, error) {
	return nil, gnet.ErrorNotYetImplemented
}

func (y *Yubikey) Attest(keyslot keyring.KeySlot) (*x509.Certificate, error) {
	return nil, gnet.ErrorNotYetImplemented
}
