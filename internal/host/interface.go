package host

import gnet "github.com/jmg292/G-Net/pkg/gneterrs"

type LogLevel uint8

const (
	Unknown LogLevel = iota
	Trace
	Debug
	Info
	Warn
	Error
	Fatal
	Panic
)

func GetHost() (Host, error) {
	return nil, gnet.ErrorNotYetImplemented
}

type Host interface {
	Log(string, LogLevel)
}
