package util

import (
	"fmt"
	"log"
	"os"

	mail "github.com/xhit/go-simple-mail/v2"
)

func SendEmail(to, title, content string) {
	senderAddr := os.Getenv("MAIL_HOST_USERNAME")
	mailServer := mail.NewSMTPClient()
	mailServer.Host = os.Getenv("MAIL_HOST")
	mailServer.Port = 587
	mailServer.Username = senderAddr
	mailServer.Password = os.Getenv("MAIL_HOST_PASSWORD")

	smtpClient, err := mailServer.Connect()
	if err != nil {
		log.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(fmt.Sprintf("Job Alert <%s>", senderAddr))
	email.AddTo(to)
	email.SetSubject(title)
	email.SetBody(mail.TextHTML, content)

	err = email.Send(smtpClient)
	if err != nil {
		log.Println(err)
	}
}
