package model

import (
	"database/sql"
	"itflow/cache"
	"itflow/db"

	"github.com/hyahm/golog"
	"gopkg.in/gomail.v2"
)

var Default *DefaultValue
var CacheEmail *Email

func InitCache() {
	Default = &DefaultValue{}
	cache.CacheRoleID = make(map[int64]cache.PageInfo)
	rolerows, err := db.Mconn.GetRows("select id, name,info from roles")
	if err != nil {
		panic(err)
	}
	for rolerows.Next() {
		var id int64
		var name, info string
		rolerows.Scan(&id, &name, &info)
		cache.CacheRoleID[id] = cache.PageInfo{
			Name: name,
			Info: info,
		}
		// CacheRoleRid[name] = id
	}
	rolerows.Close()
	// 	//默认值
	result := db.Mconn.Select(&Default, "select * from defaultvalue")
	if result.Err != nil {
		if result.Err != sql.ErrNoRows {
			panic(err)
		}
	}
	cacheemail()
}

func cacheemail() {
	CacheEmail = &Email{}
	err := db.Mconn.GetOne("select id,email,password,port,host,enable,nickname from email limit 1").
		Scan(&CacheEmail.Id, &CacheEmail.Email,
			&CacheEmail.Password,
			&CacheEmail.Port, &CacheEmail.Host,
			&CacheEmail.Enable, &CacheEmail.NickName)

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
	d := gomail.NewDialer(e.Host, e.Port, e.Email, e.Password)

	m := gomail.NewMessage()
	m.SetHeader("From", e.Email)
	m.SetHeader("To", touser...)
	m.SetAddressHeader("From", e.Email, e.NickName)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)
	// mailconf.SendMail()
	if err := d.DialAndSend(m); err != nil {
		golog.Error(err)
	}
	return
}
