package sendEmail

import (
	"doodocsTask/config"
)

type Send interface {
	SendFileToEmail(fileData []byte, toEmails []string) error
}

type SendEmailService struct {
	config config.Config
}

func NewSendEmailService(config config.Config) *SendEmailService {
	return &SendEmailService{
		config: config,
	}
}
