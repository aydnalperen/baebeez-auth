package models

import (
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

}
