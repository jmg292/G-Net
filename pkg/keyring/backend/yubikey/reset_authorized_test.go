package yubikey_test

import (
	"testing"

	"github.com/awnumar/memguard"
	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/pkg/keyring/backend/yubikey"
)

type ResetAuthorized struct {
	Test *testing.T
	yk   *yubikey.YubikeyStorageBackend
}

func (r *ResetAuthorized) openAndUnlock(pin string) {
	pincopy := make([]byte, len(pin))
	copy(pincopy, []byte(pin))
	if err := r.yk.Open(); err != nil {
		r.Test.Logf("Open failed with error: %s", err)
		r.Test.FailNow()
	}
	if err := r.yk.Unlock(memguard.NewEnclave(pincopy)); err != nil {
		r.Test.Logf("Unlock failed with error: %s", err)
		r.Test.FailNow()
	}
}

func (r *ResetAuthorized) reset() {
	if err := r.yk.Reset(); err != nil {
		r.Test.Logf("Device reset failed with error: %s", err)
		r.Test.FailNow()
	} else {
		r.Test.Logf("Device reset successful")
		if err := r.yk.Close(); err != nil {
			r.Test.Logf("Close failed with error: %s", err)
			r.Test.FailNow()
		} else {
			r.yk = yubikey.New()
			r.openAndUnlock(piv.DefaultPIN)
		}
	}
}
