package cache

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
		var id StatusId
		var name Status
		statusrows.Scan(&id, &name)
		CacheSidStatus[id] = name
		CacheStatusSid[name] = id
	}

	rolerows, err := db.Mconn.GetRows("select id,role,info from roles")
	if err != nil {
		panic(err)
	}
	for rolerows.Next() {
		var id int64
		var name string
		var info string
		rolerows.Scan(&id, &name, &info)
		CacheRidRole[id] = name
		CacheRoleRid[name] = id
		CacheRidInfo[id] = info
	}

	prows, err := db.Mconn.GetRows("select id,name from project")
	if err != nil {
		panic(err)
	}
	for prows.Next() {
		var id ProjectId
		var name Project
		prows.Scan(&id, &name)
		CachePidProject[id] = name
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

	realrows, err := db.Mconn.GetRows("select id,realname,nickname,email,bugsid,showstatus,rid,jid from user")
	if err != nil {
		panic(err)
	}
	for realrows.Next() {
		var id int64
		var name string
		var nick string
		var email string
		var bugsid int64
		var rid int64
		var jid int64
		var showstatus StoreStatusId
		realrows.Scan(&id, &name, &nick, &email, &bugsid, &showstatus, &rid, &jid)

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
	err = db.Mconn.GetOne("select count(created) from defaultvalue").Scan(&checkdefaultcount)
	if err != nil {
		golog.Error(err)
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
		var id ImportantId
		var name Important
		importantrow.Scan(&id, &name)
		CacheIidImportant[id] = name
		CacheImportantIid[name] = id
	}

	levelrow, err := db.Mconn.GetRows("select id,name from level")
	if err != nil {
		panic(err)
	}

	for levelrow.Next() {
		var id LevelId
		var name Level
		levelrow.Scan(&id, &name)
		CacheLidLevel[id] = name
		CacheLevelLid[name] = id
	}

	//默认值
	var created, complete StatusId
	err = db.Mconn.GetOne("select created, completed from defaultvalue").Scan(&created, &complete)
	if err != nil {
		panic(err)
	}
	DefaultCreateSid = created
	DefaultCompleteSid = complete

}
