package handle

import (
	"encoding/json"
	"itflow/db"
	"itflow/network/response"
	"net/http"

	"github.com/hyahm/golog"
)

//type totalProject struct {
//	CountBugs     int `json:"countbugs"`
//	CountComplete int `json:"countcomplete"`
//}

type userCount struct {
	Code        int `json:"code"`
	CountUsers  int `json:"countusers"`
	CountGroups int `json:"countgroups"`
}

func UserCount(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	uc := &userCount{}

	getusersql := "select count(id) from user"
	getgroupsql := "select count(id) from rolegroup"
	row, err := db.Mconn.GetOne(getusersql)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	err = row.Scan(&uc.CountUsers)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	row, err = db.Mconn.GetOne(getgroupsql)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = row.Scan(&uc.CountGroups)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(uc)
	w.Write(send)
	return

}

type projectCount struct {
	CountBugs     int `json:"countbugs"`
	CountComplete int `json:"countcomplete"`
	Code          int `json:"code"`
}

func ProjectCount(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	pc := &projectCount{}

	getbugs := "select count(id) from bugs"
	row, err := db.Mconn.GetOne(getbugs)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = row.Scan(&pc.CountBugs)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(pc)
	w.Write(send)
	return

}
