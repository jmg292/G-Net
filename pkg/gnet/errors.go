package gnet

const (
	ErrorNotYetImplemented ApplicationError = "not yet implemented"
	ErrorInvalidPIN        ApplicationError = "invalid pin"
)

const (
	ErrorWarrantNotFound ItemNotFound = "warrant not found"
)

const (
	ErrorInvalidSignature     ValidationError = "invalid signature"
	ErrorInvalidMessageDigest ValidationError = "invalid message digest"
	ErrorInvalidHeader        ValidationError = "invalid header"
	ErrorInvalidContentLength ValidationError = "invalid content length"
)

const (
	ErrorInvalidSigningKey              CryptoError = "invalid signing key"
	ErrorInvalidPrivateKey              CryptoError = "invalid private key"
	ErrorInvalidPublicKey               CryptoError = "invalid public key"
	ErrorInvalidKeySlot                 CryptoError = "invalid key slot"
	ErrorKeyAlreadyExists               CryptoError = "key already exists"
	ErrorKeyNotFound                    CryptoError = "key not found"
	ErrorCertificateNotFound            CryptoError = "certificate not found"
	ErrorKeyExchangeFailed              CryptoError = "key exchange failed"
	ErrorUnsupportedAlgorithm           CryptoError = "unsupported algorithm"
	ErrorUnsupportedAlgorithmForKeySlot CryptoError = "unsupported algorithm for key slot"
)

const (
	ErrorStorageLocked        TraceryError    = "block storage is locked"
	ErrorBlockExists          TraceryError    = "block already exists"
	ErrorBlockNotFound        TraceryError    = "block not found"
	ErrorManifestInvalidSize  ManifestError   = "invalid manifest size"
	ErrorBlockIndexOutOfRange SequencingError = "block index out of range"
	ErrorMalformedEntry       SequencingError = "malformed sequence map entry"
)
