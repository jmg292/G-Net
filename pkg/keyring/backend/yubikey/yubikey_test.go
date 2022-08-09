package yubikey_test

import (
	"fmt"
	"testing"

	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/keyring/backend/yubikey"
	"github.com/jmg292/G-Net/pkg/keyring/backend/yubikey/cli"
)

const resetAllowed bool = false

var yk *yubikey.YubikeyStorageBackend = yubikey.New()

func resetIfAllowed() (reset bool, err error) {
	if !resetAllowed || !cli.ResetPrompt() {
		err = gnet.ErrorResetNotAllowed
	} else {
		if err = yk.Reset(); err == nil {
			reset = true
		}
	}
	return
}

func TestMain(m *testing.M) {
	if err := yk.Open(); err != nil && err != gnet.ErrorKeystoreNotFound {
		panic(err)
	} else if err != nil && err == gnet.ErrorKeystoreNotFound {
		fmt.Println("Yubikey is not connected.  Skipping tests.")
	} else {
		defer yk.Close()
		var shouldTestKeyGeneration, shouldTestKeyRetrieval bool
		if wasReset, err := resetIfAllowed(); err != nil && err != gnet.ErrorResetNotAllowed {
			fmt.Printf("Error while attempting reset: %s\n", err)
		} else if err == nil && !wasReset {
			fmt.Println("Reset was authorized, but device was not reset")
		} else if err != nil && err == gnet.ErrorResetNotAllowed {
			shouldTestKeyRetrieval = true
		} else {
			shouldTestKeyGeneration = true
			shouldTestKeyRetrieval = true
		}
	}
}
