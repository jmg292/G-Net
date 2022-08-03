package der

import (
	"crypto/x509"
	"fmt"

	"github.com/jmg292/G-Net/pkg/gnet"
)

func rsaValidateSignature(cert *x509.Certificate, data []byte, signature []byte) error {
	return fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func ecdsaValidateSignature(cert *x509.Certificate, data []byte, signature []byte) error {
	return fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func ed25519ValidateSignature(cert *x509.Certificate, data []byte, signature []byte) error {
	return fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func x509ValidateSignature(cert *x509.Certificate, data []byte, signature []byte) error {
	switch cert.PublicKeyAlgorithm {
	case x509.RSA:
		return rsaValidateSignature(cert, data, signature)
	case x509.ECDSA:
		return ecdsaValidateSignature(cert, data, signature)
	case x509.Ed25519:
		return ed25519ValidateSignature(cert, data, signature)
	}
	return fmt.Errorf(string(gnet.ErrorUnsupportedAlgorithm))
}
