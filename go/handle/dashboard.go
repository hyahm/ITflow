package handle

import (
	"itflow/db"
	"itflow/internal/dashboard"
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

type userCount struct {
	CountUsers  int `json:"countusers"`
	CountGroups int `json:"countgroups"`
}

func UserCount(w http.ResponseWriter, r *http.Request) {

	uc := &userCount{}

	getusersql := "select count(id) from user"
	getgroupsql := "select count(id) from rolegroup"
	err := db.Mconn.GetOne(getusersql).Scan(&uc.CountUsers)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}

	err = db.Mconn.GetOne(getgroupsql).Scan(&uc.CountGroups)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}

	xmux.GetInstance(r).Response.(*response.Response).Data = uc
}

func BugCount(w http.ResponseWriter, r *http.Request) {
	xmux.GetInstance(r).Response.(*response.Response).Data = dashboard.GetCount()
}

type projectCount struct {
	CountBugs     int `json:"countbugs"`
	CountComplete int `json:"countcomplete"`
}

func ProjectCount(w http.ResponseWriter, r *http.Request) {

	pc := &projectCount{}

	getbugs := "select count(id) from bugs"
	err := db.Mconn.GetOne(getbugs).Scan(&pc.CountBugs)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	if model.Default.Completed > 0 {
		completesql := "select count(id) from bugs where sid=?"
		err := db.Mconn.GetOne(completesql, model.Default.Completed).Scan(&pc.CountComplete)
		if err != nil {
			golog.Error(err)
			xmux.GetInstance(r).Response.(*response.Response).Code = 1
			xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
			return
		}

	}
	xmux.GetInstance(r).Response.(*response.Response).Data = pc
}
