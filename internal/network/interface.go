package network

type Interface interface {
	Name() (string, error)
	Up() error
	Down() error
	Register(func([]byte) error) error
	Send(data, address []byte) error
}
