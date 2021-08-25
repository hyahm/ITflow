package perm

import (
	"itflow/cache"
	"itflow/db"
	"strings"

	"github.com/hyahm/golog"
)

type OptionPerm struct {
	Select bool
	Insert bool
	Update bool
	Delete bool
}

// func EnvPerm(uid int64) (op OptionPerm, err error) {
// 	return perm(uid, cache.CacheRoleRid["env"])
// }

// func ImportantPerm(uid int64) (op OptionPerm, err error) {
// 	return perm(uid, cache.CacheRoleRid["important"])
// }

// func LevelPerm(uid int64) (op OptionPerm, err error) {
// 	return perm(uid, cache.CacheRoleRid["level"])
// }

// func PositionPerm(uid int64) (op OptionPerm, err error) {
// 	return perm(uid, cache.CacheRoleRid["position"])
// }

// func ProjectPerm(uid int64) (op OptionPerm, err error) {
// 	return perm(uid, cache.CacheRoleRid["project"])
// }

// func StatusPerm(uid int64) (op OptionPerm, err error) {
// 	return perm(uid, cache.CacheRoleRid["status"])
// }

// func StatusgroupPerm(uid int64) (op OptionPerm, err error) {
// 	return perm(uid, cache.CacheRoleRid["statusgroup"])
// }

// func VersionPerm(uid int64) (op OptionPerm, err error) {
// 	return perm(uid, cache.CacheRoleRid["version"])
// }

func perm(uid int64, rid int64) (op OptionPerm, err error) {
	if uid == cache.SUPERID {
		op.Select = true
		op.Update = true
		op.Delete = true
		op.Insert = true
		return
	}
	var permids string
	err = db.Mconn.GetOne("select permids from rolegroup where id = (select rgid from jobs where id=(select jid from user where id=?))", uid).Scan(&permids)
	if err != nil {
		golog.Info(err)
		return
	}
	err = db.Mconn.GetOneIn("select  find, remove, revise, increase from perm where id in (?) and rid=?",
		strings.Split(permids, ","), rid).Scan(
		&op.Select, &op.Delete, &op.Update, &op.Insert,
	)
	if err != nil {
		golog.Info(err)
		return
	}
	return
}
