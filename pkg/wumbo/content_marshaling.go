package wumbo

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
)

func MarshalContent(blockContent any) ([]byte, error) {
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

func UnmarshalContent(blockContent []byte, v any) error {
	var deserializedContentBuffer bytes.Buffer
	decompressor := zlib.NewWriter(&deserializedContentBuffer)
	decompressor.Write(blockContent)
	return json.Unmarshal(deserializedContentBuffer.Bytes(), v)
}
