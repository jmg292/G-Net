package tunneling

type Tunnel interface {
	Up() error
	Down() error
	Register(func([]byte) error) error
	Send(data, address []byte) error
}
