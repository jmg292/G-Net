package gnet

const (
	ErrorNotYetImplemented ApplicationError = "not yet implemented"
)

const (
	ErrorWarrantNotFound     ItemNotFound = "warrant not found"
	ErrorKeyNotFound         ItemNotFound = "key not found"
	ErrorCertificateNotFound ItemNotFound = "certificate not found"
	ErrorKeystoreNotFound    ItemNotFound = "keystore not found"
)

const (
	ErrorUnsupportedAlgorithmForKeySlot KeystoreError = "unsupported algorithm for key slot"
	ErrorExportNotAllowed               KeystoreError = "export not allowed"
	ErrorImportNotAllowed               KeystoreError = "import not allowed"
	ErrorKeyAlreadyExists               KeystoreError = "key already exists"
	ErrorFileAlreadyExists              KeystoreError = "file already exists"
	ErrorCertAlreadyExists              KeystoreError = "certificate already exists"
	ErrorUnableToOpenKeystore           KeystoreError = "unable to open keystore"
	ErrorKeystoreLocked                 KeystoreError = "keystore is locked"
	ErrorKeystoreHandleClosed           KeystoreError = "keystore handle is closed"
	ErrorResetNotAllowed                KeystoreError = "reset not allowed"
	ErrorKeyGenFailed                   KeystoreError = "key generation failed"
)

const (
	ErrorInvalidSignature           ValidationError = "invalid signature"
	ErrorInvalidMessageDigest       ValidationError = "invalid message digest"
	ErrorInvalidHeader              ValidationError = "invalid header"
	ErrorInvalidContentLength       ValidationError = "invalid content length"
	ErrorInvalidSigningKey          ValidationError = "invalid signing key"
	ErrorInvalidPrivateKey          ValidationError = "invalid private key"
	ErrorInvalidPublicKey           ValidationError = "invalid public key"
	ErrorInvalidCertificateTemplate ValidationError = "invalid certificate template"
	ErrorInvalidCSR                 ValidationError = "invalid certificate signing request"
	ErrorInvalidAttestationCert     ValidationError = "invalid attestation certificate"
	ErrorInvalidCertificate         ValidationError = "invalid certificate"
	ErrorInvalidManagementKey       ValidationError = "invalid management key"
	ErrorInvalidPIN                 ValidationError = "invalid pin"
	ErrorInvalidKeySlot             ValidationError = "invalid key slot"
	ErrorInvalidHandle              ValidationError = "invalid handle"
	ErrorInvalidCharacter           ValidationError = "invalid character: %s"
)

const (
	ErrorKeyExchangeFailed    CryptoError = "key exchange failed"
	ErrorUnsupportedAlgorithm CryptoError = "unsupported algorithm"
)

const (
	ErrorDuplicateExtension CertificateError = "duplicate extension (asn.1 OID: %s)"
	ErrorNoSuchExtension    CertificateError = "no such extension (asn.1 OID: %s)"
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
