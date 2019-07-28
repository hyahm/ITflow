package mail

import (
	"itflow/bug/bugconfig"
	"itflow/gamail"
)

func SendMail(subject string, content string, touser []string) {
	mailconf := gamail.Newmailconfig()
	mailconf.Username = bugconfig.CacheEmail.EmailAddr
	mailconf.Password = bugconfig.CacheEmail.Password
	mailconf.Port = bugconfig.CacheEmail.Port
	mailconf.Tolist = touser
	mailconf.Subject = subject
	mailconf.Content = content
	mailconf.SendMail()
}

func TestMail(username string, password string, port int, touser string) {
	mailconf := gamail.Newmailconfig()
	mailconf.Username = username
	mailconf.Password = password
	mailconf.Port = port
	mailconf.Tolist = []string{touser}
	mailconf.Subject = "itflow发来的一声问候"
	mailconf.Content = "恭喜，邮箱验证通过"
	mailconf.SendMail()
}
