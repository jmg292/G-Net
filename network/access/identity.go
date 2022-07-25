package access

import "crypto/x509"

type IdentityType uint8

const (
	Administrator IdentityType = iota
	Device
	User
)

type Identity struct {
	Type                    IdentityType
	ParentId                []byte
	Name                    string
	SignatoryPublicKey      x509.Certificate
	AuthenticationPublicKey x509.Certificate
}
