package project

import (
	"encoding/json"
	"itflow/cache"
	"itflow/model"
	"strconv"
	"strings"

	"github.com/hyahm/golog"
)

type MyProject struct {
	Name     []string `json:"name"`
	Versions []string `json:"versions"`
	Code     int      `json:"code"`
	Msg      string   `json:"message"`
}

func (mp *MyProject) Marshal() []byte {
	send, _ := json.Marshal(mp)
	return send
}

func (mp *MyProject) Get(uid int64) []byte {
	pl, err := model.NewProjectListCheckId(uid)
	if err != nil {
		mp.Code = 1
		mp.Msg = err.Error()
		return mp.Marshal()
	}

	for _, p := range pl {
		uids := strings.Split(cache.CacheGidGroup[p.Gid].Uids, ",")
		if len(uids) == 1 && uids[0] == "" {
			continue
		}
		for _, v := range uids {
			uid64, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				continue
			}
			if uid64 == uid {
				mp.Name = append(mp.Name, p.Name.ToString())
				break
			}

		}

	}
	golog.Infof("%+v", mp)
	return mp.Marshal()
}
