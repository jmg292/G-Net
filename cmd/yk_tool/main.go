package main

import (
	"fmt"
	"syscall"

	"github.com/jmg292/G-Net/pkg/crypto/private/keyring/yubikey"
	"github.com/jmg292/G-Net/pkg/gnet"
	"golang.org/x/term"
)

const (
	// Pin constraints: https://developers.yubico.com/yubikey-piv-manager/PIN_and_Management_Key.html
	PinMinLength int = 6
	PinMaxLength int = 8
)

func validatePin(pin string) error {
	if len(pin) < PinMinLength || len(pin) > PinMaxLength {
		return gnet.ErrorInvalidPIN
	}
	return nil
}

func getPin() (string, error) {
	fmt.Print("Enter PIN: ")
	pinBytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	return string(pinBytes), validatePin(string(pinBytes))
}

func main() {
	var pin string
	var err error
	fmt.Println("Your PIN is required to unlock your Yubikey.")
	for pin == "" || err != nil {
		pin, err = getPin()
		if err != nil {
			fmt.Println(err)
		}
	}
	yubikey.NewYubikeyIdentityProvider(pin)
}
