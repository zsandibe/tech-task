package sendEmail

import (
	"doodocsTask/config"
	"io"
)

type Send interface {
	SendFileToEmail(emails []string, file io.Reader, fileName string) error
}

type SendEmailService struct {
	config config.Config
}

func NewSendEmailService(config config.Config) *SendEmailService {
	return &SendEmailService{
		config: config,
	}
}
