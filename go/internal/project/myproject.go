package project

import (
	"encoding/json"
	"itflow/db"
	"itflow/model"
	"strconv"
	"strings"

	"github.com/hyahm/golog"
)

type MyProject struct {
	Name     []string `json:"name"`
	Versions []string `json:"versions"`
	Code     int      `json:"code"`
	Msg      string   `json:"msg"`
}

func (mp *MyProject) Marshal() []byte {
	send, err := json.Marshal(mp)
	if err != nil {
		golog.Error(err)
	}
	return send
}

func (mp *MyProject) Error(msg string) []byte {
	mp.Code = 1
	mp.Msg = msg
	return mp.Marshal()
}

func (mp *MyProject) ErrorE(err error) []byte {
	return mp.Error(err.Error())
}

func (mp *MyProject) Get(uid int64) []byte {
	pl, err := model.NewProjectListCheckId(uid)
	if err != nil {
		mp.Code = 1
		mp.Msg = err.Error()
		return mp.Marshal()
	}

	for _, p := range pl {
		var uids string
		err = db.Mconn.GetOne("select ids from usergroup where id=?", p.Gid).Scan(&uids)
		if err != nil {
			golog.Error(err)
			continue
		}

		for _, v := range strings.Split(uids, ",") {
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
