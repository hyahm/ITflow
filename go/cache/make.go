package cache

import (
	"database/sql"
	"itflow/db"
)

func initCache() {

	rolerows, err := db.Mconn.GetRows("select id, name from roles")
	if err != nil {
		panic(err)
	}
	for rolerows.Next() {
		var id int64
		var name string
		rolerows.Scan(&id, &name)
		CacheRidRole[id] = name
		CacheRoleRid[name] = id
	}
	rolerows.Close()
	// 	//默认值
	err = db.Mconn.GetOne("select ifnull(min(created),0), ifnull(min(completed),0) from defaultvalue").Scan(&DefaultCreateSid, &DefaultCompleteSid)
	if err != nil {
		if err != sql.ErrNoRows {
			panic(err)
		}

	}
}
