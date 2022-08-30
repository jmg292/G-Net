package certificate

import (
	"crypto/x509"

	"github.com/jmg292/G-Net/pkg/certificate/extensions"
	gnet "github.com/jmg292/G-Net/pkg/gneterrs"
	"github.com/jmg292/G-Net/pkg/keyring"
)

type Identity struct {
	hardwareBackedCertificate
}

func (identity *Identity) GetCertificateBySlot(keyslot keyring.KeySlot) (cert *hardwareBackedCertificate, err error) {
	if extensionOid, e := extensions.GetSlotExtensionOID(keyslot); e != nil {
		err = e
	} else if certExtension, e := findExtensionByOID(identity, extensionOid); e != nil {
		err = e
	} else if certExtension.Value == nil {
		err = gnet.ErrorCertificateNotFound
	} else {
		cert, err = parseHardwareBackedCertificate(certExtension.Value)
	}
	return
}

func (identity *Identity) AttestationCertificate() (cert *x509.Certificate, err error) {
	if extAttestationCert, e := findExtensionByOID(identity, extensions.OIDVerifyProofOfOrigin); e != nil {
		err = e
	} else if extAttestationCert.Value == nil {
		err = gnet.ErrorInvalidAttestationCert
	} else {
		cert, err = x509.ParseCertificate(extAttestationCert.Value)
	}
	return
}

func (identity *Identity) SigningCertificate() (signingCert *Signing, err error) {
	if hwCert, e := identity.GetCertificateBySlot(keyring.SigningKeySlot); e != nil {
		err = e
	} else {
		signingCert = &Signing{*hwCert}
	}
	return
}

func (identity *Identity) AuthenticationCertificate() (authCert *Authentication, err error) {
	if hwCert, e := identity.GetCertificateBySlot(keyring.AuthenticationKeySlot); e != nil {
		err = e
	} else {
		authCert = &Authentication{*hwCert}
	}
	return
}

func (identity *Identity) EncryptionCertificate() (encryptionCert *Encryption, err error) {
	if hwCert, e := identity.GetCertificateBySlot(keyring.EncryptionKeySlot); e != nil {
		err = e
	} else {
		encryptionCert = &Encryption{*hwCert}
	}
	return
}
