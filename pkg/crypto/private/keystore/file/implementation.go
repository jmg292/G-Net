package file

import (
	"crypto"
	"crypto/x509"
	"fmt"
	"path/filepath"

	gcrypt "github.com/jmg292/G-Net/pkg/crypto"
	"github.com/jmg292/G-Net/pkg/crypto/public"
	"github.com/jmg292/G-Net/pkg/gnet"
)

func (f *fileKeyStore) Name() string {
	return filepath.Base(f.path)
}

func (f *fileKeyStore) GetKeyId(keyId gcrypt.KeySlot) ([]byte, error) {
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) Unlock(pin []byte) error {
	return fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) Validate() error {
	return fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) ManagementKey() (managementKey []byte, err error) {
	if f.managementKey == nil {
		err = fmt.Errorf(string(gnet.ErrorKeyNotFound))
	} else {
		managementKey = f.managementKey
	}
	return
}

func (f *fileKeyStore) KeyEncryptionKey() ([]byte, error) {
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) Lock() error {
	return fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) DestroyKey(slot gcrypt.KeySlot) error {
	return fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) CreateKey(keyType gcrypt.SupportedKeyType, keySlot gcrypt.KeySlot) error {
	return fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) GetPublicKey(keySlot gcrypt.KeySlot) (crypto.PublicKey, error) {
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) PublicKeyRing() (public.KeyRing, error) {
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) GetPrivateKey(keySlot gcrypt.KeySlot) (crypto.PrivateKey, error) {
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) GetPrivateKeyBytes(keySlot gcrypt.KeySlot) ([]byte, error) {
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) AttestationCertificate() (*x509.Certificate, error) {
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}

func (f *fileKeyStore) Attest(gcrypt.KeySlot) (*x509.Certificate, error) {
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}
