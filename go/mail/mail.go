package mail

import (
	"crypto/tls"
	"strings"

	"gopkg.in/gomail.v2"
)

func TestMail(host, username string, password string, port int, touser, nickname string) error {
	d := gomail.NewDialer(host, port, username, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if nickname == "" {
		nickname = strings.Split(username, "@")[0]
	}
	m := gomail.NewMessage()
	m.SetHeader("From", username)
	m.SetHeader("To", strings.Split(touser, ",")...)
	m.SetAddressHeader("Cc", username, nickname)
	m.SetHeader("Subject", "验证您的邮箱是否能收到邮件")
	m.SetBody("text/html", "恭喜， 您的邮箱可以使用")
	// m.Attach("/home/Alex/lolcat.jpg")
	return d.DialAndSend(m)
}
