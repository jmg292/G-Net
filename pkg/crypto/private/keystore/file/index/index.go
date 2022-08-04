package index

import (
	"fmt"

	"github.com/jmg292/G-Net/internal/utilities/convert"
	"github.com/jmg292/G-Net/pkg/gnet"
)

const size int = 8

type Index [size]byte

func New(signCertSize uint16, authCertSize uint16, encCertSize uint16, devCertSize uint16) (idx Index) {
	copy(idx[:2], convert.UInt16ToBytes(signCertSize))
	copy(idx[2:4], convert.UInt16ToBytes(authCertSize))
	copy(idx[4:6], convert.UInt16ToBytes(encCertSize))
	copy(idx[6:], convert.UInt16ToBytes(devCertSize))
	return idx
}

func (i *Index) LoadOffsets(indexBytes []byte) (err error) {
	if len(indexBytes) >= size {
		copy(i[:], indexBytes[:8])
	} else {
		err = fmt.Errorf(string(gnet.ErrorInvalidContentLength))
	}
	return
}
