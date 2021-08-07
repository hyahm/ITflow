package model

import (
	"encoding/json"
	"itflow/db"

	"github.com/hyahm/golog"
)

type Important struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func GetImportantKeyNameByUid() ([]KeyName, error) {
	rows, err := db.Mconn.GetRows("select id,name from importants")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	kns := make([]KeyName, 0)
	for rows.Next() {
		kn := KeyName{}
		err = rows.Scan(&kn.ID, &kn.Name)
		if err != nil {
			golog.Error(err)
			continue
		}
		kns = append(kns, kn)
	}
	return kns, nil
}

type ResposeImportant struct {
	ImportantList []*Important `json:"importantlist"`
	Code          int          `json:"code"`
	Msg           string       `json:"msg"`
}

func (li *ResposeImportant) Marshal() []byte {
	send, err := json.Marshal(li)
	if err != nil {
		golog.Error(err)
	}
	return send
}
func (li *ResposeImportant) Error(msg string) []byte {
	li.Code = 1
	li.Msg = msg
	return li.Marshal()
}

func (li *ResposeImportant) ErrorE(err error) []byte {

	return li.Error(err.Error())
}
