package keyring_test

import (
	"crypto/rand"
	"testing"

	"github.com/jmg292/G-Net/pkg/keyring"
)

func getTestByteSlice(length int) (test []byte) {
	test = make([]byte, length)
	rand.Read(test)
	return
}

func TestPublic(t *testing.T) {
	if keyring, err := keyring.NewPrivate(NewTestKeyStore()); err != nil {
		t.Errorf("Instantiation failed: %s", err)
	} else if public := keyring.Public(); public == nil {
		t.Error("nil public key")
	}
}

func TestSign(t *testing.T) {
	testHashSize := 64 // No need to support more than 512 bits right now
	if keyring, err := keyring.NewPrivate(NewTestKeyStore()); err != nil {
		t.Errorf("Instantiation failed: %s", err)
	} else if _, err := keyring.Sign(rand.Reader, getTestByteSlice(testHashSize), nil); err != nil {
		t.Errorf("Sign failed: %s", err)
	}
}

func TestDerypt(t *testing.T) {
	testDataSize := 64 // Arbitrary size, data is unused by test backends
	if keyring, err := keyring.NewPrivate(NewTestKeyStore()); err != nil {
		t.Errorf("Instantiation failed: %s", err)
	} else if _, err := keyring.Decrypt(rand.Reader, getTestByteSlice(testDataSize), nil); err != nil {
		t.Errorf("Decrypt failed: %s", err)
	}
}
