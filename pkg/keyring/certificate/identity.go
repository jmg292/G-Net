package certificate

import (
	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/keyring"
	"github.com/jmg292/G-Net/pkg/keyring/certificate/extensions"
)

type Identity struct {
	hardwareBackedCertificate
}

func (identity *Identity) GetCertificateBySlot(keyslot keyring.KeySlot) (cert *hardwareBackedCertificate, err error) {
	if extensionOid, e := extensions.GetOIDByKeyslot(keyslot); e != nil {
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
