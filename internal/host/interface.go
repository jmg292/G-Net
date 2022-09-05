package host

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

func GetHost() Host {
	return nil
}

type Host interface {
	Log(string, LogLevel)
}
