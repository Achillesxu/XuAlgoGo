// Package gmail
// Time    : 2022/8/21 13:17
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package gmail

import (
	"gopkg.in/mail.v2"
	"os"
	"testing"
)

func TestSimpleMail(t *testing.T) {
	m := mail.NewMessage()
	m.SetHeader("From", "achilles.xu@outlook.com")
	m.SetHeader("To", "18682193124@163.com")
	m.SetAddressHeader("Cc", "yuqingxushiyin@gmail.com", "xu shi yin")
	m.SetHeader("Subject", "Go language send mail test!!!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Oh my zsh</i>!")
	d := mail.NewDialer("smtp.office365.com", 587, "achilles.xu@outlook.com", os.Getenv("MAIL_PASS"))
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Send the email to some one
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
