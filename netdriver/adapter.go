package netdriver

// Functions dealing with the `Adapter` object
var (
	pCreateAdapter  = tunnelDriver.NewProc("WintunCreateAdapter")
	pOpenAdapter    = tunnelDriver.NewProc("WintunOpenAdapter")
	pCloseAdapter   = tunnelDriver.NewProc("WintunCloseAdapter")
	pGetAdapterLuid = tunnelDriver.NewProc("WintunGetAdapterLuid")
)