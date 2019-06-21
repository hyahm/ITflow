package handle

import (
	"encoding/json"
	"galog"
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
		conn, _, err := logtokenmysql(r)
		errorcode := &errorstruct{}

		if err != nil {
			galog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		uc := &userCount{}

		getusersql := "select count(id) from user"
		getgroupsql := "select count(id) from rolegroup"
		err = conn.GetOne(getusersql).Scan(&uc.CountUsers)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		err = conn.GetOne(getgroupsql).Scan(&uc.CountGroups)
		if err != nil {
			galog.Error(err.Error())
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
		conn, _, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			galog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		pc := &projectCount{}

		getbugs := "select count(id) from bugs"
		err = conn.GetOne(getbugs).Scan(&pc.CountBugs)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		send, _ := json.Marshal(pc)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
