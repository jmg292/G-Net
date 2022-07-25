package tracery

type FileSystemTracery struct {
	storagePath string
}

func NewFileSystemTracery(storagePath string) *FileSystemTracery {
	return &FileSystemTracery{
		storagePath: storagePath,
	}
}
