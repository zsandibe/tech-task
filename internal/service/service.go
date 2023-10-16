package service

import (
	"doodocsTask/config"
	archiveCreate "doodocsTask/internal/service/archiveCreate"
	archiveInfo "doodocsTask/internal/service/archiveInfo"
	sendEmail "doodocsTask/internal/service/sendEmail"
)

type Service struct {
	archiveInfo.Info
	archiveCreate.Create
	sendEmail.Send
	Config config.Config
}

func NewService(cfg config.Config) *Service {
	return &Service{
		Info:   archiveInfo.NewArchiveInfoService(),
		Create: archiveCreate.NewArchiveCreateService(),
		Send:   sendEmail.NewSendEmailService(cfg),
		Config: cfg,
	}
}
