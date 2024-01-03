package service

import (
	"Say-Hi/notification/constants"
	"Say-Hi/notification/contracts"
	"fmt"
	"gopkg.in/gomail.v2"
)

type SendEmailService struct {
}

func NewSendEmailService() *SendEmailService {
	return &SendEmailService{}
}

func (s *SendEmailService) SendEmail(data *contracts.SendEmailRequest) error {

	message := gomail.NewMessage()
	message.SetHeader(constants.From, constants.HostEmail)             // Sender email address
	message.SetHeader(constants.To, data.Email)                        // Recipient email address
	message.SetHeader(constants.Subject, constants.VerifyEmailSubject) // Email subject
	message.SetBody(constants.ContentTextPlain, fmt.Sprintf(constants.VerifyEmailMessage, data.Name, data.OTP))

	// Set up the email sender
	sender := gomail.NewDialer(
		constants.SMTPServerAddress,    // SMTP server address
		constants.SMTPServerPort,       // SMTP server port (587 for TLS, 465 for SSL)
		constants.HostEmail,            // Your email address
		constants.HostEmailAppPassword, // Your email password or app-specific password
	)

	if err := sender.DialAndSend(message); err != nil {
		return err
	}

	return nil
}
