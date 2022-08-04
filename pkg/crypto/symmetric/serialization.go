package signet

import (
	"fmt"

	"github.com/jmg292/G-Net/pkg/crypto"
	"github.com/jmg292/G-Net/pkg/gnet"
	"golang.org/x/crypto/chacha20poly1305"
)

func (s *sealed) toBytes() []byte {
	sealedBytes := append(s.nonce, s.ephemeralPublicKey...)
	return append(sealedBytes, s.content...)
}

func sealedFromBytes(data []byte) (*sealed, error) {
	if len(data) <= int(crypto.X25519KeySize)+chacha20poly1305.NonceSizeX {
		return nil, fmt.Errorf(string(gnet.ErrorInvalidContentLength))
	}
	return &sealed{
		nonce:              data[:chacha20poly1305.NonceSizeX],
		ephemeralPublicKey: data[chacha20poly1305.NonceSizeX:crypto.X25519KeySize],
		content:            data[chacha20poly1305.NonceSizeX+crypto.X25519KeySize:],
	}, nil
}
