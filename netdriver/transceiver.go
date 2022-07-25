package netdriver

// Functions dealing with packet reception
var (
	pGetReadWaitEvent     = tunnelDriver.NewProc("WintunGetReadWaitEvent")
	pReceivePacket        = tunnelDriver.NewProc("WintunReceivePacket")
	pReleaseReceivePacket = tunnelDriver.NewProc("WintunReleaseReceivePacket")
)

// Functions dealing with packet transmission
var (
	pAllocateSendPacket = tunnelDriver.NewProc("WintunAllocateSendPacket")
	pSendPacket         = tunnelDriver.NewProc("WintunSendPacket")
)