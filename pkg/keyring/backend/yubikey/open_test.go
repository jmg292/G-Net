package yubikey_test

import (
	"testing"

	"github.com/jmg292/G-Net/pkg/keyring/backend/yubikey"
)

func newOpenAndUnlockedYubikey(pin []byte, t *testing.T) (yk *Yubikey, err error) {
	if yk, err = yubikey.New(); err != nil {
		t.Errorf("failed to get new yubikey: %s", err)
	} else if err = yk.Open(); err != nil {
		t.Errorf("failed to open yubikey: %s", err)
	} else if err = yk.Unlock(pin); err != nil {
		t.Errorf("failed to unlock yubikey: %s", err)
	}
	return
}
