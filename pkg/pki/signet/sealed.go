package signet

import (
	"crypto"
	"crypto/rand"
	"fmt"

	"github.com/jmg292/G-Net/pkg/gnet"
	"golang.org/x/crypto/chacha20poly1305"
)

type sealed struct {
	nonce              []byte
	ephemeralPublicKey []byte
	content            []byte
}

func newSealedContent(data []byte, publicKey crypto.PublicKey) (*sealed, error) {
	s := sealed{
		nonce:   make([]byte, chacha20poly1305.NonceSizeX),
		content: make([]byte, len(data)+chacha20poly1305.Overhead),
	}
	if _, err := rand.Read(s.nonce); err != nil {
		return nil, err
	}
	if publicKeyBytes, ok := publicKey.([]byte); !ok {
		return nil, fmt.Errorf(string(gnet.ErrorInvalidPublicKey))
	} else {
		s.ephemeralPublicKey = publicKeyBytes
	}
	return &s, nil
}
