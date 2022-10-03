package main

import (
	"fmt"
	"time"

	"github.com/jmg292/G-Net/internal/host/packet"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/debug"
)

var eventLog debug.Log

type svcGuardian struct {
	filter *packet.Filter
}

// Execute implements Windows service registration and control functionality for G-Net Guardian
// (see: https://pkg.go.dev/golang.org/x/sys/windows/svc#Handler)
func (guardian *svcGuardian) Execute(args []string, chanRequest <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	// TODO: Guardian setup code here
	tick := time.Tick(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			// TODO: Update packet filter statistics
		case req := <-chanRequest:
			switch req.Cmd {
			case svc.Interrogate:
				changes <- req.CurrentStatus
				time.Sleep(100 * time.Millisecond)
				changes <- req.CurrentStatus
			case svc.Stop:
				eventLog.Warning(1, "Guardian service ignored Stop request from Windows service manager.")
			case svc.Shutdown:
				eventLog.Info(1, "Guardian service shutting down.")
				changes <- svc.Status{State: svc.StopPending}
				return
			case svc.Pause:
				eventLog.Warning(1, "Guardian service ignored Pause request from Windows service manager.")
			default:
				eventLog.Warning(1, fmt.Sprintf("Guardian service received unexpected control request #%d", req))
			}
		}
	}
}
