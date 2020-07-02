package assist

import (
	"itflow/cache"
	"strconv"
	"strings"
)

type Names []string

func (names Names) RealNameToUsers() string {
	// nickname 转 id的字符串拼接
	ul := make([]string, 0)
	for _, v := range names {
		if udd, ok := cache.CacheRealNameUid[v]; ok {
			ul = append(ul, strconv.FormatInt(udd, 10))
		}
	}
	return strings.Join(ul, ",")
}

func FormatUserlistToShow(userlist string) Names {
	// 用户id转 真实姓名
	al := make([]string, 0)
	ul := strings.Split(userlist, ",")
	for _, v := range ul {
		uid, _ := strconv.Atoi(v)
		al = append(al, cache.CacheUidRealName[int64(uid)])
	}
	return al
}
