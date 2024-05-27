package email

import "full/pkg/logger"

type EmailServiceImpl struct{}

func NewEmailService() IEmailService {
	return &EmailServiceImpl{}
}

func (receiver *EmailServiceImpl) SendEmail(to, subject, body string) error {
	logger.Infof("Email sent to to: %s, with subject: %s and body: %s!", to, subject, body)
	return nil
}
