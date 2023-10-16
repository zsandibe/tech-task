package archiveCreate

import "mime/multipart"

type ArchiveCreateProvider interface {
	CreateArchive(files []*multipart.FileHeader, validMimeTypes map[string]bool) ([]byte, error)
}

type ArchiveCreateHandler struct {
	create ArchiveCreateProvider
}

func NewArchiveCreateHandler(create ArchiveCreateProvider) *ArchiveCreateHandler {
	return &ArchiveCreateHandler{
		create: create,
	}
}
