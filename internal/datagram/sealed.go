package datagram

type Sealed struct {
	EphemeralKey []byte
	Payload      []byte
	Signature    []byte
}
