package tpm

import (
	"fmt"
	"io"

	"github.com/google/go-tpm/tpm2"
	"github.com/google/go-tpm/tpmutil"
)

type Tpm2IdentityProvider struct {
	ownerAuth  []byte
	tpmHandle  io.ReadWriteCloser
	primaryKey tpmutil.Handle
	publicKey  tpm2.Public
}

func NewTpm2IdentityProvider(ownerAuth []byte) (*Tpm2IdentityProvider, error) {
	if ownerAuth == nil {
		return nil, fmt.Errorf("no OwnerAuth value provided")
	}
	return &Tpm2IdentityProvider{
		ownerAuth: ownerAuth,
	}, nil
}

func (idProvider *Tpm2IdentityProvider) Open() error {
	var err error = nil
	idProvider.tpmHandle, err = tpm2.OpenTPM()
	if err != nil {
		return err
	}
	return err
}

func (idProvider *Tpm2IdentityProvider) Close() error {
	var err error = nil
	idProvider.primaryKey = tpm2.HandleNull
	if idProvider.tpmHandle != nil {
		err = idProvider.tpmHandle.Close()
		idProvider.tpmHandle = nil
	}
	return err
}

func (idProvider *Tpm2IdentityProvider) GenerateKey() error {
	return nil
}

func (idProvider *Tpm2IdentityProvider) Certificate() ([]byte, error) {
	// var err error
	// idProvider.publicKey, _, _, err = tpm2.ReadPublic(idProvider.tpmHandle, idProvider.primaryKey)
	/*
		if err != nil {
			return nil, error
		}
		certificateTemplate := x509.Certificate{
			KeyUsage:    x509.KeyUsageCertSign,
			ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		}
		x509.CreateCertificate(rand.Reader)
	*/
	return nil, nil
}

func (idProvider *Tpm2IdentityProvider) Sign(data []byte) ([]byte, error) {
	return nil, nil
}
