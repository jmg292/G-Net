package gnet

type ApplicationError string

const (
	ErrorNotYetImplemented ApplicationError = "not yet implemented"
	ErrorInvalidPIN        ApplicationError = "invalid pin"
)

type ItemNotFound ApplicationError

const (
	ErrorWarrantNotFound ItemNotFound = "warrant not found"
)

type ValidationError ApplicationError

const (
	ErrorInvalidSignature     ValidationError = "invalid signature"
	ErrorInvalidMessageDigest ValidationError = "invalid message digest"
	ErrorInvalidHeader        ValidationError = "invalid header"
)

type CryptoError ValidationError

const (
	ErrorInvalidKeySlot                 CryptoError = "invalid key slot"
	ErrorKeyAlreadyExists               CryptoError = "key already exists"
	ErrorUnsupportedAlgorithm           CryptoError = "unsupported algorithm"
	ErrorUnsupportedAlgorithmForKeySlot CryptoError = "unsupported algorithm for key slot"
)

type TraceryError ApplicationError
type ManifestError TraceryError
type SequencingError TraceryError

const (
	ErrorStorageLocked        TraceryError    = "block storage is locked"
	ErrorBlockExists          TraceryError    = "block already exists"
	ErrorBlockNotFound        TraceryError    = "block not found"
	ErrorManifestInvalidSize  ManifestError   = "invalid manifest size"
	ErrorBlockIndexOutOfRange SequencingError = "block index out of range"
	ErrorMalformedEntry       SequencingError = "malformed sequence map entry"
)
