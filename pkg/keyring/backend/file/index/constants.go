package index

type offset uint16

const (
	preamble        offset = 0
	versionNumber   offset = 7
	indexBytes      offset = 8
	salt            offset = 16
	adminSlot       offset = 48
	keySlotBase     offset = 112
	certificateBase offset = 624
)
