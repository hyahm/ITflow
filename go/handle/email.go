package handle

import (
	"encoding/json"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/email"
	"itflow/internal/response"
	"itflow/mail"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func TestEmail(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	getemail := xmux.GetData(r).Data.(*email.Email)

	mail.TestMail(getemail.Host, getemail.EmailAddr, getemail.Password, getemail.Port, getemail.To)
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func SaveEmail(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	getemail := xmux.GetData(r).Data.(*email.Email)

	if getemail.Id < 1 {
		var err error
		errorcode.Id, err = db.Mconn.Insert("insert into email(email,password,port,host,enable) values(?,?,?,?,?)", getemail.EmailAddr,
			getemail.Password, getemail.Port, getemail.Host, getemail.Enable)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
	} else {
		_, err := db.Mconn.Update("update email set email=?,password=?,port=?,host=?,enable=? where id=?", getemail.EmailAddr,
			getemail.Password, getemail.Port, getemail.Host, getemail.Enable, getemail.Id)
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
	cache.CacheEmail.Id = errorcode.Id
	errorcode.Id = getemail.Id
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return
}

func GetEmail(w http.ResponseWriter, r *http.Request) {

	email := &cache.Email{}

	email = cache.CacheEmail
	send, _ := json.Marshal(email)
	w.Write(send)
	return

}
