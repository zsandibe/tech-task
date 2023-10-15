package delivery

import (
	createHandler "doodocsTask/internal/delivery/archiveCreate"
	infoHandler "doodocsTask/internal/delivery/archiveInfo"
	sendHandler "doodocsTask/internal/delivery/sendEmail"
)

type Handler struct {
	Create *createHandler.ArchiveCreateHandler
	Info   *infoHandler.ArchiveInfoHandler
	Send   *sendHandler.SendEmailHandler
}

func NewHandler() *Handler {

}
