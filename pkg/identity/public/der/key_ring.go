package der

import "crypto/x509"

type keyRing struct {
	name               string
	signingCert        *x509.Certificate
	authenticationCert *x509.Certificate
}
