package project

import (
	"encoding/json"
	"itflow/cache"
	"itflow/model"

	"github.com/hyahm/golog"
)

type ResProjectList struct {
	ProjectList []*ReqProject `json:"projectlist"`
	Code        int           `json:"code"`
	Msg         string        `json:"message"`
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
		rpl.Code = 1
		rpl.Msg = err.Error()
		return rpl.Marshal()
	}
	golog.Info(ps)
	for _, p := range ps {
		rp := &ReqProject{
			ProjectName: p.Name,
		}
		rp.Id = p.Id
		ug, ok := cache.CacheUGidUserGroup[p.Gid]
		if !ok {
			continue
		}
		rp.GroupName = ug.Name

		rpl.ProjectList = append(rpl.ProjectList, rp)
	}
	return rpl.Marshal()
}
