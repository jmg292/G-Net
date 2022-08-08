package file

import (
	"crypto"
	"crypto/x509"
	"path/filepath"

	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/keyring"
	"github.com/jmg292/G-Net/pkg/keyring/kdf"
)

func (f *fileKeyStore) Name() string {
	filename := filepath.Base(f.metadata.Path())
	return filename[:len(filename)-len(filepath.Ext(filename))]
}

func (f *fileKeyStore) Open() error {
	return gnet.ErrorNotYetImplemented
}

func (f *fileKeyStore) Unlock(pin []byte) error {
	return gnet.ErrorNotYetImplemented
}

func (f *fileKeyStore) ManagementKey() (managementKey []byte, err error) {
	return nil, gnet.ErrorNotYetImplemented
}

func (f *fileKeyStore) KeyEncryptionKey() (kek []byte, err error) {
	if managementKey, err := f.ManagementKey(); err == nil {
		kek = kdf.DeriveKey(managementKey, kek)
	}
	return
}

func (f *fileKeyStore) Lock() error {
	return gnet.ErrorNotYetImplemented
}

func (f *fileKeyStore) DestroyKey(slot keyring.KeySlot) error {
	return gnet.ErrorNotYetImplemented
}

func (f *fileKeyStore) CreateKey(keyType keyring.SupportedKeyType, keySlot keyring.KeySlot) error {
	return gnet.ErrorNotYetImplemented
}

func (f *fileKeyStore) GetPublicKey(keySlot keyring.KeySlot) (crypto.PublicKey, error) {
	return nil, gnet.ErrorNotYetImplemented
}

func (f *fileKeyStore) GetPrivateKey(keySlot keyring.KeySlot) (crypto.PrivateKey, error) {
	return nil, gnet.ErrorNotYetImplemented
}

func (f *fileKeyStore) GetPrivateKeyBytes(keySlot keyring.KeySlot) ([]byte, error) {
	return nil, gnet.ErrorNotYetImplemented
}

func (f *fileKeyStore) AttestationCertificate() (*x509.Certificate, error) {
	return nil, gnet.ErrorNotYetImplemented
}

func (f *fileKeyStore) Attest(keyring.KeySlot) (*x509.Certificate, error) {
	return nil, gnet.ErrorNotYetImplemented
}
