package file

import (
	"crypto"
	"crypto/x509"
	"fmt"
	"path/filepath"

	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/keyring"
	"github.com/jmg292/G-Net/pkg/keyring/kdf"
)

func (f *fileKeyStore) Name() string {
	return filepath.Base(f.metadata.Path())
}

func (f *fileKeyStore) GetKeyId(keyId keyring.KeySlot) ([]byte, error) {
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) Unlock(pin []byte) error {
	return fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) Validate() error {
	return fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) ManagementKey() (managementKey []byte, err error) {
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) KeyEncryptionKey() (kek []byte, err error) {
	if managementKey, err := f.ManagementKey(); err == nil {
		kek = kdf.DeriveKey(managementKey, kek)
	}
	return
}

func (f *fileKeyStore) Lock() error {
	return fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) DestroyKey(slot keyring.KeySlot) error {
	return fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) CreateKey(keyType keyring.SupportedKeyType, keySlot keyring.KeySlot) error {
	return fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) GetPublicKey(keySlot keyring.KeySlot) (crypto.PublicKey, error) {
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) GetPrivateKey(keySlot keyring.KeySlot) (crypto.PrivateKey, error) {
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) GetPrivateKeyBytes(keySlot keyring.KeySlot) ([]byte, error) {
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) AttestationCertificate() (*x509.Certificate, error) {
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) Attest(keyring.KeySlot) (*x509.Certificate, error) {
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}
