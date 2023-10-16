package delivery

import (
	createHandler "doodocsTask/internal/delivery/archiveCreate"
	infoHandler "doodocsTask/internal/delivery/archiveInfo"
	sendHandler "doodocsTask/internal/delivery/sendEmail"
	s "doodocsTask/internal/service"
)

type Handler struct {
	Create  *createHandler.ArchiveCreateHandler
	Info    *infoHandler.ArchiveInfoHandler
	Send    *sendHandler.SendEmailHandler
	service *s.Service
}

func NewHandler(service *s.Service) *Handler {
	return &Handler{
		Info:   infoHandler.NewArchiveInfoHandler(service.Info),
		Create: createHandler.NewArchiveCreateHandler(service.Create),
		Send:   sendHandler.NewSendEmailHandler(service.Send),
	}
}
