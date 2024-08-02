package utils

import (
	"net/smtp"
	"strconv"

	"github.com/edaywalid/chat-app/configs"
)

type EmailService struct {
	config *configs.Config
}

func NewEmailService(config *configs.Config) *EmailService {
	return &EmailService{config: config}
}

func (s *EmailService) SendEmail(subject, body, to string) error {
	auth := smtp.PlainAuth("", s.config.SMTPUser, s.config.SMTPPass, s.config.SMTPHost)

	mime := "MIME-version: 1.0;\r\n" +
		"Content-Type: text/plain; charset=\"UTF-8\";\r\n"

	msg := []byte("To: " + to + "\r\n" +
		"From: " + s.config.SMTPUser + "\r\n" +
		"Subject: " + subject + "\r\n" +
		mime + "\r\n" +
		body)

	err := smtp.SendMail(
		s.config.SMTPHost+":"+strconv.Itoa(s.config.SMTPPort),
		auth,
		s.config.SMTPUser,
		[]string{to},
		msg,
	)
	if err != nil {
		return err
	}
	return nil
}
