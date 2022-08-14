package yubikey_test

import (
	"testing"

	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/pkg/keyring/backend/yubikey"
)

const pin []byte = []byte(piv.DefaultPIN)

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

	}
}
