package handle

import (
	"itflow/db"
	"itflow/internal/email"
	"itflow/mail"
	"itflow/model"
	"itflow/response"
	"net/http"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func TestEmail(w http.ResponseWriter, r *http.Request) {

	getemail := xmux.GetInstance(r).Data.(*email.Email)
	if getemail.Nickname == "" {
		getemail.Nickname = strings.Split(getemail.EmailAddr, "@")[0]
	}
	golog.Infof("%#v", *getemail)
	err := mail.TestMail(getemail.Host, getemail.EmailAddr, getemail.Password, getemail.Port, getemail.To, getemail.Nickname)
	if err != nil {
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		return
	}

}

func SaveEmail(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	getemail := xmux.GetInstance(r).Data.(*email.Email)

	if getemail.Id < 1 {
		result := db.Mconn.Insert("insert into email(email,password,port,host,enable,nickname) values(?,?,?,?,?,?)", getemail.EmailAddr,
			getemail.Password, getemail.Port, getemail.Host, getemail.Enable, getemail.Nickname)
		if result.Err != nil {
			golog.Error(result.Err)
			xmux.GetInstance(r).Response.(*response.Response).Msg = result.Err.Error()
			xmux.GetInstance(r).Response.(*response.Response).Code = 1
			return
		}
		errorcode.ID = result.LastInsertId
	} else {
		result := db.Mconn.Update("update email set email=?,password=?,port=?,host=?,enable=?,nickname=? where id=?", getemail.EmailAddr,
			getemail.Password, getemail.Port, getemail.Host, getemail.Enable, getemail.Id, getemail.Nickname)
		if result.Err != nil {
			golog.Error(result.Err)
			xmux.GetInstance(r).Response.(*response.Response).Msg = result.Err.Error()
			xmux.GetInstance(r).Response.(*response.Response).Code = 1
			return
		}
	}
	model.CacheEmail.Enable = getemail.Enable
	model.CacheEmail.Host = getemail.Host
	model.CacheEmail.Password = getemail.Password
	model.CacheEmail.Email = getemail.EmailAddr
	model.CacheEmail.Port = getemail.Port
	model.CacheEmail.NickName = getemail.Nickname
	model.CacheEmail.Id = errorcode.ID
	xmux.GetInstance(r).Response.(*response.Response).ID = getemail.Id
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
	xmux.GetInstance(r).Response.(*response.Response).Data = model.CacheEmail
}
