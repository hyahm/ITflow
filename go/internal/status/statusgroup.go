package status

import (
	"itflow/db"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/gomysql"
)

type StatusGroup struct {
	Id         int64    `json:"id" type:"int" need:"是" default:"0" information:"无效"`
	StatusList []string `json:"checklist, omitempty" type:"array" need:"是" default:"[]" information:"状态名列表"`
	Name       string   `json:"name" type:"int" need:"是" default:"0" information:"状态组名"`
}

func (status *StatusGroup) GetIds() (string, error) {
	rows, err := db.Mconn.GetRowsIn("select id from status where name in (?)",
		(gomysql.InArgs)(status.StatusList).ToInArgs())

	if err != nil {
		golog.Error(err)
		return "", err
	}
	ids := make([]string, 0)
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			golog.Error(err)
			continue
		}
		ids = append(ids, id)
	}
	rows.Close()
	return strings.Join(ids, ","), nil
}
