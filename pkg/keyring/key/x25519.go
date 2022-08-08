package key

import (
	"crypto"
	"crypto/rand"
	"crypto/subtle"
	"io"

	"github.com/cloudflare/circl/dh/x25519"
	"github.com/jmg292/G-Net/pkg/gnet"
)

// Implement crypto.PrivateKey for x25519.Key
type X25519PrivateKey struct {
	private x25519.Key
	public  x25519.Key
}

func (key *X25519PrivateKey) Bytes() []byte {
	return append(key.private[:], key.public[:]...)
}

func (key *X25519PrivateKey) PublicBytes() []byte {
	return key.public[:]
}

func (key *X25519PrivateKey) PublicKey() crypto.PublicKey {
	return &X25519PublicKey{
		public: &key.public,
	}
}

func (key *X25519PrivateKey) Equal(x crypto.PrivateKey) (equal bool) {
	if unwrapped, ok := x.(*X25519PrivateKey); ok {
		equal = subtle.ConstantTimeCompare(key.Bytes(), unwrapped.Bytes()) == 1
	}
	return
}

// Implement crypto.PublicKey for x25519.Key
type X25519PublicKey struct {
	public *x25519.Key
}

func (key *X25519PublicKey) Bytes() []byte {
	return key.public[:]
}

func (key *X25519PublicKey) Equal(x crypto.PublicKey) (equal bool) {
	if unwrapped, ok := x.(*X25519PublicKey); ok {
		equal = subtle.ConstantTimeCompare(key.Bytes(), unwrapped.Bytes()) == 1
	}
	return
}

func NewX25519PrivateKey(seed []byte) *X25519PrivateKey {
	var key X25519PrivateKey
	subtle.ConstantTimeCopy(1, seed, key.private[:])
	x25519.KeyGen(&key.public, &key.private)
	return &key
}

func GenerateX25519KeyPair() (*X25519PublicKey, *X25519PrivateKey) {
	var key X25519PrivateKey
	io.ReadFull(rand.Reader, key.private[:])
	x25519.KeyGen(&key.public, &key.private)
	return key.PublicKey().(*X25519PublicKey), &key
}

func DeserializeX25519Key(keyBytes []byte) (*X25519PrivateKey, error) {
	if len(keyBytes) < (x25519.Size * 2) {
		return nil, gnet.ErrorInvalidContentLength
	}
	var key X25519PrivateKey
	subtle.ConstantTimeCopy(1, keyBytes[:x25519.Size], key.private[:])
	subtle.ConstantTimeCopy(1, keyBytes[x25519.Size:], key.public[:])
	return &key, nil
}

func DeserializeX25519PublicKey(keyBytes []byte) (*X25519PublicKey, error) {
	if len(keyBytes) < x25519.Size {
		return nil, gnet.ErrorInvalidContentLength
	}
	var key x25519.Key
	subtle.ConstantTimeCopy(1, keyBytes[:x25519.Size], key[:])
	return &X25519PublicKey{public: &key}, nil
}
