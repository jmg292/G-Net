package certificate

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"fmt"

	"github.com/jmg292/G-Net/pkg/gnet"
)

type Extensible interface {
	Certificate() *x509.Certificate
	Extensions() map[string]pkix.Extension
}

func parseExtensions(extended Extensible) (extensions map[string]pkix.Extension, err error) {
	var unhandledExtensions []pkix.Extension
	if extended.Certificate() == nil {
		err = gnet.ErrorInvalidCertificate
	} else {
		if extended.Certificate().Extensions != nil {
			unhandledExtensions = append(unhandledExtensions, extended.Certificate().Extensions...)
		}
		if extended.Certificate().ExtraExtensions != nil {
			unhandledExtensions = append(unhandledExtensions, extended.Certificate().ExtraExtensions...)
		}
	}
	for _, extension := range unhandledExtensions {
		oid := extension.Id.String()
		if _, keyExists := extensions[oid]; keyExists {
			err = fmt.Errorf(gnet.ErrorDuplicateExtension.Error(), oid)
			break
		} else {
			extensions[oid] = extension
		}
	}
	return
}

func findExtensionByOID(cert Extensible, oid asn1.ObjectIdentifier) (extension *pkix.Extension, err error) {
	if cert.Extensions() == nil {
		err = gnet.ErrorInvalidCertificate
	} else if value, exists := cert.Extensions()[oid.String()]; !exists {
		err = fmt.Errorf(gnet.ErrorNoSuchExtension.Error(), oid.String())
	} else {
		extension = &value
	}
	return
}
