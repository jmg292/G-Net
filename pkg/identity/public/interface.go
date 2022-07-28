package public

import "crypto/x509"

type KeyRing interface {
	Name() string
	Fingerprint() []byte
	SigningCertificate() *x509.Certificate
	AuthenticationCertificate() *x509.Certificate
	VerifySignature([]byte, []byte) error
	VerifyAuthentication([]byte, []byte) error
}
