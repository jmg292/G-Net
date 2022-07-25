package identity

type Key interface {
	Open() error
	Close() error
	GenerateKey() error
	Fingerprint() []byte
	Certificate() ([]byte, error)
	Sign([]byte) ([]byte, error)
}

type PublicKey interface {
	Fingerprint() []byte
	Verify([]byte, []byte) error
	Encrypt([]byte) ([]byte, error)
}
