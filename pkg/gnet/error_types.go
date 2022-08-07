package gnet

type ApplicationError string

func (e ApplicationError) Error() string {
	return string(e)
}

type KeystoreError ApplicationError

func (e KeystoreError) Error() string {
	return string(e)
}

type ItemNotFound ApplicationError

func (e ItemNotFound) Error() string {
	return string(e)
}

type ValidationError ApplicationError

func (e ValidationError) Error() string {
	return string(e)
}

type CryptoError ValidationError

func (e CryptoError) Error() string {
	return string(e)
}

type TraceryError ApplicationError

func (e TraceryError) Error() string {
	return string(e)
}

type ManifestError TraceryError

func (e ManifestError) Error() string {
	return string(e)
}

type SequencingError TraceryError

func (e SequencingError) Error() string {
	return string(e)
}
