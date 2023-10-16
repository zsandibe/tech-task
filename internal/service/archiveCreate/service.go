package archiveCreate

import "mime/multipart"

type Create interface {
	CreateArchive(files []*multipart.FileHeader, validMimeTypes map[string]bool) ([]byte, error)
}

type ArchiveCreateService struct {
}

func NewArchiveCreateService() *ArchiveCreateService {
	return &ArchiveCreateService{}
}
