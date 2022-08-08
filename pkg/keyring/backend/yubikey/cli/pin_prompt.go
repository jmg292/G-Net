package cli

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	"github.com/jmg292/G-Net/pkg/gnet"
)

const (
	// https://developers.yubico.com/yubikey-piv-manager/PIN_and_Management_Key.html
	pinMinLength = 6
	pinMaxLength = 8
)

func assertOnlyAsciiCharacters(s string) (err error) {
	// https://stackoverflow.com/a/53069799
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			err = fmt.Errorf(gnet.ErrorInvalidCharacter.Error(), s[i])
			break
		}
	}
	return
}

func PINPrompt() (pin string, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Your PIN is required to authorize this action.")
	fmt.Printf("Please enter your PIN: ")
	if pin, err = reader.ReadString('\n'); err != nil {
		pin = ""
	} else if len(pin) < pinMinLength || len(pin) > pinMaxLength {
		pin = ""
		err = gnet.ErrorInvalidPIN
	} else if err = assertOnlyAsciiCharacters(pin); err != nil {
		pin = ""
	}
	return
}
