package tracery

import "gnet/tracery/wumbo"

type SqliteTracery struct {
	filePath     string
	currentState []byte
	currentBlock wumbo.Block
	rootBlock    wumbo.Block
}

func NewSqliteTracery(databasePath string) *SqliteTracery {
	return &SqliteTracery{
		filePath: databasePath,
	}
}

func (tracery *SqliteTracery) GetCurrentState() []byte {
	return tracery.currentState
}

func (tracery *SqliteTracery) GetCurrentBlock() wumbo.Block {
	return tracery.currentBlock
}

func (tracery *SqliteTracery) GetRootBlock() wumbo.Block {
	return tracery.rootBlock
}
