package project

import (
	"encoding/json"
	"itflow/db"
	"itflow/model"

	"github.com/hyahm/golog"
)

type ResProjectList struct {
	ProjectList []*ReqProject `json:"projectlist"`
	Code        int           `json:"code"`
	Msg         string        `json:"msg"`
}

func (rpl *ResProjectList) Marshal() []byte {
	send, _ := json.Marshal(rpl)
	return send
}

func GetList(uid int64) []byte {
	// 项目列表， 只要跟自己有关的都可以看到
	rpl := &ResProjectList{
		ProjectList: make([]*ReqProject, 0),
	}
	ps, err := model.NewProjectListCheckId(uid)
	if err != nil {
		golog.Error(err)
		rpl.Code = 1
		rpl.Msg = err.Error()
		return rpl.Marshal()
	}
	for _, p := range ps {
		rp := &ReqProject{
			ProjectName: p.Name,
		}
		rp.Id = p.Id
		err = db.Mconn.GetOne("select name from usergroup where id=?", p.UGid).Scan(&rp.GroupName)
		if err != nil {
			golog.Info(err)
			continue
		}
		// ug, ok := cache.CacheUGidUserGroup[p.Gid]
		// if !ok {
		// 	continue
		// }

		rpl.ProjectList = append(rpl.ProjectList, rp)
	}
	return rpl.Marshal()
}
