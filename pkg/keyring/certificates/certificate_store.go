package certificates

import (
	"crypto/x509"
)

type certificateIndex int

const (
	Signing certificateIndex = iota
	SigningKeyAttestation
	Authentication
	AuthenticationKeyAttestation
	EncryptionKeyAttestation
	Device
	DeviceKeyAttestation
)
const CertificateCount = 7

type CertificateStore [CertificateCount]*x509.Certificate

func Parse(derBytes []byte) (cs *CertificateStore, err error) {
	var certstore CertificateStore
	if err = certstore.ParseDER(derBytes); err == nil {
		cs = &certstore
	}
	return
}
