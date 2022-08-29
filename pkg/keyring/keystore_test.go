package keyring_test

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"errors"
	"io"

	"github.com/jmg292/G-Net/pkg/keyring"
)

var ErrNotImplemented = errors.New("not implemented")

// Dummy implementation of keyring.KeyInfo
type testKeyInfo struct {
	slot keyring.KeySlot
}

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

// Dummy implementation of crypto.Decrypter
type testDecryptWrapper struct {
	key *ecdsa.PrivateKey
}

func (t *testDecryptWrapper) Public() crypto.PublicKey {
	return t.key.Public()
}

func (t *testDecryptWrapper) Decrypt(rand io.Reader, msg []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error) {
	return nil, nil
}

// Dummy implementation of KeyStore
type testKeyStore struct {
	privateKey *ecdsa.PrivateKey
}

func NewTestKeyStore() *testKeyStore {
	key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return &testKeyStore{privateKey: key}
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
func (t *testKeyStore) GetPrivateKey(slot keyring.KeySlot) (key crypto.PrivateKey, err error) {
	if slot == keyring.EncryptionKeySlot {
		key = &testDecryptWrapper{key: t.privateKey}
	} else {
		key = t.privateKey
	}
	return
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
