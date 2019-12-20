package asset

import (
	"github.com/hyahm/golog"
	"itflow/bug/bugconfig"
	"itflow/db"
	"strings"
)

func CheckPerm(role string, nickname string) (bool, error) {
	rid := bugconfig.CacheUidRid[bugconfig.CacheNickNameUid[nickname]]
	var rolestring string
	row, err := db.Mconn.GetOne("select rolelist from rolegroup where id=?", rid)
	if err != nil {
		golog.Error(err.Error())
		return false, err
	}
	err = row.Scan(&rolestring)
	if err != nil {
		golog.Error(err.Error())
		return false, err
	}
	for _, v := range strings.Split(rolestring, ",") {
		if v == role {
			return true, nil
		}
	}
	return false, nil
}
