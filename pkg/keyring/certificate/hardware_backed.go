package certificate

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"fmt"

	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/keyring/certificate/extensions"
)

type hardwareBackedCertificate struct {
	Certificate   *x509.Certificate
	proofOfOrigin *x509.Certificate
	extensions    map[string]pkix.Extension
}

func (h *hardwareBackedCertificate) loadExtensions() (err error) {
	var unhandledExtensions []pkix.Extension
	if h.Certificate.Extensions != nil {
		unhandledExtensions = append(unhandledExtensions, h.Certificate.Extensions...)
	}
	if h.Certificate.ExtraExtensions != nil {
		unhandledExtensions = append(unhandledExtensions, h.Certificate.ExtraExtensions...)
	}
	for _, extension := range unhandledExtensions {
		oid := extension.Id.String()
		if _, keyExists := h.extensions[oid]; keyExists {
			err = fmt.Errorf(gnet.ErrorDuplicateExtension.Error(), oid)
			break
		} else {
			h.extensions[oid] = extension
		}
	}
	return
}

func (h *hardwareBackedCertificate) loadProofOfOrigin() (err error) {
	if h.proofOfOrigin != nil {
		err = gnet.ErrorCertAlreadyExists
	} else if attestationCertDer, e := h.GetExtensionByOID(extensions.OIDProofOfOrigin); e != nil {
		err = e
	} else if attestationCert, e := x509.ParseCertificate(attestationCertDer.Value); e != nil {
		err = e
	} else {
		h.proofOfOrigin = attestationCert
	}
	return
}

func (h *hardwareBackedCertificate) ProofOfOrigin() (attestationCert *x509.Certificate, err error) {
	if h.proofOfOrigin == nil {
		err = gnet.ErrorInvalidAttestationCert
	} else {
		attestationCert = h.proofOfOrigin
	}
	return
}

func (h *hardwareBackedCertificate) GetExtensionByOID(oid asn1.ObjectIdentifier) (extension *pkix.Extension, err error) {
	if value, exists := h.extensions[oid.String()]; !exists {
		err = fmt.Errorf(gnet.ErrorNoSuchExtension.Error(), oid.String())
	} else {
		extension = &value
	}
	return
}

func (h *hardwareBackedCertificate) LoadCertificate(cert *x509.Certificate) (err error) {
	if h.Certificate != nil {
		err = gnet.ErrorCertAlreadyExists
	} else if cert == nil {
		err = gnet.ErrorInvalidCertificate
	} else {
		h.Certificate = cert
		if err = h.loadExtensions(); err == nil {
			err = h.loadProofOfOrigin()
		}
	}
	return
}

func (h *hardwareBackedCertificate) LoadCertificateDER(der []byte) (err error) {
	if cert, e := x509.ParseCertificate(der); e != nil {
		err = e
	} else {
		h.LoadCertificate(cert)
	}
	return
}
