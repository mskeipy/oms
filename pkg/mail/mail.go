package mail

import (
	"dropx/pkg/config"
	"github.com/go-gomail/gomail"
	"log"
)

func Config() *gomail.Dialer {
	d := gomail.NewDialer(
		config.Global.MailHost,
		config.Global.MailPort,
		config.Global.MailUsername,
		config.Global.MailPassword,
	)
	return d
}

func Send(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.Global.MailFromAddress)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	if err := Config().DialAndSend(m); err != nil {
		log.Println("SendMail error:", err)
		return err
	}

	log.Println("Email sent to", to)
	return nil
}
