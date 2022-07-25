package message

import (
	"bytes"
	"compress/zlib"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"gnet/control/protocol/payload"
	"gnet/identity"
)

type messageFormatter struct {
	signingKey identity.Key
}

func NewMessageFormatter(signingKey any) *messageFormatter {
	return &messageFormatter{
		signingKey: signingKey.(identity.Key),
	}
}

func (v *messageFormatter) Seal(message any, recipient any) (*SealedMessage, error) {

	// Generate a header for this message
	messageHeader := CreateHeader(v.signingKey.Fingerprint(), message.(payload.Content).Type())
	rand.Read(messageHeader.EncryptionKey)
	rand.Read(messageHeader.Nonce)

	// Marshal message to bytes and compress it
	var compressionBuffer bytes.Buffer
	compressor := zlib.NewWriter(&compressionBuffer)
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	compressor.Write(messageBytes)

	// Encrypt the compressed message bytes
	encryptedPayload, err := encryptPayload(messageHeader.EncryptionKey, messageHeader.Nonce, compressionBuffer.Bytes())
	if err != nil {
		return nil, err
	}
	compressionBuffer.Reset()

	// Marshal header to bytes
	headerBytes, err := json.Marshal(messageHeader)
	if err != nil {
		return nil, err
	}

	// Compress and encrypt marshalled header bytes
	compressor.Write(headerBytes)
	encryptedHeader, err := recipient.(identity.PublicKey).Encrypt(compressionBuffer.Bytes())
	if err != nil {
		return nil, err
	}

	// Sign (encryptedHeader | encryptedPayload)
	signature, err := v.signingKey.Sign(append(encryptedHeader, encryptedPayload...))
	if err != nil {
		return nil, err
	}

	// Return the sealed message
	return &SealedMessage{
		Sender:    base64.StdEncoding.EncodeToString(v.signingKey.Fingerprint()),
		Payload:   base64.StdEncoding.EncodeToString(encryptedPayload),
		Header:    base64.StdEncoding.EncodeToString(encryptedHeader),
		Signature: base64.StdEncoding.EncodeToString(signature),
	}, nil
}
