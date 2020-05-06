package bugconfig

import (
	"errors"
	"itflow/db"

	"github.com/hyahm/golog"
)

func initCache() {

	statusrows, err := db.Mconn.GetRows("select id,name from status")
	if err != nil {
		panic(err)
	}
	for statusrows.Next() {
		var id int64
		var name string
		statusrows.Scan(&id, &name)
		CacheSidStatus[id] = name
		CacheStatusSid[name] = id
	}

	rolerows, err := db.Mconn.GetRows("select id,role from roles")
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

	prows, err := db.Mconn.GetRows("select id,name from projectname")
	if err != nil {
		panic(err)
	}
	for prows.Next() {
		var id int64
		var name string
		prows.Scan(&id, &name)
		CachePidName[id] = name
		CacheProjectPid[name] = id
	}

	jobrows, err := db.Mconn.GetRows("select id,name from jobs")
	if err != nil {
		panic(err)
	}
	for jobrows.Next() {
		var id int64
		var name string
		jobrows.Scan(&id, &name)
		CacheJidJobname[id] = name
		CacheJobnameJid[name] = id
	}

	erows, err := db.Mconn.GetRows("select id,envname from environment")
	if err != nil {
		panic(err)
	}
	for erows.Next() {
		var id int64
		var name string
		erows.Scan(&id, &name)
		CacheEidName[id] = name
		CacheEnvNameEid[name] = id
	}

	headerrows, err := db.Mconn.GetRows("select id,name from header")
	if err != nil {
		panic(err)
	}
	for headerrows.Next() {
		var id int64
		var name string
		headerrows.Scan(&id, &name)
		CacheHidHeader[id] = name
		CacheHeaderHid[name] = id
	}

	realrows, err := db.Mconn.GetRows("select id,realname,nickname,email,bugsid,level,showstatus,rid,jid from user")
	if err != nil {
		panic(err)
	}
	for realrows.Next() {
		var id int64
		var name string
		var nick string
		var email string
		var bugsid int64
		var level int
		var rid int64
		var jid int64
		var showstatus string
		realrows.Scan(&id, &name, &nick, &email, &bugsid, &level, &showstatus, &rid, &jid)
		if level == 0 {
			SUPERID = id
		}
		CacheUidSgid[id] = bugsid // 状态组id
		CacheUidNickName[id] = nick
		CacheUidRealName[id] = name
		CacheUidRid[id] = rid
		CacheUidJid[id] = jid
		CacheNickNameUid[nick] = id
		CacheRealNameUid[name] = id
		CacheUidFilter[id] = showstatus
		CacheUidEmail[id] = email

	}

	versionrows, err := db.Mconn.GetRows("select id,name from version")
	if err != nil {
		panic(err)
	}
	for versionrows.Next() {
		var vid int64
		var role string
		versionrows.Scan(&vid, &role)
		CacheVersionNameVid[role] = vid
		CacheVidName[vid] = role

	}

	deprows, err := db.Mconn.GetRows("select id,name from statusgroup")
	if err != nil {
		panic(err)
	}
	for deprows.Next() {
		var pid int64
		var name string
		deprows.Scan(&pid, &name)
		CacheSgidGroup[pid] = name
	}
	rgrows, err := db.Mconn.GetRows("select id,name from rolegroup")
	if err != nil {
		panic(err)
	}
	for rgrows.Next() {
		var id int64
		var name string
		rgrows.Scan(&id, &name)
		CacheRidGroup[id] = name
	}
	grouprows, err := db.Mconn.GetRows("select id,name from usergroup")
	if err != nil {
		panic(err)
	}

	for grouprows.Next() {
		var id int64
		var name string
		grouprows.Scan(&id, &name)
		CacheGidGroup[id] = name
	}

	typerows, err := db.Mconn.GetRows("select id,name from types")
	if err != nil {
		panic(err)
	}

	for typerows.Next() {
		var id int64
		var name string
		typerows.Scan(&id, &name)
		CacheTidName[id] = name
		CacheNameTid[name] = id
	}

	//检查默认值是否只有一行
	var checkdefaultcount int
	row, err := db.Mconn.GetOne("select count(status) from defaultvalue")
	if err != nil {
		golog.Error(err)
		panic(err)
	}
	err = row.Scan(&checkdefaultcount)
	if err != nil {
		panic(err)
	}
	if checkdefaultcount != 1 {
		panic(errors.New("defaultvalue表行数只能是1"))
	}
	importantrow, err := db.Mconn.GetRows("select id,name from importants")
	if err != nil {
		panic(err)
	}

	for importantrow.Next() {
		var id int64
		var name string
		importantrow.Scan(&id, &name)
		CacheIidImportant[id] = name
		CacheImportantIid[name] = id
	}

	levelrow, err := db.Mconn.GetRows("select id,name from level")
	if err != nil {
		panic(err)
	}

	for levelrow.Next() {
		var id int64
		var name string
		levelrow.Scan(&id, &name)
		CacheLidLevel[id] = name
		CacheLevelLid[name] = id
	}

	//默认值
	var status int64
	var important int64
	var level int64
	row, err = db.Mconn.GetOne("select status,important,level from defaultvalue")
	if err != nil {
		panic(err)
	}
	err = row.Scan(&status, &important, &level)
	if err != nil {
		panic(err)
	}
	if _, ok := CacheSidStatus[status]; ok {
		CacheDefault["status"] = status
	}

	if _, ok := CacheIidImportant[important]; ok {
		CacheDefault["important"] = important
	}
	if _, ok := CacheLidLevel[level]; ok {
		CacheDefault["level"] = level
	}

}
