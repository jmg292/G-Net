package yubikey_test

import (
	"testing"

	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/pkg/keyring"
	"github.com/jmg292/G-Net/pkg/keyring/backend/yubikey"
)

func TestGetName(t *testing.T) {
	if yk, err := yubikey.New(); err != nil {
		t.Errorf("failed to open yubikey: %s", err)
	} else if name, err := yk.Name(); err != nil {
		t.Errorf("failed to get name: %s", err)
	} else {
		t.Logf("Yubikey name: %s", name)
	}
}

func TestCreateKey(t *testing.T) {
	if yk, err := newOpenAndUnlockedYubikey([]byte(piv.DefaultPIN), t); err == nil {
		t.Logf("Opened Yubikey handle.")
		defer yk.Close()
		for _, params := range keyGenTestParams {
			t.Logf("Attempting to create key of type %d in slot %d.  Success expected? %t", params.Type, params.Slot, params.ExpectSuccess)
			if err := yk.CreateKey(params.Slot, params.Type); err != nil && params.ExpectSuccess {
				t.Errorf("KeyGen of type %d on slot %d failed with error: %s (expected success)", params.Type, params.Slot, err)
			} else if err == nil && !params.ExpectSuccess {
				t.Errorf("KeyGen of type %d on slot %d succeeded. (expected failure)", params.Type, params.Slot)
			} else {
				t.Logf("KeyGen of type %d on slot %d succeeded!", params.Type, params.Slot)
			}
		}
	} else {
		t.Errorf("failed to open yubikey.  Error: %s", err)
	}
}

func TestGetPrivateKey(t *testing.T) {
	if yk, err := newOpenAndUnlockedYubikey([]byte(piv.DefaultPIN), t); err == nil {
		defer yk.Close()
		for i := 0; i < 5; i++ {
			if key, err := yk.GetPrivateKey(keyring.KeySlot(i)); err != nil {
				t.Errorf("failed to get private key for slot %d. error: %s", key, err)
			} else if key == nil {
				t.Logf("got nil key from slot: %d", i)
			} else {
				t.Logf("slot %d passed", i)
			}
		}
	} else {
		t.Errorf("failed to open yubikey.  Error: %s", err)
	}
}

func TestGetPublicKey(t *testing.T) {
	if yk, err := openYubikey(t); err == nil {
		defer yk.Close()
		for i := 0; i < 5; i++ {
			if pubkey, err := yk.GetPublicKey(keyring.KeySlot(i)); err != nil {
				t.Errorf("failed to get public key for slot %d. error: %s", i, err)
			} else if pubkey == nil {
				t.Logf("got nil public key from slot: %d", i)
			} else {
				t.Logf("slot %d passed", i)
			}
		}
	} else {
		t.Errorf("failed to open yubikey.  Error: %s", err)
	}
}

func TestGetCertificate(t *testing.T) {
	if yk, err := openYubikey(t); err == nil {
		defer yk.Close()
		for i := 0; i < 5; i++ {
			if cert, err := yk.GetCertificate(keyring.KeySlot(i)); err != nil {
				t.Errorf("failed to get public key for slot %d. error: %s", i, err)
			} else if cert == nil {
				t.Logf("got nil cert key from slot: %d", i)
			} else {
				t.Logf("slot %d passed", i)
			}
		}
	} else {
		t.Errorf("failed to open yubikey.  error: %s", err)
	}
}

func TestAttestationCertificate(t *testing.T) {
	if yk, err := openYubikey(t); err == nil {
		defer yk.Close()
		if _, err := yk.AttestationCertificate(); err != nil {
			t.Errorf("failed to get attestation certificate. error: %s", err)
		} else {
			t.Logf("passed")
		}
	} else {
		t.Errorf("failed to open yubikey.  error: %s", err)
	}
}

func TestAttest(t *testing.T) {
	if yk, err := openYubikey(t); err == nil {
		defer yk.Close()
		for i := 0; i < 5; i++ {
			if cert, err := yk.Attest(keyring.KeySlot(i)); err != nil {
				t.Errorf("failed to get attestation certificate. error: %s", err)
			} else if cert == nil {
				t.Logf("got nil attestation cert for slot %d", i)
			} else {
				t.Logf("slot %d passed", i)
			}
		}
	}
}
