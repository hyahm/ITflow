package cache

import (
	"database/sql"
	"itflow/db"
)

var CacheRoleID map[int64]PageInfo

func initCache() {
	CacheRoleID = make(map[int64]PageInfo)
	rolerows, err := db.Mconn.GetRows("select id, name,info from roles")
	if err != nil {
		panic(err)
	}
	for rolerows.Next() {
		var id int64
		var name, info string
		rolerows.Scan(&id, &name, &info)
		CacheRoleID[id] = PageInfo{
			Name: name,
			Info: info,
		}
		// CacheRoleRid[name] = id
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
