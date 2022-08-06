package index_test

import (
	"crypto/rand"
	"os"
	"testing"

	"github.com/jmg292/G-Net/internal/utilities/convert"
	"github.com/jmg292/G-Net/pkg/crypto/private/keystore/file/index"
)

const certificateBaseOffset = 512

var testData [index.Size]byte
var (
	testSignCertSize uint16 = 0
	testAuthCertSize uint16 = 0
	testEncCertSize  uint16 = 0
	testDevCertSize  uint16 = 0
)

func testSetup() {
	rand.Read(testData[:])
	testSignCertSize = convert.BytesToUInt16(testData[:2])
	testAuthCertSize = convert.BytesToUInt16(testData[2:4])
	testEncCertSize = convert.BytesToUInt16(testData[4:6])
	testDevCertSize = convert.BytesToUInt16(testData[6:])
}

func getIndex() *index.Index {
	return index.New(testSignCertSize, testAuthCertSize, testEncCertSize, testDevCertSize)
}

func TestMain(m *testing.M) {
	testSetup()
	os.Exit(m.Run())
}

func TestEmptyIndex(t *testing.T) {
	idx := index.Empty()
	if !idx.IsEmpty() {
		t.Errorf("Empty index is not empty")
	}
}

func TestLoadOffsets(t *testing.T) {
	idx := index.Empty()
	t.Logf("index.LoadOffsets(0x%x)", testData)
	idx.LoadOffsets(testData[:])
	if idx.IsEmpty() {
		t.Errorf("Failed to load offsets")
	}
}

func TestNewIndex(t *testing.T) {
	t.Logf("index.New(%d, %d, %d, %d)", testSignCertSize, testAuthCertSize, testEncCertSize, testDevCertSize)
	idx := getIndex()
	if idx.IsEmpty() {
		t.Errorf("Failed to create index with data")
	}
}

func testCertificate(expectedOffset int, expectedSize int, actualOffset int, actualSize int, t *testing.T) {
	t.Logf("expecting offset: %d, length %d", expectedOffset, expectedSize)
	if actualOffset != expectedOffset {
		t.Errorf("expected offset: %d, got %d", expectedOffset, actualOffset)
	}
	if actualSize != expectedSize {
		t.Errorf("expected size: %d, got %d", expectedSize, actualSize)
	}
}

func TestSigningCertificate(t *testing.T) {
	idx := getIndex()
	expectedOffset := certificateBaseOffset
	testCertificate(expectedOffset, int(testSignCertSize), idx.SigningCertificateOffset(), idx.SigningCertificateSize(), t)
}

func TestAuthenticationCertificate(t *testing.T) {
	idx := getIndex()
	expectedOffset := int(certificateBaseOffset + testSignCertSize)
	testCertificate(expectedOffset, int(testAuthCertSize), idx.AuthenticationCertificateOffset(), idx.AuthenticationCertificateSize(), t)
}

func TestEncryptionCertfiicate(t *testing.T) {
	idx := getIndex()
	expectedOffset := int(certificateBaseOffset + testSignCertSize + testAuthCertSize)
	testCertificate(expectedOffset, int(testEncCertSize), idx.EncryptionCertificateOffset(), idx.EncryptionCertificateSize(), t)
}

func TestDeviceCertificate(t *testing.T) {
	idx := getIndex()
	expectedOffset := int(certificateBaseOffset + testSignCertSize + testAuthCertSize + testEncCertSize)
	testCertificate(expectedOffset, int(testDevCertSize), idx.DeviceCertificateOffset(), idx.DeviceCertificateSize(), t)
}
