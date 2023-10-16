package archiveInfo

import "doodocsTask/internal/models"

type Info interface {
	GetInfoByArchive(fileData []byte, fileName string) (*models.Archive, error)
}

type ArchiveInfoService struct {
}

func NewArchiveInfoService() *ArchiveInfoService {
	return &ArchiveInfoService{}
}
