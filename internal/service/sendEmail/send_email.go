package sendEmail

import (
	"crypto/tls"
	"net/smtp"
)

func (s *SendEmailService) SendFileToEmail(fileData []byte, toEmails []string) error {
	auth := smtp.PlainAuth("", s.config.SMTPUsername, s.config.SMTPPassword, s.config.SMTPServer)

	tlsConfig := &tls.Config{
		ServerName: s.config.SMTPServer,
	}

	for _, toEmail := range toEmails {
		// Create the email message
		from := s.config.SMTPUsername

		to := toEmail

		// Establish a TLS connection
		client, err := smtp.Dial(s.config.SMTPServer + ":" + s.config.SMTPPort)
		if err != nil {
			return err
		}

		err = client.StartTLS(tlsConfig)
		if err != nil {
			return err
		}

		// Authenticate and send the email
		err = client.Auth(auth)
		if err != nil {
			return err
		}

		err = client.Mail(from)
		if err != nil {
			return err
		}

		err = client.Rcpt(to)
		if err != nil {
			return err
		}

		w, err := client.Data()
		if err != nil {
			return err
		}

		_, err = w.Write(fileData)
		if err != nil {
			return err
		}

		err = w.Close()
		if err != nil {
			return err
		}

		err = client.Quit()
		if err != nil {
			return err
		}
	}

	return nil
}
