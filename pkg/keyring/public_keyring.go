package keyring

// Public represents the public components of a Backend implementation
type Public struct {
	Keyring Keystore
}

// NewPublic creates a new instance of Public from an instance of Backend
// The provided instance must be opened and available for use
func NewPublic(backend Keystore) (keyring *Public, err error) {
	if err = VerifyKeyAvailability(backend, false); err == nil {
		keyring = &Public{Keyring: backend}
	}
	return
}

type PublicKeyRing interface {
	DeviceInfo
	// VerifySignature uses a certificate's public
	// signing key to verify signed data
	VerifySignature(data, signature []byte) bool

	// VerifyAuthentication uses a certificate's public
	// authentication key to verify authenticated data
	VerifyAuthentication(data, authentication []byte) bool

	// VerifyAttestation uses a certificate's embedded
	// attestation certificate to verify key attesatation data
	VerifyKeySlotAttestation(slot KeySlot, attestation []byte) bool

	// Seal implements cipher.AEAD for keyring
	Seal(dst, nonce, plaintext, additionalData []byte) ([]byte, error)
}
