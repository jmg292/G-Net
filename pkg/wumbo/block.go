package wumbo

import (
	"bytes"
	"compress/zlib"
	"encoding/json"

	"github.com/jmg292/G-Net/identity"
	"golang.org/x/crypto/sha3"
)

type Block struct {
	Header    *WumboHeader
	Content   []byte
	Signature []byte
}

func (block *Block) Digest() []byte {
	digestFunction := sha3.New256()
	digestFunction.Sum(block.Header.AsBytes()[:WumboHeaderByteCount-4])
	return digestFunction.Sum(block.Content)
}

func (block *Block) Validate(issuer any) error {
	return issuer.(identity.PublicKey).Verify(block.Digest(), block.Signature)
}

func (block *Block) UnmarshalContent(v any) error {
	var deserializedContentBuffer bytes.Buffer
	decompressor := zlib.NewWriter(&deserializedContentBuffer)
	decompressor.Write(block.Content)
	return json.Unmarshal(deserializedContentBuffer.Bytes(), v)
}
