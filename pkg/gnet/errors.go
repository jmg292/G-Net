package gnet

const (
	ErrorNotYetImplemented ApplicationError = "not yet implemented"
)

const (
	ErrorWarrantNotFound     ItemNotFound = "warrant not found"
	ErrorKeyNotFound         ItemNotFound = "key not found"
	ErrorCertificateNotFound ItemNotFound = "certificate not found"
)

const (
	ErrorUnsupportedAlgorithmForKeySlot KeystoreError = "unsupported algorithm for key slot"
	ErrorExportNotAllowed               KeystoreError = "export not allowed"
	ErrorImportNotAllowed               KeystoreError = "import not allowed"
	ErrorKeyAlreadyExists               KeystoreError = "key already exists"
	ErrorFileAlreadyExists              KeystoreError = "file already exists"
	ErrorUnableToOpenKeystore           KeystoreError = "unable to open keystore"
	ErrorKeystoreNotFound               KeystoreError = "keystore not found"
	ErrorKeystoreLocked                 KeystoreError = "keystore is locked"
)

const (
	ErrorInvalidSignature     ValidationError = "invalid signature"
	ErrorInvalidMessageDigest ValidationError = "invalid message digest"
	ErrorInvalidHeader        ValidationError = "invalid header"
	ErrorInvalidContentLength ValidationError = "invalid content length"
	ErrorInvalidSigningKey    ValidationError = "invalid signing key"
	ErrorInvalidPrivateKey    ValidationError = "invalid private key"
	ErrorInvalidPublicKey     ValidationError = "invalid public key"
	ErrorInvalidPIN           ValidationError = "invalid pin"
	ErrorInvalidKeySlot       ValidationError = "invalid key slot"
	ErrorInvalidHandle        ValidationError = "invalid handle"
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
