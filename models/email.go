package models

import (
	"fmt"
	"net/smtp"
	"os"
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
	smtpHost := "smtp.gmail.com"

	smtpPort := "587"

	auth := smtp.PlainAuth("", h.from, h.password, smtpHost)

	message := []byte("This is a test message!")
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, h.from, destination, message)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
