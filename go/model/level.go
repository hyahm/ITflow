package model

import (
	"encoding/json"
	"itflow/db"

	"github.com/hyahm/golog"
)

type Level struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func GetLevelKeyNameByUid() ([]KeyName, error) {
	rows, err := db.Mconn.GetRows("select id,name from level")
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

type RequestLevel struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Code int    `json:"code"`
}

type Levels struct {
	Levels []*Level `json:"levels"`
	Code   int      `json:"code"`
	Msg    string   `json:"msg"`
}

type Update_level struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	OldName string `json:"oldname"`
}

func (ll *Levels) Marshal() []byte {
	send, err := json.Marshal(ll)
	if err != nil {
		golog.Error(err)
	}
	return send
}

func (ll *Levels) Error(msg string) []byte {
	ll.Code = 1
	ll.Msg = msg
	return ll.Marshal()
}

func (ll *Levels) ErrorE(err error) []byte {
	return ll.Error(err.Error())
}
