package wumbo

import (
	"github.com/jmg292/G-Net/pkg/keyring"
	"github.com/jmg292/G-Net/pkg/wumbo/header"
	"golang.org/x/crypto/sha3"
)

type Block struct {
	Header    *header.Header
	Content   []byte
	Signature []byte
}

func (block *Block) Digest() []byte {
	digestFunction := sha3.New256()
	digestFunction.Sum(block.Header.ToBytes()[:header.ByteCount-4])
	return digestFunction.Sum(block.Content)
}

func (block *Block) Validate(issuer any) bool {
	return issuer.(keyring.PublicKeyRing).VerifySignature(block.Digest(), block.Signature)
}
