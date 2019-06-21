package gamail

import (
	"crypto/tls"
	"galog"
	"strings"
	"time"
)

func doc(subject string, content string, user []string) {
	mailconf := Newmailconfig()
	mailconf.Username = "test@126.com"
	mailconf.Password = "123456789"
	mailconf.Content = content
	mailconf.Subject = subject
	mailconf.Tolist = user
	mailconf.SendMail()
}

type mailConfig struct {
	Host        string //smtp.126.com
	Port        int    // 25
	Username    string
	Password    string
	Subject     string
	ContentType string // text/html    text
	Content     string
	Tolist      []string // sendto list
	AttachPath  string
}

func Newmailconfig() *mailConfig {
	return &mailConfig{
		ContentType: "text/html",
		Tolist:      make([]string, 0),
	}
}

func (mc *mailConfig) AddUser(mailaddr string) {
	if mc.Tolist == nil {
		mc.Tolist = make([]string, 0)
	}
	mc.Tolist = append(mc.Tolist, mailaddr)
}

func (mc *mailConfig) SetContent(content string) {
	mc.Content = content
}

func (mc *mailConfig) SendMail() {

	if mc.Username == "" {
		galog.Email("username is empty")
		return
	}
	if mc.Host == "" {
		hl := strings.Split(mc.Username, "@")
		if len(hl) < 2 {
			galog.Email("username %s is not vaild", mc.Username)
			return
		}
		mc.Host = "smtp." + hl[1]
	}

	if mc.Port == 0 {
		mc.Port = 25
	}
	if len(mc.Tolist) < 1 {
		galog.Email("sender is empty")
		return
	}
	d := NewDialer(mc.Host, mc.Port, mc.Username, mc.Password)
	if mc.Port != 25 {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	msg := NewMessage()

	msg.SetHeader("From", mc.Username)
	msg.SetHeader("To", mc.Tolist...)
	msg.SetBody(mc.ContentType, mc.Content)
	if mc.AttachPath != "" {
		msg.Attach(mc.AttachPath)
	}

	msg.SetHeader("Subject", mc.Subject)

	if err := d.DialAndSend(msg); err != nil {
		galog.Email("time:%s,error:%v", time.Now().Unix(), err)
		return
	}
	tousers := ""
	for _, v := range mc.Tolist {
		if tousers == "" {
			tousers = v
		} else {
			tousers = tousers + "," + v
		}
	}
	galog.Email("time:%s,to: %v,send mail successed", time.Now().Unix(), tousers)

}
