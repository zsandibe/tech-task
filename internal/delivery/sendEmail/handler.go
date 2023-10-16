package sendEmail

import "io"

type SendEmailProvider interface {
	SendFileToEmail(emails []string, file io.Reader, fileName string) error
}

type SendEmailHandler struct {
	send SendEmailProvider
}

func NewSendEmailHandler(send SendEmailProvider) *SendEmailHandler {
	return &SendEmailHandler{
		send: send,
	}
}
