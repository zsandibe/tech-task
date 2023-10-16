package archiveInfo

import "doodocsTask/internal/models"

type ArchiveInfoProvider interface {
	GetInfoByArchive(fileData []byte, fileName string) (*models.Archive, error)
}

type ArchiveInfoHandler struct {
	info ArchiveInfoProvider
}

func NewArchiveInfoHandler(info ArchiveInfoProvider) *ArchiveInfoHandler {
	return &ArchiveInfoHandler{
		info: info,
	}
}
