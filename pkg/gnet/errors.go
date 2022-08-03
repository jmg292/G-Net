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
	ErrorInvalidContentLength ValidationError = "invalid content length"
)

type CryptoError ValidationError

const (
	ErrorInvalidSigningKey              CryptoError = "invalid signing key"
	ErrorInvalidPublicKey               CryptoError = "invalid public key"
	ErrorInvalidKeySlot                 CryptoError = "invalid key slot"
	ErrorKeyAlreadyExists               CryptoError = "key already exists"
	ErrorKeyNotFound                    CryptoError = "key not found"
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
