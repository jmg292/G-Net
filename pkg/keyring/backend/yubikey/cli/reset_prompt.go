package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var resetPrompt []string = []string{
	"Preparing to reset the PIV applet on your Yubikey.\n",
	"If you choose to proceed, all keys and certificates will be wiped from your Yubikey.\n",
	"These keys cannot be regenerated.  This action is irreversible.\n",
	"Would you like to continue? [y/N]: ",
}

func ResetPrompt() (confirmation bool) {
	for _, msg := range resetPrompt {
		fmt.Print(msg)
	}
	reader := bufio.NewReader(os.Stdin)
	if response, _ := reader.ReadString('\n'); strings.HasPrefix(strings.ToLower(response), "y") {
		confirmation = true
		fmt.Println("Reset authorized, proceeding...")
	} else {
		fmt.Println("Reset not authorized, skipping...")
	}
	return
}
