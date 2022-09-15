package models

import (
	"crypto/tls"
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

type Host struct {
	from     string
	password string
}

func (h *Host) ConnectToMailHost() {
	FROM := os.Getenv("MAIL_HOST")
	PASSWORD := os.Getenv("MAIL_PASSWORD")

	h.from = FROM
	h.password = PASSWORD
}
func SendEmail(destination string, verifCode string) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "baebeez.secure@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", destination)

	// Set E-Mail subject
	m.SetHeader("Subject", "Baebeez Activation Code")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "Your activation code: "+verifCode)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "baebeez.secure@gmail.com", "baebeez-914")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}
