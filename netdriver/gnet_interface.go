package netdriver

import (
	"fmt"

	"golang.org/x/sys/windows"
)

type GNetInterface struct {
	handle uintptr
}

func NewGNetInterface() (*GNetInterface, bool, error) {
	_, _, err := pSetLogger.Call(windows.NewCallback(writeLog))
	if err != nil {
		return nil, false, fmt.Errorf("log init: %s", err)
	}
	return nil, false, nil
}
