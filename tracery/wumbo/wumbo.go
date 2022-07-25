package wumbo

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"gnet/identity"
	"gnet/management"

	"golang.org/x/crypto/sha3"
)

type Block struct {
	Header    *WumboHeader
	Content   []byte
	Signature []byte
}

func serializeBlockContent(blockContent any) ([]byte, error) {
	jsonContent, err := json.Marshal(blockContent)
	if err != nil {
		return nil, err
	}
	var serializedContentBuffer bytes.Buffer
	compressor := zlib.NewWriter(&serializedContentBuffer)
	if _, err := compressor.Write(jsonContent); err != nil {
		return nil, err
	}
	return serializedContentBuffer.Bytes(), nil
}

func IssueBlock(precedingBlockId []byte, data any, issuer any) (*Block, error) {
	contentType := data.(management.Command).PacketType()
	blockContent, err := serializeBlockContent(data)
	if err != nil {
		return nil, err
	}
	issuedBlock := Block{
		Header:  NewWumboHeader(precedingBlockId, contentType, len(blockContent), issuer),
		Content: blockContent,
	}
	if issuedBlock.Signature, err = issuer.(identity.Key).Sign(issuedBlock.Digest()); err != nil {
		return nil, err
	}
	return &issuedBlock, nil
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
