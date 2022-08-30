package key

import (
	"testing"

	"github.com/jmg292/G-Net/pkg/keyring"
)

func testFunctionalityForKeyType(keyType keyring.SupportedKeyType, t *testing.T) {
	public, private, err := GenerateKeyPair(keyType)
	if err != nil {
		t.Errorf("failed to generate key.  Type: %d; Error: %s", keyType, err.Error())
	}
	if publicBytes, err := PublicToBytes(keyType, public); err != nil {
		t.Errorf("failed to serialize public key.  Type: %d; Error: %s", keyType, err.Error())
	} else {
		t.Logf("public bytes: 0x%x", publicBytes)
	}
	if privateBytes, err := PrivateToBytes(keyType, private); err != nil {
		t.Errorf("failed to serialize private key. Type: %d, Error: %s", keyType, err.Error())
	} else {
		t.Logf("private bytes: 0x%x", privateBytes)
	}
}

func TestP256KeyFunctionality(t *testing.T) {
	testFunctionalityForKeyType(keyring.EC256Key, t)
}

func TestP384KeyFunctionality(t *testing.T) {
	testFunctionalityForKeyType(keyring.EC384Key, t)
}

func TestX25519KeyFunctionality(t *testing.T) {
	testFunctionalityForKeyType(keyring.X25519Key, t)
}
