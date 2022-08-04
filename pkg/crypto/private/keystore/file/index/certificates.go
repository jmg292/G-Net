package index

import (
	"github.com/jmg292/G-Net/internal/utilities/convert"
	"github.com/jmg292/G-Net/pkg/crypto"
)

func (i *Index) getCertificateSize(keySlot crypto.KeySlot) int {
	relativeOffset := int(keySlot) * 2
	return int(convert.BytesToUInt16(i[relativeOffset : relativeOffset+2]))
}

func (*Index) SigningCertificateOffset() int {
	return int(certificateBase)
}

func (i *Index) SigningCertificateSize() int {
	return i.getCertificateSize(crypto.SigningKeySlot)
}

func (i *Index) AuthenticationCertificateOffset() int {
	return i.SigningCertificateOffset() + i.SigningCertificateSize()
}

func (i *Index) AuthenticationCertificateSize() int {
	return i.getCertificateSize(crypto.AuthenticationKeySlot)
}

func (i *Index) EncryptionCertificateOffset() int {
	return i.AuthenticationCertificateOffset() + i.AuthenticationCertificateSize()
}

func (i *Index) EncryptionCertificateSize() int {
	return i.getCertificateSize(crypto.EncryptionKeySlot)
}

func (i *Index) DeviceCertificateOffset() int {
	return i.EncryptionCertificateOffset() + i.EncryptionCertificateSize()
}

func (i *Index) DeviceCertificateSize() int {
	return i.getCertificateSize(crypto.DeviceKeySlot)
}
