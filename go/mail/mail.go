package mail

import (
	"crypto/tls"
	"itflow/cache"
	"strings"

	"gopkg.in/gomail.v2"
)

func SendMail(subject string, content string, touser ...string) error {
	d := gomail.NewDialer("smtp.example.com", cache.CacheEmail.Port, cache.CacheEmail.EmailAddr, cache.CacheEmail.Password)

	m := gomail.NewMessage()
	m.SetHeader("From", cache.CacheEmail.EmailAddr)
	m.SetHeader("To", touser...)
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)
	// mailconf.SendMail()
	return d.DialAndSend(m)
}

func TestMail(username string, password string, port int, touser string) error {
	host := "smtp." + strings.Split(username, "@")[1]
	d := gomail.NewDialer(host, port, username, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", username)
	m.SetHeader("To", strings.Split(touser, ",")...)
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "验证您的邮箱是否能收到邮件")
	m.SetBody("text/html", "恭喜， 您的邮箱可以使用")
	// m.Attach("/home/Alex/lolcat.jpg")

	return d.DialAndSend(m)
	// mailconf.Username = username
	// mailconf.Password = password
	// mailconf.Port = port
	// mailconf.Tolist = []string{touser}
	// mailconf.Subject = "itflow发来的一声问候"
	// mailconf.Content = "恭喜，邮箱验证通过"
	// mailconf.SendMail()
}
