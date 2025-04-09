package mail

import (
	"esim/config"
	"net/smtp"
)

const smtpHost = "smtp.gmail.com"
const smtpPort = "587"

type Mailer interface {
	SendVerificationCode(toEmail string, code int) error
}

type mailer struct {
	addr     string
	password string
}

func (m mailer) SendVerificationCode(toEmail string, code int) error {
	msg := []byte(verificationCodeMsg(code))

	auth := smtp.PlainAuth("", m.addr, m.password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, m.addr, []string{toEmail}, msg)
	if err != nil {
		return err
	}

	return nil
}

func New(cfg config.Config) Mailer {
	return mailer{
		cfg.Mail.Address(),
		cfg.Mail.Password(),
	}
}
