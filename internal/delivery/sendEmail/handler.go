package sendEmail

type SendEmailProvider interface {
	SendFileToEmail(fileData []byte, toEmails []string) error
}

type SendEmailHandler struct {
	send SendEmailProvider
}

func NewSendEmailHandler(send SendEmailProvider) *SendEmailHandler {
	return &SendEmailHandler{
		send: send,
	}
}
