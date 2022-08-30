package identity

import (
	"crypto/x509"

	gnet "github.com/jmg292/G-Net/pkg/gneterrs"
)

func mapCertMemberstoCSR(cert *x509.Certificate) (csr *x509.CertificateRequest, err error) {
	if cert == nil {
		err = gnet.ErrorInvalidCertificateTemplate
	} else {
		// Only the fields below are parsed by x509.CreateCertificateRequest
		// See: https://pkg.go.dev/crypto/x509#CreateCertificateRequest
		csr = &x509.CertificateRequest{
			SignatureAlgorithm: cert.SignatureAlgorithm,
			Subject:            cert.Subject,
			DNSNames:           cert.DNSNames,
			EmailAddresses:     cert.EmailAddresses,
			URIs:               cert.URIs,
			ExtraExtensions:    cert.ExtraExtensions,
		}
	}
	return
}

func mapCSRMembersToCertTemplate(csr *x509.CertificateRequest, template *x509.Certificate) (cert *x509.Certificate, err error) {
	if csr == nil {
		err = gnet.ErrorInvalidCSR
	} else if template == nil {
		err = gnet.ErrorInvalidCertificateTemplate
	} else {
		// Only the fields below are parsed by x509.CreateCertificate
		// See: https://pkg.go.dev/crypto/x509#CreateCertificate
		cert = &x509.Certificate{
			AuthorityKeyId:              template.AuthorityKeyId, // Only used if cert is self-signed
			BasicConstraintsValid:       template.BasicConstraintsValid,
			CRLDistributionPoints:       template.CRLDistributionPoints,
			DNSNames:                    csr.DNSNames,
			EmailAddresses:              csr.EmailAddresses,
			ExcludedDNSDomains:          template.ExcludedDNSDomains,
			ExcludedEmailAddresses:      template.ExcludedEmailAddresses,
			ExcludedIPRanges:            template.ExcludedIPRanges,
			ExcludedURIDomains:          template.ExcludedURIDomains,
			ExtKeyUsage:                 template.ExtKeyUsage,
			ExtraExtensions:             csr.ExtraExtensions,
			IPAddresses:                 template.IPAddresses,
			IsCA:                        template.IsCA,
			IssuingCertificateURL:       template.IssuingCertificateURL,
			KeyUsage:                    template.KeyUsage,
			MaxPathLen:                  template.MaxPathLen,
			MaxPathLenZero:              template.MaxPathLenZero,
			NotAfter:                    template.NotAfter,
			NotBefore:                   template.NotBefore,
			OCSPServer:                  template.OCSPServer,
			PermittedDNSDomains:         template.PermittedDNSDomains,
			PermittedDNSDomainsCritical: template.PermittedDNSDomainsCritical,
			PermittedEmailAddresses:     template.PermittedEmailAddresses,
			PermittedIPRanges:           template.PermittedIPRanges,
			PermittedURIDomains:         template.PermittedURIDomains,
			PolicyIdentifiers:           template.PolicyIdentifiers,
			SerialNumber:                template.SerialNumber,
			SignatureAlgorithm:          csr.SignatureAlgorithm,
			Subject:                     csr.Subject,
			SubjectKeyId:                template.SubjectKeyId,
			URIs:                        csr.URIs,
			UnknownExtKeyUsage:          template.UnknownExtKeyUsage,
		}
	}
	return
}
