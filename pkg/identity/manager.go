package identity

// CertificateManager is responsible for managing a keyring's identity certificate.
// Identity certificates are used to uniquely identify and interact with hardware keyrings.
// They are comprised of multiple component certificates, derived from keys stored within the keyring
// Each certificate is each responsible for a specific security service, including:
//   • Identity Verification
//   • User Authentication
//   • Device Authentication
//   • Encryption
//   • Non-Repudiation
// In addition, these certificates use special PKIX extensions to guarantee:
//   • The specific hardware used to generate PIV keys,
//   • The security controls protecting each PIV key,
//   • The manufacturer of the storage backend,
//   • And the privileges granted to the certificate holder
// Finally, the certificate manager is capable of generating, serializing, and parsing
// signing requests, which are used by a certificate authority to tie a keyring's identity
// into broader public key infrastructure.
type CertificateManager struct {
	keyring keyring.Storage
}

func New(keyring backend.Storage) (certman *CertificateManager, err error) {
	if _, err = keyring.Name(); err == nil {
		certman = &CertificateManager{keyring: keyring}
	}
	return
}
