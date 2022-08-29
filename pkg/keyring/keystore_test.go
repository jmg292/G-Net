package keyring_test

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/x509"
	"errors"

	"github.com/jmg292/G-Net/pkg/keyring"
)

var ErrNotImplemented = errors.New("not implemented")

// Dummy implementation of keyring.KeyInfo
type testKeyInfo struct {
	slot keyring.KeySlot
}

// Dummy implementation of
func (info *testKeyInfo) KeySlot() keyring.KeySlot {
	return info.slot
}

func (info *testKeyInfo) KeyType() keyring.KeyType {
	return keyring.NilKey
}

func (info *testKeyInfo) KeyIdentifier() []byte {
	return []byte("Test Key")
}

func (info *testKeyInfo) KeySecurityPolicy() []byte {
	return []byte("Test Key")
}

type testKeyStore struct {
	privateKey *ecdsa.PrivateKey
}

func (*testKeyStore) KeyInfo(slot keyring.KeySlot) (keyring.KeyInfo, error) {
	return &testKeyInfo{slot: slot}, nil
}

func (*testKeyStore) Open() error {
	// Not used in this package
	return ErrNotImplemented
}

func (*testKeyStore) Unlock([]byte) error {
	// Not used in this package
	return ErrNotImplemented
}

func (*testKeyStore) Lock() error {
	// Not used in this package
	return ErrNotImplemented
}
func (*testKeyStore) Close() error {
	// Not used in this package
	return ErrNotImplemented
}
func (*testKeyStore) CreateKey(_ keyring.KeySlot, _ keyring.KeyType) error {
	// Not used in this package
	return ErrNotImplemented
}
func (t *testKeyStore) GetPrivateKey(_ keyring.KeySlot) (crypto.PrivateKey, error) {
	return t.privateKey, nil
}
func (t *testKeyStore) GetPublicKey(_ keyring.KeySlot) (crypto.PublicKey, error) {
	return t.privateKey.PublicKey, nil
}

func (*testKeyStore) GetCertificate(_ keyring.KeySlot) (*x509.Certificate, error) {
	// Not used in this package
	return nil, ErrNotImplemented
}

func (*testKeyStore) SetCertificate(_ keyring.KeySlot, _ *x509.Certificate) error {
	// Not used in this package
	return ErrNotImplemented
}

func (*testKeyStore) AttestationCertificate() (*x509.Certificate, error) {
	// Not used in this package
	return nil, ErrNotImplemented
}

func (*testKeyStore) Attest(_ keyring.KeySlot) (*x509.Certificate, error) {
	// Not used in this package
	return nil, ErrNotImplemented
}
