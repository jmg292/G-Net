package private

import "github.com/jmg292/G-Net/pkg/identity/public"

type PrivateKeyRing interface {
	Name() string
	Fingerprint() []byte
	PublicKeyRing() public.KeyRing
	Sign([]byte) ([]byte, error)
	Authenticate([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
}
