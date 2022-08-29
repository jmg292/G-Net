package keyring

import (
	"crypto"
	"crypto/x509"
)

type KeyInfo interface {
	KeySlot() KeySlot
	KeyType() KeyType
	KeyIdentifier() []byte
	KeySecurityPolicy() []byte
}

type Keystore interface {

	// KeyInfo returns an instance of KeyInfo
	// for the key stored in the specified slot.
	//
	// KeyInfo can only be called after establishing an
	// unprivileged session with the keystore.
	//
	// It returns an error in the following scenarios:
	//
	//    • If there are no open sessions, KeyInfo will
	//      return an error that wraps ErrClosed
	//    • If no key exists in the specified slot, KeyInfo
	//      will return an error that wraps ErrNotFound
	//    • If an unspecified error occurs, KeyInfo will
	//      return the error without modification
	KeyInfo(KeySlot) (KeyInfo, error)

	// Open attempts to establish an unprivileged
	// session with the keystore.
	//
	// It returns an error in the following scenarios:
	//
	//    • If a the keystore isn't available to establish
	//      a session, Open will return an error that wraps
	//      ErrNotFound
	//    • If an unspecified error occurred while attempting
	//      to create a key in the specified slot, Open
	//      will return the error without modification
	Open() error

	// Unlock uses the provided information to
	// establish a privileged session with the
	// keystore.
	//
	// Unlock can only be called after establishing an
	// unprivileged session with the keystore.
	//
	// It returns an error in the following scenarios:
	//
	//    • If there are no open unprivileged sessions,
	//      Unlock will return an error that wraps ErrClosed
	//    • If the provided information can't be used to
	//      establish an authenticated session, Unlock
	//      will return an error that wraps ErrInvalid
	//    • If an unspecified error occurs, Unlock will
	//      return the error without modification
	Unlock([]byte) error

	// Lock terminates all privileged keystore sessions.
	// It does not terminate unprivileged sessions.
	//
	// It returns an error in the following scenarios:
	//
	//    • If no privileged session exists when Lock is called,
	//      it will return terminating a privileged session, or if no
	//    • If an unspecified error occurs, Lock will
	//      return the error without modification
	Lock() error

	// Close terminates all open keystore connections,
	// including privileged sessions.
	//
	// It returns an error in the following scenarios:
	//
	//    • If there are no open sessions, Close will
	//      return an error that wraps ErrClosed
	Close() error

	// CreateKey attempts to create a key of the specified
	// KeyType within the specified KeySlot. If the NilKey
	// KeyType is passed to CreateKey, it will attempt to
	// destroy the key in the specified KeySlot.
	//
	// CreateKey can only be used after a privileged session
	// is established.
	//
	// It returns an error in the following scenarios:
	//
	//    • If there are no open sessions available for use,
	//      CreateKey will return an error that wraps ErrClosed
	//    • If there are no privileged sessions available
	//      for use, CreateKey will return an error that wraps
	//      ErrPermission
	//    • If a KeyType is not NilKey and a key already exists
	//      in the specified slot, CreateKey will return an
	//      error that wraps ErrExist
	//    • If the specified KeySlot does not support the specified
	//      KeyType, CreateKey will return an error that wraps
	//      ErrInvalid
	//    • If an unspecified error occurs, CreateKey will
	//      return the error without modification
	CreateKey(KeySlot, KeyType) error

	// GetPrivateKey attempts to retrieve a handle to the
	// private key stored within in the specified KeySlot.  The handle
	// is returned as an instance of crypto.PrivateKey
	//
	// Private keys cannot be exported from the keystore in which
	// they were generated.  Therefore, the handle returned by
	// GetPrivateKey does not provide a method for accessing a
	// private key directly.  Instead, the key handle represents
	// a channel that is used to:
	//
	//    • Pass data to and from the keystore.
	//    • Direct the keystore to perform cryptographic
	//      operations on that data using the key stored
	//      within KeySlot.
	//
	// GetPrivateKey can only be called after a privileged session
	// is established.
	//
	// It returns an error in the following scenarios:
	//
	//    • If there are no open sessions available for use,
	//      GetPrivateKey will return an error that wraps
	//      ErrClosed
	//    • If there are no privileged sessions available for use,
	//      GetPrivateKey will return an error that wraps
	//      ErrPermission
	//    • If no key is stored within the specified slot,
	//      GetPrivateKey will return an error that wraps
	//      ErrNotExist
	//    • If an unspecified error occurs, CreateKey will
	//      return the error without modification
	GetPrivateKey(KeySlot) (crypto.PrivateKey, error)

	// GetPublicKey attempts to retrieve a handle to the public key
	// stored within in the specified KeySlot. The handle is returned
	// as an instance of crypto.PublicKey.
	//
	// GetPublicKey can only be called after establishing an
	// an unprivileged session with the keystore.
	//
	// It returns an error in the following scenarios:
	//
	//    • If there are no open sessions available for use,
	//      GetPublicKey will return an error that wraps
	//      ErrClosed
	//    • If no key is stored within the specified slot,
	//      GetPublicKey will return an error that wraps
	//      ErrNotExist
	//    • If an unspecified error occurs, GetPublicKey
	//      will return the error without modification
	GetPublicKey(KeySlot) (crypto.PublicKey, error)

	// GetCertificate attempts to retrieve a handle to the
	// *x509.Certificate certificate stored within the specified
	// KeySlot.
	//
	// GetCertificate can only be called after establishing an
	// unprivileged session with the keystore.
	//
	// It returns an error in the following scenarios:
	//
	//    • If there are no open sessions available for use,
	//      GetCertificate will return an error that wraps
	//      ErrClosed
	//    • If no certificate is stored within the specified
	//      KeySlot, GetCertificate will return an error that
	//      wraps ErrNotExist
	//    • If an unspecified error occurs, GetCertificate
	//      will return the error without modification
	GetCertificate(KeySlot) (*x509.Certificate, error)

	// SetCertificate attempts to store an *x509.Certificate within
	// the specified KeySlot. If a nil *x509.Certificate is passed
	// to SetCertificate, it will attempt to destroy the certificate
	// in the specified KeySlot.
	//
	// SetCertificate can only be called after establishing an
	// unprivileged session with the keystore.
	//
	// It returns an error in the following scenarios:
	//
	//    • If there are no open sessions available for use,
	//      SetCertificate will return an error that wraps
	//      ErrClosed
	//    • If the KeySlot already contains an certificate,
	//      SetCertificate will return an error that wraps
	//      ErrNotExist
	//    • If an unspecified error occurs, SetCertificate
	//      will return the error without modification
	SetCertificate(KeySlot, *x509.Certificate) error

	// AttestationCertificate attempts to retrieve the keystore's
	// device attestation certificate. This certificate is generated
	// and signed by the keystore's manufacturer, and is stored
	// within a keystore when the device is manufactured.
	//
	AttestationCertificate() (*x509.Certificate, error)

	// Attest generates an *x509.Certificate that is used
	// to provide a non-repudiable declaration of the state
	// keys stored within the specified KeySlot.
	//
	// This *x509.Certificate is used to determine:
	//
	//    • If a key was generated by the keystore, or if it
	//      was generated externally and later imported into
	//      the KeySlot
	//    • The access control policies in place to restrict
	//      unauthorized use of the private key stored within
	//      the KeySlot
	//    • If the keystore is capable of exporting the key
	//      stored within the KeySlot
	//
	// It returns an error in the following scenarios:
	//
	//    • If there are no open sessions available for use,
	//      Attest will return an error that wraps ErrClosed
	//    • If the KeyStore is unable to attest to the state
	//      of a KeySlot, Attest will return an error that
	//      wraps ErrInvalid
	//    • If an unspecified error occurs, Attest will return
	//      the error without modification
	Attest(KeySlot) (*x509.Certificate, error)
}

// VerifyKeyAvailability attempts to retrieve the each key
// in either the public or private components to ensure the
// keystore is open, unlocked, and ready for use.  It returns
// the first error encountered during key retrieval, or nil if
// all keys were retrieved successfully
func VerifyKeyAvailability(backend Keystore, checkPrivateKeys bool) (err error) {
	for keyslot := SigningKeySlot; keyslot < ManagementKeySlot; keyslot++ {
		if checkPrivateKeys {
			if _, err = backend.GetPrivateKey(keyslot); err != nil {
				break
			}
		} else if _, err = backend.GetPublicKey(keyslot); err != nil {
			break
		}
	}
	return
}
