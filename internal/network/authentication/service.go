package authentication

import "github.com/jmg292/G-Net/pkg/keyring"

// authentication.Service
type Service struct {
	keyring *keyring.HardwareKeyRing
}
