package cache

import (
	"database/sql"
	"itflow/db"

	"github.com/hyahm/golog"
	"gopkg.in/gomail.v2"
)

type Email struct {
	Host      string `json:"host"`
	Enable    bool   `json:"enable"`
	Id        int64  `json:"id"`
	Port      int    `json:"port"`
	EmailAddr string `json:"emailaddr"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	Code      int    `json:"code"`
}

func cacheemail() {
	CacheEmail = &Email{}
	err := db.Mconn.GetOne("select id,email,password,port,host,enable,nickname from email limit 1").
		Scan(&CacheEmail.Id, &CacheEmail.EmailAddr, &CacheEmail.Password, &CacheEmail.Port, &CacheEmail.Host, &CacheEmail.Enable, &CacheEmail.Nickname)

	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		panic(err)
	}
}

func (e *Email) SendMail(subject string, content string, touser ...string) {
	golog.Infof("%#v", *e)
	if !e.Enable {
		return
	}
	d := gomail.NewDialer(e.Host, e.Port, e.EmailAddr, e.Password)

	m := gomail.NewMessage()
	m.SetHeader("From", e.EmailAddr)
	m.SetHeader("To", touser...)
	m.SetAddressHeader("From", e.EmailAddr, e.Nickname)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)
	// mailconf.SendMail()
	if err := d.DialAndSend(m); err != nil {
		golog.Error(err)
	}
	return
}
