package gnet

type ApplicationError string

const (
	ErrorNotYetImplemented ApplicationError = "not yet implemented"
)

type ItemNotFound ApplicationError

const (
	ErrorWarrantNotFound ItemNotFound = "warrant not found"
)

type ValidationError ApplicationError
type CryptoError ValidationError

const (
	ErrorInvalidSignature     ValidationError = "invalid signature"
	ErrorInvalidMessageDigest ValidationError = "invalid message digest"
	ErrorInvalidHeader        ValidationError = "invalid header"
	ErrorUnsupportedAlgorithm CryptoError     = "unsupported algorithm"
)

type TraceryError ApplicationError
type ManifestError TraceryError

const (
	ErrorStorageLocked        TraceryError  = "block storage is locked"
	ErrorBlockExists          TraceryError  = "block already exists"
	ErrorBlockNotFound        TraceryError  = "block not found"
	ErrorManifestInvalidSize  ManifestError = "invalid manifest size"
	ErrorBlockIndexOutOfRange ManifestError = "block index out of range"
)
