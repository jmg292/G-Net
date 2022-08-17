package yubikey_test

import (
	"testing"

	"github.com/jmg292/G-Net/pkg/keyring/backend/yubikey"
)

func openYubikey(t *testing.T) (yk *yubikey.Yubikey, err error) {
	if yk, err = yubikey.New(); err != nil {
		t.Errorf("failed to get new yubikey: %s", err)
	}
	return
}

func newOpenAndUnlockedYubikey(pin []byte, t *testing.T) (yk *yubikey.Yubikey, err error) {
	if yk, err = openYubikey(t); err == nil {
		err = yk.Unlock(pin)
	}
	return
}
