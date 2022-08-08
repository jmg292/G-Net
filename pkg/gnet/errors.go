package gnet

const (
	ErrorNotYetImplemented ApplicationError = "not yet implemented"
	ErrorInvalidPIN        ApplicationError = "invalid pin"
)

const (
	ErrorWarrantNotFound     ItemNotFound = "warrant not found"
	ErrorKeyNotFound         ItemNotFound = "key not found"
	ErrorCertificateNotFound ItemNotFound = "certificate not found"
)

const (
	ErrorInvalidKeySlot                 KeystoreError = "invalid key slot"
	ErrorUnsupportedAlgorithmForKeySlot KeystoreError = "unsupported algorithm for key slot"
	ErrorKeyAlreadyExists               KeystoreError = "key already exists"
	ErrorFileAlreadyExists              KeystoreError = "file already exists"
)

const (
	ErrorInvalidSignature     ValidationError = "invalid signature"
	ErrorInvalidMessageDigest ValidationError = "invalid message digest"
	ErrorInvalidHeader        ValidationError = "invalid header"
	ErrorInvalidContentLength ValidationError = "invalid content length"
	ErrorInvalidSigningKey    ValidationError = "invalid signing key"
	ErrorInvalidPrivateKey    ValidationError = "invalid private key"
	ErrorInvalidPublicKey     ValidationError = "invalid public key"
)

const (
	ErrorKeyExchangeFailed    CryptoError = "key exchange failed"
	ErrorUnsupportedAlgorithm CryptoError = "unsupported algorithm"
)

const (
	ErrorStorageLocked TraceryError = "block storage is locked"
	ErrorBlockExists   TraceryError = "block already exists"
	ErrorBlockNotFound TraceryError = "block not found"
)

const (
	ErrorManifestInvalidSize ManifestError = "invalid manifest size"
)

const (
	ErrorBlockIndexOutOfRange SequencingError = "block index out of range"
	ErrorMalformedEntry       SequencingError = "malformed sequence map entry"
)
