package index

import (
	"bytes"

	"github.com/jmg292/G-Net/internal/utilities/convert"
	"github.com/jmg292/G-Net/pkg/gnet"
)

const Size int = 8

type Index [Size]byte

var empty Index

func Empty() *Index {
	var idx Index
	return &idx
}

func New(signCertSize uint16, authCertSize uint16, encCertSize uint16, devCertSize uint16) (idx *Index) {
	idx = Empty()
	copy(idx[:2], convert.UInt16ToBytes(signCertSize))
	copy(idx[2:4], convert.UInt16ToBytes(authCertSize))
	copy(idx[4:6], convert.UInt16ToBytes(encCertSize))
	copy(idx[6:], convert.UInt16ToBytes(devCertSize))
	return
}

func (i *Index) LoadOffsets(indexBytes []byte) (err error) {
	if len(indexBytes) >= Size {
		copy(i[:], indexBytes[:8])
	} else {
		err = gnet.ErrorInvalidContentLength
	}
	return
}

func (i *Index) IsEmpty() bool {
	return bytes.Equal(empty[:], i[:])
}
