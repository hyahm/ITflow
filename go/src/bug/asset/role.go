package asset

import (
	"bug/bugconfig"
	"gadb"
	"galog"
	"strings"
)

func CheckPerm(role string, nickname string, conn *gadb.Db) (bool, error) {
	rid := bugconfig.CacheUidRid[bugconfig.CacheNickNameUid[nickname]]
	var rolestring string
	err := conn.GetOne("select rolelist from rolegroup where id=?", rid).Scan(&rolestring)
	if err != nil {
		galog.Error(err.Error())
		return false, err
	}
	for _, v := range strings.Split(rolestring, ",") {
		if v == role {
			return true, nil
		}
	}
	return false, nil
}
