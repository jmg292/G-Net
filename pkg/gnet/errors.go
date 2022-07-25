package gnet

type ApplicationError string

const (
	ErrorNotYetImplemented ApplicationError = "not yet implemented"
)

type ItemNotFound ApplicationError

const (
	ErrorWarrantNotFound ItemNotFound = "warrant not found"
)

type CryptoError ApplicationError
type ValidationError CryptoError

const (
	ErrorUnsupportedAlgorithm CryptoError     = "unsupported algorithm"
	ErrorInvalidSignature     ValidationError = "invalid signature"
	ErrorInvalidMessageDigest ValidationError = "invalid message digest"
)
