package keyring

import (
	"crypto"
	"io"
)

type PrivateKeyRing interface {
	DeviceInfo
	crypto.PrivateKey
	crypto.Signer
	crypto.Decrypter
}

// Private is a crypto.PrivateKey implementation for keyring.Backend.
// It implements crypto.PrivateKey, crypto.Signer and crypto.Decrypter
type Private struct {
	Keyring Keystore
	public  *Public
}

// NewPrivate creates and returns a new instance of Private using the provided Backend
// The Backend must be open, unlocked, and available for use
func NewPrivate(backend Keystore) (keyring *Private, err error) {
	if err = VerifyKeyAvailability(backend, true); err == nil {
		keyring = &Private{Keyring: backend}
		keyring.public, err = NewPublic(backend)
	}
	return
}

func (private Private) signWithKey(key crypto.PrivateKey, rand io.Reader, digest []byte, opts crypto.SignerOpts) (signature []byte, err error) {
	if signer, ok := key.(crypto.Signer); !ok {
		err = ErrInvalidSigningKey
	} else {
		signature, err = signer.Sign(rand, digest, opts)
	}
	return
}

// Public implements crypto.PrivateKey for Private.
// It returns the public key associated with this private key.
func (private Private) Public() crypto.PublicKey {
	return &Public{Keyring: private.Keyring}
}

// Sign implements crypto.Signer for Private
// It uses the Backend's signing key to sign the provided digest
func (private Private) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) (signature []byte, err error) {
	if signingKey, e := private.Keyring.GetPrivateKey(SigningKeySlot); e != nil {
		err = e
	} else {
		signature, err = private.signWithKey(signingKey, rand, digest, opts)
	}
	return
}

// Decrypt implements crypto.Decrypter for Private
// It uses the Backend's encryption key to decrypt the provided msg
func (private Private) Decrypt(rand io.Reader, msg []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error) {
	if encryptionKey, e := private.Keyring.GetPrivateKey(EncryptionKeySlot); e != nil {
		err = e
	} else if decrypter, ok := encryptionKey.(crypto.Decrypter); !ok {
		err = ErrInvalidDecryptionKey
	} else {
		plaintext, err = decrypter.Decrypt(rand, msg, opts)
	}
	return
}

// Authenticate provides a convenience wrapper around retrieving
// and signing a digest with a stored Authentication key
func (private Private) Authenticate(rand io.Reader, digest []byte, opts crypto.SignerOpts) (signature []byte, err error) {
	if authKey, e := private.Keyring.GetPrivateKey(AuthenticationKeySlot); e != nil {
		err = e
	} else {
		signature, err = private.signWithKey(authKey, rand, digest, opts)
	}
	return
}
