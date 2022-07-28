package wumbo

import (
	"bytes"
	"compress/zlib"
	"encoding/json"

	"github.com/jmg292/G-Net/internal/datagrams"
	"github.com/jmg292/G-Net/pkg/identity/private"
)

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
	contentType := data.(datagrams.Datagram).Type()
	blockContent, err := serializeBlockContent(data)
	if err != nil {
		return nil, err
	}
	issuedBlock := Block{
		Header:  NewWumboHeader(precedingBlockId, contentType, len(blockContent), issuer),
		Content: blockContent,
	}
	if issuedBlock.Signature, err = issuer.(private.KeyRing).Sign(issuedBlock.Digest()); err != nil {
		return nil, err
	}
	return &issuedBlock, nil
}
