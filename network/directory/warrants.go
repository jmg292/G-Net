package directory

import (
	"fmt"

	"github.com/jmg292/G-Net/network/access"
	"github.com/jmg292/G-Net/utilities/gnet"
)

func (network *directory) getWarrantBlockIndex(warrantId []byte) (uint64, error) {
	blockIndex, err := network.tracery.Manifest().GetAdminWarrantBlockIndex(warrantId)
	if err != nil {
		blockIndex, err = network.tracery.Manifest().GetDeviceWarrantBlockIndex(warrantId)
	}
	return blockIndex, err
}

func (network *directory) GetIdentityFromWarrantId(warrantId []byte) (*access.Identity, error) {
	blockIndex, err := network.getWarrantBlockIndex(warrantId)
	if err != nil {
		return nil, err
	}
	warrantBlock, err := network.tracery.GetBlockByIndex(blockIndex)
	if err != nil {
		return nil, err
	}
	return nil, fmt.Errorf(string(gnet.ErrorNotYetImplemented))
}
