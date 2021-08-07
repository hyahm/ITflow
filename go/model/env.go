package model

import (
	"itflow/db"

	"github.com/hyahm/golog"
)

type environment struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func GetEnvKeyNameByUid() ([]KeyName, error) {
	rows, err := db.Mconn.GetRows("select id,name from environment")
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
