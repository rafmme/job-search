package util

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/opensaucerer/goaxios"
	mail "github.com/xhit/go-simple-mail/v2"
)

func SendEmail(to, title, content, via string) {
	if strings.ToLower((via)) == "ext" {
		a := goaxios.GoAxios{
			Url:    os.Getenv("EXT_MAIL_SERVICE_URL"),
			Method: "POST",
			Body: map[string]string{
				"to":      to,
				"subject": title,
				"html":    content,
			},
		}

		_, _, d, err := a.RunRest()
		if err != nil {
			log.Printf("err: %v", err)
		}

		log.Println(d)
		return
	}

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
