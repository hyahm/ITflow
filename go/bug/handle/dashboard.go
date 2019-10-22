package handle

import (
	"encoding/json"
	"github.com/hyahm/golog"
	"itflow/bug/bugconfig"
	"net/http"
)

//type totalProject struct {
//	CountBugs     int `json:"countbugs"`
//	CountComplete int `json:"countcomplete"`
//}

type userCount struct {
	Code        int `json:"statuscode"`
	CountUsers  int `json:"countusers"`
	CountGroups int `json:"countgroups"`
}

func UserCount(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		 _, err := logtokenmysql(r)
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

		uc := &userCount{}

		getusersql := "select count(id) from user"
		getgroupsql := "select count(id) from rolegroup"
		err = bugconfig.Bug_Mysql.GetOne(getusersql).Scan(&uc.CountUsers)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		err = bugconfig.Bug_Mysql.GetOne(getgroupsql).Scan(&uc.CountGroups)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		send, _ := json.Marshal(uc)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

type projectCount struct {
	CountBugs     int `json:"countbugs"`
	CountComplete int `json:"countcomplete"`
	Code          int `json:"statuscode"`
}

func ProjectCount(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		 _, err := logtokenmysql(r)
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

		pc := &projectCount{}

		getbugs := "select count(id) from bugs"
		err = bugconfig.Bug_Mysql.GetOne(getbugs).Scan(&pc.CountBugs)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		send, _ := json.Marshal(pc)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
