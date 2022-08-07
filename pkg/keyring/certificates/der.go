package certificates

import (
	"crypto/x509"
	"fmt"

	"github.com/jmg292/G-Net/pkg/gnet"
)

func (certstore *CertificateStore) DERSize() (size int, err error) {
	for i := 0; i < len(certstore); i++ {
		if cert, e := certstore.GetCertificate(certificateIndex(i)); e != nil {
			err = e
			break
		} else {
			size += len(cert.Raw)
		}
	}
	return
}

func (certstore *CertificateStore) DER() (derBytes []byte, err error) {
	if size, e := certstore.DERSize(); e != nil {
		err = e
	} else {
		derBytes = make([]byte, size)
		for i := 0; i < len(certstore); i++ {
			if cert, e := certstore.GetCertificate(certificateIndex(i)); e != nil {
				err = e
			} else {
				derBytes = append(derBytes, cert.Raw...)
			}
		}
	}
	return
}

func (certstore *CertificateStore) ParseDER(derBytes []byte) (err error) {
	if certs, e := x509.ParseCertificates(derBytes); e != nil {
		err = e
	} else if len(certs) < len(certstore) {
		err = fmt.Errorf(string(gnet.ErrorInvalidContentLength))
	} else {
		copy(certstore[:], certs)
	}
	return
}
