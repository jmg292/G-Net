package certificate

import (
	"crypto/x509"
	"crypto/x509/pkix"
)

type HardwareBackedCertificate interface {
	Extensible
	ProvableOrigin
}

type hardwareBackedCertificate struct {
	certificate   *x509.Certificate
	proofOfOrigin *x509.Certificate
	extensions    *map[string]pkix.Extension
}

// Implement certificate.Extensible and certificate.ProvableOrigin for hardwareBackedCertificate
func (h hardwareBackedCertificate) Certificate() *x509.Certificate {
	return h.certificate
}

// Implement certificate.Extensible for hardwareBackedCertificate
func (h hardwareBackedCertificate) Extensions() (extensions *map[string]pkix.Extension) {
	if h.extensions != nil {
		extensions = h.extensions
	}
	return
}

// Implement certificate.ProvableOrigin for hardwareBackedCertificate
func (h hardwareBackedCertificate) ProofOfOrigin() (cert *x509.Certificate, err error) {
	if h.proofOfOrigin != nil {
		cert = h.proofOfOrigin
	} else {
		cert, err = GetProofOfOrigin(h)
	}
	return
}

func parseHardwareBackedCertificate(der []byte) (hwCert *hardwareBackedCertificate, err error) {
	if cert, e := x509.ParseCertificate(der); e != nil {
		err = e
	} else {
		hwCert = &hardwareBackedCertificate{certificate: cert}
		hwCert.extensions, err = parseExtensions(hwCert)
	}
	return
}
