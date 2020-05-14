package handle

import (
	"encoding/json"
	"itflow/app/bugconfig"
	"itflow/app/mail"
	"itflow/db"
	"itflow/network/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func TestEmail(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	getemail := xmux.GetData(r).Data.(*bugconfig.Email)

	mail.TestMail(getemail.EmailAddr, getemail.Password, getemail.Port, getemail.To)
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func SaveEmail(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	getemail := xmux.GetData(r).Data.(*bugconfig.Email)

	if getemail.Id < 0 {
		var err error
		bugconfig.CacheEmail.Id, err = db.Mconn.Insert("insert into email(email,password,port,createuser,createbug,passbug) values(?,?,?,?,?,?)", getemail.EmailAddr, getemail.Password, getemail.Port, getemail.CreateUser, getemail.CreateBug, getemail.PassBug)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
	} else {
		_, err := db.Mconn.Update("update email set email=?,password=?,port=?,createuser=?,createbug=?,passbug=? where id=?", getemail.EmailAddr, getemail.Password, getemail.Port, getemail.CreateUser, getemail.CreateBug, getemail.PassBug, getemail.Id)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
	}
	bugconfig.CacheEmail.PassBug = getemail.PassBug
	bugconfig.CacheEmail.CreateUser = getemail.CreateUser
	bugconfig.CacheEmail.CreateUser = getemail.CreateUser
	bugconfig.CacheEmail.EmailAddr = getemail.EmailAddr
	bugconfig.CacheEmail.Port = getemail.Port
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return
}

func GetEmail(w http.ResponseWriter, r *http.Request) {

	email := &bugconfig.Email{}

	email = bugconfig.CacheEmail
	email.Password = ""
	send, _ := json.Marshal(email)
	w.Write(send)
	return

}
