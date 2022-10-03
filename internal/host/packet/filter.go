package packet

import (
	"sync"

	"github.com/williamfhe/godivert"
)

type Filter struct {
	handle *godivert.WinDivertHandle
}

var (
	instance      *Filter
	instanceMutex = &sync.Mutex{}
)

func NewFilter() (filter *Filter, err error) {
	instanceMutex.Lock()
	defer instanceMutex.Unlock()

	if instance == nil {
		// TODO: Do not divert network i/o from G-Net services
		if handle, e := godivert.NewWinDivertHandle(""); e != nil {
			err = e
		} else {
			instance = &Filter{handle: handle}
		}
	}

	filter = instance
	return
}
