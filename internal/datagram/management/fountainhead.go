package management

import "github.com/go-piv/piv-go/piv"

type Fountainhead struct {
	DomainName              string
	FounderName             string
	SigningKey              []byte
	SignatureAlgorithm      piv.Algorithm
	AuthenticationKey       []byte
	AuthenticationAlgorithm piv.Algorithm
}

func (*Fountainhead) Type() Datagram {
	return NetworkRoot
}
