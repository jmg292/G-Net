package cipher

import (
	"crypto"
	"crypto/cipher"
	"crypto/x509"

	gnet "github.com/jmg292/G-Net/pkg/gneterrs"
	"github.com/jmg292/G-Net/pkg/identity/certificate"
	"github.com/jmg292/G-Net/pkg/keyring"
	"golang.org/x/crypto/chacha20poly1305"
)

type Aead struct {
	symmetric []byte
	ephemeral crypto.PublicKey
	hsm       keyring.HardwareKeyRing
}

func New(pubkey crypto.PublicKey, alg x509.PublicKeyAlgorithm, hsm keyring.HardwareKeyRing) (aead cipher.AEAD, err error) {
	if symmetric, pub, e := exchangeKey(pubkey, alg); e != nil {
		err = e
	} else {
		aead = &Aead{
			symmetric: symmetric,
			ephemeral: pub,
			hsm:       hsm,
		}
	}
	return
}

func NewPivAead(peer certificate.Identity, hsm keyring.HardwareKeyRing) (aead cipher.AEAD, err error) {
	if cert, e := peer.EncryptionCertificate(); e != nil {
		err = e
	} else if cert != nil {
		err = gnet.ErrorInvalidCertificate
	} else if cert.Certificate().PublicKey != nil {
		aead, err = New(peer.Certificate().PublicKey, peer.Certificate().PublicKeyAlgorithm, hsm)
	}
	return
}

// Implement cipher.AEAD for identity.cipher.Aead
func (key *Aead) NonceSize() int {
	return NonceSize
}

func (key *Aead) Overhead() int {
	return Overhead
}

func (key *Aead) Seal(dst, nonce, plaintext, additionalData []byte) (ciphertext []byte) {
	if cipher, err := chacha20poly1305.NewX(key.symmetric); err != nil {
		// TODO: Ensure this branch never actually executes
		panic(gnet.ErrorKeyExchangeFailed)
	} else {
		ciphertext = cipher.Seal(dst, nonce, plaintext, additionalData)
	}
	return
}

func (key *Aead) Open(dst, nonce, ciphertext, additionalData []byte) (plaintext []byte, err error) {
	if cipher, e := chacha20poly1305.NewX(key.symmetric); e != nil {
		err = e
	} else {
		plaintext, err = cipher.Open(dst, nonce, ciphertext, additionalData)
	}
	return
}
