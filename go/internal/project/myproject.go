package project

import (
	"encoding/json"
	"itflow/model"
	"strconv"
	"strings"
)

type MyProject struct {
	Name []string `json:"name"`
	Code int      `json:"code"`
	Msg  string   `json:"message"`
}

func (mp *MyProject) Marshal() []byte {
	send, _ := json.Marshal(mp)
	return send
}

func (mp *MyProject) Get(uid int64) []byte {
	pl, err := model.NewProjectList()
	if err != nil {
		mp.Code = 1
		mp.Msg = err.Error()
		return mp.Marshal()
	}

	for _, p := range pl {
		uids := strings.Split(p.Participant, ",")
		if len(uids) == 1 && uids[0] == "" {
			continue
		}
		for _, v := range uids {
			uid64, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				continue
			}
			if uid64 == uid {
				mp.Name = append(mp.Name, p.Name)
				break
			}

		}

	}
	return mp.Marshal()
}
