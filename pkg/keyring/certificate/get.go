package certificate

import (
	"crypto/x509"

	"github.com/jmg292/G-Net/pkg/gnet"
)

func (certstore *CertificateStore) GetCertificate(idx certificateIndex) (cert *x509.Certificate, err error) {
	if idx < 0 || int(idx) > len(certstore) {
		err = gnet.ErrorInvalidKeySlot
	} else {
		cert = certstore[idx]
	}
	if cert == nil {
		err = gnet.ErrorCertificateNotFound
	}
	return
}
