package service

import (
	archiveCreate "doodocsTask/internal/service/archiveCreate"
	archiveInfo "doodocsTask/internal/service/archiveInfo"
)

type Service struct {
	archiveInfo.Info
	archiveCreate.Create
}

func NewService() *Service {
	return &Service{
		Info:   archiveInfo.NewArchiveInfoService(),
		Create: archiveCreate.NewArchiveCreateService(),
	}
}
