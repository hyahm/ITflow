package project

import (
	"encoding/json"
	"itflow/cache"
	"itflow/model"
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
	rpl := &ResProjectList{
		ProjectList: make([]*ReqProject, 0),
	}
	ps, err := model.NewProjectListCheckId(uid)
	if err != nil {
		rpl.Code = 1
		rpl.Msg = err.Error()
		return rpl.Marshal()
	}

	for _, p := range ps {
		rp := &ReqProject{
			ProjectName: p.Name,
		}
		rp.Id = p.Id
		rp.GroupName = cache.CacheGidGroup[p.Gid].Name
		// 防止空数据错误， 正常环境是不会出现这个问题的
		// pl := strings.Split(p.Gid, ",")
		// if len(pl) == 1 && pl[0] == "" {
		// 	rpl.ProjectList = append(rpl.ProjectList, rp)
		// 	continue
		// }

		rpl.ProjectList = append(rpl.ProjectList, rp)
	}
	return rpl.Marshal()
}
