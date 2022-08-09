package yubikey_test

import (
	"fmt"
	"testing"

	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/keyring/backend/yubikey"
)

const resetAllowed bool = true

func TestMain(m *testing.M) {
	yk := yubikey.New()
	if err := yk.Open(); err != nil && err != gnet.ErrorKeystoreNotFound {
		panic(err)
	} else if err != nil && err == gnet.ErrorKeystoreNotFound {
		fmt.Println("Yubikey is not connected.  Skipping tests.")
	} else {
		yk.Close()
		m.Run()
	}
}

func TestKeyGeneration(t *testing.T) {
	if !resetAllowed {
		t.Skip("Reset not authorized.")
	} else {
		t.Run("InvalidTypeCombinations", func(t *testing.T) {
			tests := ResetAuthorized{Test: t, yk: yubikey.New()}
			tests.TestInvalidKeyGeneration()
		})
		t.Run("ValidTypeCombinations", func(t *testing.T) {
			tests := ResetAuthorized{Test: t, yk: yubikey.New()}
			tests.TestValidKeyGeneration()
		})
		t.Run("ErrorKeyAlreadyExists", func(t *testing.T) {
			tests := ResetAuthorized{Test: t, yk: yubikey.New()}
			tests.TestKeyAlreadyExists()
		})
	}
}
