package gnet

type ApplicationError string

func (e ApplicationError) Error() string {
	return string(e)
}

type KeystoreError string

func (e KeystoreError) Error() string {
	return string(e)
}

type ItemNotFound string

func (e ItemNotFound) Error() string {
	return string(e)
}

type ValidationError string

func (e ValidationError) Error() string {
	return string(e)
}

type CryptoError string

func (e CryptoError) Error() string {
	return string(e)
}

type ConnectionError string

func (e ConnectionError) Error() string {
	return string(e)
}

type CertificateError string

func (e CertificateError) Error() string {
	return string(e)
}

type TraceryError string

func (e TraceryError) Error() string {
	return string(e)
}

type ManifestError string

func (e ManifestError) Error() string {
	return string(e)
}

type SequencingError string

func (e SequencingError) Error() string {
	return string(e)
}
