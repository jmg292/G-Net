package index

import (
	"github.com/jmg292/G-Net/internal/utilities/convert"
	"github.com/jmg292/G-Net/pkg/crypto"
)

func (i *index) getCertificateSize(keySlot crypto.KeySlot) int {
	relativeOffset := int(keySlot) * 2
	return int(convert.BytesToUInt16(i[relativeOffset : relativeOffset+2]))
}

func (*index) SigningCertificateOffset() int {
	return int(certificateBase)
}

func (i *index) SigningCertificateSize() int {
	return i.getCertificateSize(crypto.SigningKeySlot)
}

func (i *index) AuthenticationCertificateOffset() int {
	return i.SigningCertificateOffset() + i.SigningCertificateSize()
}

func (i *index) AuthenticationCertificateSize() int {
	return i.getCertificateSize(crypto.AuthenticationKeySlot)
}

func (i *index) EncryptionCertificateOffset() int {
	return i.AuthenticationCertificateOffset() + i.AuthenticationCertificateSize()
}

func (i *index) EncryptionCertificateSize() int {
	return i.getCertificateSize(crypto.EncryptionKeySlot)
}

func (i *index) DeviceCertificateOffset() int {
	return i.EncryptionCertificateOffset() + i.EncryptionCertificateSize()
}

func (i *index) DeviceCertificateSize() int {
	return i.getCertificateSize(crypto.DeviceKeySlot)
}
