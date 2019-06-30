package handle

import (
	"bug/bugconfig"
	"bug/mail"
	"encoding/json"
	"github.com/hyahm/golog"
	"io/ioutil"
	"net/http"
)

func TestEmail(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		conn, _, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()

		getemail := &bugconfig.Email{}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorGetData())
			return
		}
		err = json.Unmarshal(b, getemail)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		mail.TestMail(getemail.EmailAddr, getemail.Password, getemail.Port, getemail.To)
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func SaveEmail(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		conn, _, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()

		getemail := &bugconfig.Email{}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorGetData())
			return
		}
		err = json.Unmarshal(b, getemail)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}

		var id int64
		if getemail.Id < 0 {
			id, err = conn.InsertWithID("insert into email(email,password,port,createuser,createbug,passbug) values(?,?,?,?,?,?)", getemail.EmailAddr, getemail.Password, getemail.Port, getemail.CreateUser, getemail.CreateBug, getemail.PassBug)
			if err != nil {
				golog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		} else {
			_, err = conn.Update("update email set email=?,password=?,port=?,createuser=?,createbug=?,passbug=? where id=?", getemail.EmailAddr, getemail.Password, getemail.Port, getemail.CreateUser, getemail.CreateBug, getemail.PassBug, getemail.Id)
			if err != nil {
				golog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}
		bugconfig.CacheEmail.PassBug = getemail.PassBug
		bugconfig.CacheEmail.CreateUser = getemail.CreateUser
		bugconfig.CacheEmail.CreateUser = getemail.CreateUser
		bugconfig.CacheEmail.EmailAddr = getemail.EmailAddr
		bugconfig.CacheEmail.Id = id
		bugconfig.CacheEmail.Port = getemail.Port
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func GetEmail(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		conn, _, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		email := &bugconfig.Email{}

		email = bugconfig.CacheEmail
		email.Password = ""
		send, _ := json.Marshal(email)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
