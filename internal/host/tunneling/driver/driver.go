package driver

// Functions dealing with the `Driver` object
var (
	pDeleteDriver            = tunnelDriver.NewProc("WintunDeleteDriver")
	pGetRunningDriverVersion = tunnelDriver.NewProc("WintunGetRunningDriverVersion")
)
