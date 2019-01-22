package controllers

import (
	"crypto/tls"

	gomail "gopkg.in/gomail.v2"
)

func SendMail(emailReceiver string, message string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "cms.system.hblab@gmail.com")
	m.SetHeader("To", emailReceiver)
	m.SetHeader("Subject", "Confirm Register!")
	m.SetBody("text/html", message)
	d := gomail.NewDialer("smtp.gmail.com", 587, "cms.system.hblab", "a123456789b")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
