package model

import "itflow/db"

// keyvalue 值
type KeyName struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func GetUserKeyName() ([]KeyName, error) {
	rows, err := db.Mconn.GetRows("select id,nickname from user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	kns := make([]KeyName, 0)
	for rows.Next() {
		kn := KeyName{}
		err = rows.Scan(&kn.ID, &kn.Name)
		if err != nil {
			continue
		}
		kns = append(kns, kn)
	}
	return kns, nil
}
