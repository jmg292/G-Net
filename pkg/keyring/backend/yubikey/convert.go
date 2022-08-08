package yubikey

import (
	"github.com/go-piv/piv-go/piv"
	"github.com/jmg292/G-Net/pkg/gnet"
	"github.com/jmg292/G-Net/pkg/keyring"
)

func convertToPivSlot(s keyring.KeySlot) (slot piv.Slot, err error) {
	switch s {
	case keyring.SigningKeySlot:
		slot = piv.SlotSignature
	case keyring.AuthenticationKeySlot:
		slot = piv.SlotAuthentication
	case keyring.DeviceKeySlot:
		slot = piv.SlotCardAuthentication
	default:
		err = gnet.ErrorInvalidKeySlot
	}
	return
}

func convertToPivAlg(s keyring.SupportedKeyType) (alg piv.Algorithm, err error) {
	switch s {
	case keyring.EC256Key:
		alg = piv.AlgorithmEC256
	case keyring.EC384Key:
		alg = piv.AlgorithmEC384
	default:
		err = gnet.ErrorUnsupportedAlgorithm
	}
	return
}