package handle

import (
	"encoding/json"
	"itflow/db"
	"itflow/internal/dashboard"
	"itflow/internal/response"
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
	err := db.Mconn.GetOne(getusersql).Scan(&uc.CountUsers)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	err = db.Mconn.GetOne(getgroupsql).Scan(&uc.CountGroups)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(uc)
	w.Write(send)
	return

}

func BugCount(w http.ResponseWriter, r *http.Request) {
	w.Write(dashboard.GetCount())
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
	err := db.Mconn.GetOne(getbugs).Scan(&pc.CountBugs)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(pc)
	w.Write(send)
	return

}
