package adminslot_test

import (
	"bytes"
	"crypto/rand"
	"os"
	"testing"

	"github.com/jmg292/G-Net/pkg/crypto/private/keystore/file/adminslot"
	"golang.org/x/crypto/chacha20poly1305"
)

type slotIndex int

const (
	managementKeySlot slotIndex = iota
	kdfSaltSlot
)

const keySize int = 24

var testData [adminslot.Size]byte
var pin [8]byte
var salt [32]byte

func testSetup() {
	rand.Read(testData[:adminslot.Size-chacha20poly1305.Overhead])
	rand.Read(pin[:])
	rand.Read(salt[:])
}

func getAdminSlot() *adminslot.AdminSlot {
	slot := adminslot.Empty()
	slot.Load(testData[:])
	return slot
}

func getLockedAdminSlot() *adminslot.AdminSlot {
	slot := getAdminSlot()
	slot.Lock(pin[:], salt[:])
	return slot
}

func equalsTestDataSlot(idx slotIndex, keyContent []byte, t *testing.T) bool {
	slotOffset := keySize * int(idx)
	testDataSlotContent := testData[slotOffset : slotOffset+keySize]
	t.Logf("test content: 0x%x", testDataSlotContent)
	return bytes.Equal(testDataSlotContent, keyContent)
}

func keyContentMatches(slot *adminslot.AdminSlot, idx slotIndex, t *testing.T) bool {
	var keyContent []byte
	switch idx {
	case managementKeySlot:
		keyContent = slot.ManagementKey()
	case kdfSaltSlot:
		keyContent = slot.KdfSalt()
	}
	t.Logf("key content: 0x%x", keyContent)
	return equalsTestDataSlot(idx, keyContent, t)
}

func TestEmpty(t *testing.T) {
	slot := adminslot.Empty()
	if !slot.IsEmpty() {
		t.Errorf("slot contents: 0x%x", slot[:])
	}
}

func TestGetManagementKey(t *testing.T) {
	slot := getAdminSlot()
	if !equalsTestDataSlot(managementKeySlot, slot.ManagementKey(), t) {
		t.Fail()
	}
}

func TestGetKdfSalt(t *testing.T) {
	slot := getAdminSlot()
	if !equalsTestDataSlot(kdfSaltSlot, slot.KdfSalt(), t) {
		t.Fail()
	}
}

func TestLockSlot(t *testing.T) {
	slot := getLockedAdminSlot()
	for i := 0; i < 2; i++ {
		if keyContentMatches(slot, slotIndex(i), t) {
			t.Fail()
		}
	}
}

func TestUnlockSlot(t *testing.T) {
	slot := getLockedAdminSlot()
	slot.Unlock(pin[:], salt[:])
	for i := 0; i < 2; i++ {
		if !keyContentMatches(slot, slotIndex(i), t) {
			t.Fail()
		}
	}
}

func TestMain(m *testing.M) {
	testSetup()
	os.Exit(m.Run())
}
