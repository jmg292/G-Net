package memory

import (
	"github.com/jmg292/G-Net/internal/tracery/manifest/entries"
	"github.com/jmg292/G-Net/pkg/keyring"
)

type memoryManifest struct {
	devices map[string]entries.DeviceManifest
	users   map[string]keyring.PublicKeyRing
}
