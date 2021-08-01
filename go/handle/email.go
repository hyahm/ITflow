package handle

import (
	"encoding/json"
	"fmt"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/email"
	"itflow/mail"
	"itflow/response"
	"net/http"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func TestEmail(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	getemail := xmux.GetInstance(r).Data.(*email.Email)
	if getemail.Nickname == "" {
		getemail.Nickname = strings.Split(getemail.EmailAddr, "@")[0]
	}
	golog.Infof("%#v", *getemail)
	err := mail.TestMail(getemail.Host, getemail.EmailAddr, getemail.Password, getemail.Port, getemail.To, getemail.Nickname)
	if err != nil {
		w.Write(errorcode.ErrorE(err))
		return
	}
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func SaveEmail(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	getemail := xmux.GetInstance(r).Data.(*email.Email)

	if getemail.Id < 1 {
		var err error
		errorcode.ID, err = db.Mconn.Insert("insert into email(email,password,port,host,enable,nickname) values(?,?,?,?,?,?)", getemail.EmailAddr,
			getemail.Password, getemail.Port, getemail.Host, getemail.Enable, getemail.Nickname)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
	} else {
		_, err := db.Mconn.Update("update email set email=?,password=?,port=?,host=?,enable=?,nickname=? where id=?", getemail.EmailAddr,
			getemail.Password, getemail.Port, getemail.Host, getemail.Enable, getemail.Id, getemail.Nickname)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
	}
	cache.CacheEmail.Enable = getemail.Enable
	cache.CacheEmail.Host = getemail.Host
	cache.CacheEmail.Password = getemail.Password
	cache.CacheEmail.EmailAddr = getemail.EmailAddr
	cache.CacheEmail.Port = getemail.Port
	cache.CacheEmail.Nickname = getemail.Nickname
	cache.CacheEmail.Id = errorcode.ID
	errorcode.ID = getemail.Id
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return
}

func GetEmail(w http.ResponseWriter, r *http.Request) {
	// errorcode := &response.Response{}

	// id := xmux.GetInstance(r).Get("uid")
	// email := &cache.Email{}
	// var email, password, host string
	// // email = cache.CacheEmail
	// var eid int64
	// var port int
	// var enable bool
	// err := db.Mconn.GetOne("select id, email,password,port, enable,host from user where id=?", id).Scan(
	// 	&eid, &email, &password, &port, &enable, &host)
	// if err != nil {
	// 	w.Write(errorcode.ErrorE(err))
	// 	return
	// }
	send := fmt.Sprintf(`{"code": 0, "email": "%s", "id": %d, "password": "%s", "port": %d, "enable": %t, "host":"%s", "nickname": "%s"}`,
		cache.CacheEmail.EmailAddr, cache.CacheEmail.Id, cache.CacheEmail.Password,
		cache.CacheEmail.Port, cache.CacheEmail.Enable, cache.CacheEmail.Host, cache.CacheEmail.Nickname)
	w.Write([]byte(send))
	return

}

func MyEmail(w http.ResponseWriter, r *http.Request) {
	// errorcode := &response.Response{}

	id := xmux.GetInstance(r).Get("uid")
	// email := &cache.Email{}
	// var email, password, host string
	// // email = cache.CacheEmail
	// var eid int64
	// var port int
	// var enable bool
	// err := db.Mconn.GetOne("select id, email,password,port, enable,host from user where id=?", id).Scan(
	// 	&eid, &email, &password, &port, &enable, &host)
	// if err != nil {
	// 	w.Write(errorcode.ErrorE(err))
	// 	return
	// }
	var email string
	err := db.Mconn.GetOne("select email from user where id=?", id).Scan(&email)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"code": 2, "msg": "%s"}`, err.Error())))
		return
	}
	w.Write([]byte(fmt.Sprintf(`{"code": 0, "email": "%s"}`, email)))
	return

}
