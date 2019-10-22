package asset

import (
	"github.com/hyahm/golog"
	"itflow/bug/bugconfig"
	"strings"
)

func CheckPerm(role string, nickname string) (bool, error) {
	rid := bugconfig.CacheUidRid[bugconfig.CacheNickNameUid[nickname]]
	var rolestring string
	err := bugconfig.Bug_Mysql.GetOne("select rolelist from rolegroup where id=?", rid).Scan(&rolestring)
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
