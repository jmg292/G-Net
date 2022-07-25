package netdriver

import "syscall"

var tunnelDriver = syscall.NewLazyDLL("wintun.dll")

// Functions dealing with logging
var (
	pSetLogger     = tunnelDriver.NewProc("WintunSetLogger")	
)
