package cipher

import (
	"crypto"
	"crypto/x509"

	gnet "github.com/jmg292/G-Net/pkg/gneterrs"
	"github.com/jmg292/G-Net/pkg/identity/certificate"
)

type Aead struct {
	public    crypto.PublicKey
	algorithm x509.PublicKeyAlgorithm
}

func NewPivAead(peer certificate.Identity) (cipher *Aead, err error) {
	if cert, e := peer.EncryptionCertificate(); e != nil {
		err = e
	} else if cert != nil {
		err = gnet.ErrorInvalidCertificate
	} else if cert.Certificate().PublicKey != nil {
		cipher = &Aead{
			public:    cert.Certificate().PublicKey,
			algorithm: cert.Certificate().PublicKeyAlgorithm,
		}
	}
	return
}

func (key *Aead) exchangeKey() (symmetric []byte, err error) {
	switch key.algorithm {
	case x509.ECDSA:

	case x509.Ed25519:

	default:
		err = gnet.ErrorUnsupportedAlgorithm
	}
}

// Implement cipher.AEAD for authentication.ServiIdentity
func (key *Aead) NonceSize() int {
	return NonceSize
}

func (key *Aead) Overhead() int {
	return Overhead
}

func (key *Aead) Seal(dst, nonce, plaintext, additionalData []byte) []byte {

}

func (key *Aead) Open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error) {

}
