package index

import (
	"bytes"
	"fmt"

	"github.com/jmg292/G-Net/internal/utilities/convert"
	"github.com/jmg292/G-Net/pkg/gnet"
)

const Size int = 8

type index [Size]byte

func Empty() *index {
	var idx index
	return &idx
}

func New(signCertSize uint16, authCertSize uint16, encCertSize uint16, devCertSize uint16) (idx *index) {
	idx = Empty()
	copy(idx[:2], convert.UInt16ToBytes(signCertSize))
	copy(idx[2:4], convert.UInt16ToBytes(authCertSize))
	copy(idx[4:6], convert.UInt16ToBytes(encCertSize))
	copy(idx[6:], convert.UInt16ToBytes(devCertSize))
	return
}

func (i *index) LoadOffsets(indexBytes []byte) (err error) {
	if len(indexBytes) >= Size {
		copy(i[:], indexBytes[:8])
	} else {
		err = fmt.Errorf(string(gnet.ErrorInvalidContentLength))
	}
	return
}

func (i *index) IsEmpty() bool {
	return bytes.Equal(Empty()[:], i[:])
}
