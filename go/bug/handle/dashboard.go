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
	Code        int `json:"code"`
	CountUsers  int `json:"countusers"`
	CountGroups int `json:"countgroups"`
}

func UserCount(w http.ResponseWriter, r *http.Request) {

	_, err := logtokenmysql(r)
	errorcode := &errorstruct{}

	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	uc := &userCount{}

	getusersql := "select count(id) from user"
	getgroupsql := "select count(id) from rolegroup"
	err = bugconfig.Bug_Mysql.GetOne(getusersql).Scan(&uc.CountUsers)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	err = bugconfig.Bug_Mysql.GetOne(getgroupsql).Scan(&uc.CountGroups)
	if err != nil {
		golog.Error(err.Error())
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

	_, err := logtokenmysql(r)
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	pc := &projectCount{}

	getbugs := "select count(id) from bugs"
	err = bugconfig.Bug_Mysql.GetOne(getbugs).Scan(&pc.CountBugs)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(pc)
	w.Write(send)
	return

}
