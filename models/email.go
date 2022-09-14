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
func SendEmail(destination string, verifCode string) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "baebeez.support@protonmail.com")
	msg.SetHeader("To", destination)
	msg.SetHeader("Subject", "Baebeez Verification Code")
	msg.SetBody("text/html", "<b>Your Verification Code : </b>"+verifCode)
	msg.Attach("")

	n := gomail.NewDialer("smtp.gmail.com", 587, "aydinaalperen@gmail.com", "baebeez-914")

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}
}
