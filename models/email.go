package models

import (
	"os"

	gomail "gopkg.in/gomail.v2"
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
func (h *Host) SendEmail(destination []string) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "aydinaalperen@gmail.com")
	msg.SetHeader("To", "aydinaalperen@gmail.com")
	msg.SetHeader("Subject", "test")
	msg.SetBody("text/html", "<b>Gomail test</b>")
	msg.Attach("")

	n := gomail.NewDialer("smtp.gmail.com", 587, "aydinaalperen@gmail.com", "alptseren61")

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}
}
