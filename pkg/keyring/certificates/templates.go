package certificates

import "crypto/x509"

var signingCertTemplate = &x509.Certificate{
	BasicConstraintsValid: true,
	IsCA:                  true,
	MaxPathLen:            1,
	KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign | x509.KeyUsageCRLSign | x509.KeyUsageContentCommitment,
}

var authenticationCertTemplate = &x509.Certificate{
	BasicConstraintsValid: true,
	MaxPathLen:            0,
	MaxPathLenZero:        true,
	KeyUsage:              x509.KeyUsageContentCommitment | x509.KeyUsageDigitalSignature,
}

var encryptionCertTemplate = &x509.Certificate{
	BasicConstraintsValid: true,
	MaxPathLen:            0,
	MaxPathLenZero:        true,
	KeyUsage:              x509.KeyUsageDataEncipherment | x509.KeyUsageKeyAgreement | x509.KeyUsageKeyEncipherment | x509.KeyUsageEncipherOnly,
}

var deviceCertTemplate = &x509.Certificate{
	BasicConstraintsValid: true,
	MaxPathLen:            0,
	MaxPathLenZero:        true,
	KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageContentCommitment,
}
