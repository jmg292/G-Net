package netdriver

// Functions dealing with the `Session` object
var (
	pStartSession = tunnelDriver.NewProc("WintunStartSession")
	pEndSession   = tunnelDriver.NewProc("WintunEndSession")
)