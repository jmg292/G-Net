package yubikey_test

import (
	"testing"

	"github.com/go-piv/piv-go/piv"
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
	return
}

func TestCreateKey(t *testing.T) {
	if yk, err := newOpenAndUnlockedYubikey([]byte(piv.DefaultPIN), t); err == nil {
		defer yk.Close()
		for _, params := range keyGenTestParams {
			if err := yk.CreateKey(params.Slot, params.Type); err != nil && params.ExpectSuccess {
				t.Errorf("KeyGen of type %d on slot %d failed with error: %s (expected success)", params.Type, params.Slot, err)
			} else if err == nil && !params.ExpectSuccess {
				t.Errorf("KeyGen of type %d on slot %d succeeded. (expected failure)", params.Type, params.Slot)
			} else {
				t.Logf("KeyGen of type %d on slot %d succeeded!", params.Type, params.Slot)
			}
		}
	}
}
