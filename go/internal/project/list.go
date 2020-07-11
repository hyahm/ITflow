package project

import (
	"encoding/json"
	"itflow/cache"
	"itflow/model"
	"strconv"
	"strings"
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

func GetList() []byte {
	rpl := &ResProjectList{
		ProjectList: make([]*ReqProject, 0),
	}
	ps, err := model.NewProjectList()
	if err != nil {
		rpl.Code = 1
		rpl.Msg = err.Error()
		return rpl.Marshal()
	}

	for _, p := range ps {
		rp := &ReqProject{
			SelectUser: make([]string, 0),
		}
		rp.Id = p.Id
		rp.ProjectName = p.Name
		// 防止空数据错误， 正常环境是不会出现这个问题的
		pl := strings.Split(p.Participant, ",")
		if len(pl) == 1 && pl[0] == "" {
			rpl.ProjectList = append(rpl.ProjectList, rp)
			continue
		}
		for _, v := range pl {
			uid, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				rpl.Code = 1
				rpl.Msg = err.Error()
				return rpl.Marshal()
			}
			rp.SelectUser = append(rp.SelectUser, cache.CacheUidRealName[uid])
		}
		rpl.ProjectList = append(rpl.ProjectList, rp)
	}
	return rpl.Marshal()
}
