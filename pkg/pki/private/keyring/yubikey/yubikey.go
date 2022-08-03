package yubikey

import (
	"crypto"
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"strings"

	"github.com/go-piv/piv-go/piv"
)

var KeyTemplate = piv.Key{
	Algorithm:   piv.AlgorithmEC384,
	PINPolicy:   piv.PINPolicyAlways,
	TouchPolicy: piv.TouchPolicyAlways,
}

type YubikeyIdentityProvider struct {
	pin                       string
	yubikeyHandle             *piv.YubiKey
	AuthenticationCertificate *x509.Certificate
	authenticationKey         crypto.PrivateKey
	SigningCertificate        *x509.Certificate
	signingKey                crypto.PrivateKey
}

func NewYubikeyIdentityProvider(pin string) *YubikeyIdentityProvider {
	if pin == "" {
		pin = piv.DefaultPIN
	}
	return &YubikeyIdentityProvider{
		pin: pin,
	}
}

func (idProvider *YubikeyIdentityProvider) openStoredKeysAndCertificates() error {
	var err error = nil
	if idProvider.AuthenticationCertificate, err = idProvider.yubikeyHandle.Certificate(piv.SlotAuthentication); err != nil {
		return err
	}
	if idProvider.authenticationKey, err = idProvider.yubikeyHandle.PrivateKey(
		piv.SlotAuthentication,
		idProvider.AuthenticationCertificate.PublicKey,
		piv.KeyAuth{PIN: idProvider.pin},
	); err != nil {
		return err
	}
	if idProvider.SigningCertificate, err = idProvider.yubikeyHandle.Certificate(piv.SlotSignature); err != nil {
		return err
	}
	idProvider.signingKey, err = idProvider.yubikeyHandle.PrivateKey(
		piv.SlotSignature,
		idProvider.SigningCertificate.PublicKey,
		piv.KeyAuth{PIN: idProvider.pin},
	)
	return err
}

func (idProvider *YubikeyIdentityProvider) Open() error {
	smartCards, err := piv.Cards()
	if err != nil {
		return err
	}
	for _, cardName := range smartCards {
		if strings.Contains(strings.ToLower(cardName), "yubikey") {
			if idProvider.yubikeyHandle, err = piv.Open(cardName); err != nil {
				return err
			}
			break
		}
	}
	if idProvider.yubikeyHandle == nil {
		return fmt.Errorf("unable to open Yubikey")
	}
	return idProvider.openStoredKeysAndCertificates()
}

func (idProvider *YubikeyIdentityProvider) Close() error {
	var err error
	idProvider.pin = ""
	idProvider.SigningCertificate, idProvider.signingKey = nil, nil
	idProvider.AuthenticationCertificate, idProvider.authenticationKey = nil, nil
	if idProvider.yubikeyHandle != nil {
		err = idProvider.yubikeyHandle.Close()
	}
	return err
}

func (idProvider *YubikeyIdentityProvider) GenerateKey() error {
	if idProvider.AuthenticationCertificate != nil || idProvider.authenticationKey != nil || idProvider.SigningCertificate != nil || idProvider.signingKey != nil {
		return fmt.Errorf("key already exists, unable to generate a new one")
	}
	yubikeyMetadata, err := idProvider.yubikeyHandle.Metadata(idProvider.pin)
	if err != nil {
		return err
	}
	if _, err = idProvider.yubikeyHandle.GenerateKey(*yubikeyMetadata.ManagementKey, piv.SlotAuthentication, KeyTemplate); err != nil {
		return err
	}
	if _, err = idProvider.yubikeyHandle.GenerateKey(*yubikeyMetadata.ManagementKey, piv.SlotSignature, KeyTemplate); err != nil {
		return err
	}
	return idProvider.openStoredKeysAndCertificates()
}

func (idProvider *YubikeyIdentityProvider) Certificate() ([]byte, error) {
	if idProvider.AuthenticationCertificate == nil {
		return nil, fmt.Errorf("nil certificate found. Call Open() or GenerateKey() before attempting to retrieve the certificate")
	}
	return idProvider.AuthenticationCertificate.Raw, nil
}

func (idProvider *YubikeyIdentityProvider) Sign(data []byte) ([]byte, error) {
	if idProvider.authenticationKey == nil {
		return nil, fmt.Errorf("nil authentication key found.  Call Open() or GenerateKey() before attempting to use the authentication key")
	}
	signer, okay := idProvider.authenticationKey.(crypto.Signer)
	if !okay {
		return nil, fmt.Errorf("unable to use stored private key for signature operations")
	}
	return signer.Sign(rand.Reader, data, crypto.SHA3_256)
}
