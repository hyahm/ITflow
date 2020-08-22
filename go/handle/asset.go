package handle

import (
	"errors"
	"itflow/cache"
	"itflow/classify"
	"itflow/db"
	"strings"
	"time"
)

var NotFoundToken = errors.New("not found token")

func sortpermlist(permlist []string) []string {
	l := len(cache.CacheSidStatus)

	newlist := make([]string, 0)
	for i := 0; i < l; i++ {
		for _, v := range permlist {
			if cache.CacheSidStatus[cache.StatusId(i)].ToString() == v {
				newlist = append(newlist, v)
			}
		}
	}
	return newlist
}

// 插入到log表中
func InsertLog(classify classify.Classify, content string, ipaddr string) error {
	logsql := "insert into log(exectime,classify,content,ip) values(?,?,?,?)"
	ip := strings.Split(ipaddr, ":")[0]
	if ip != "127.0.0.1" {
		_, err := db.Mconn.Insert(logsql, time.Now().Unix(), classify, content, ip)
		if err != nil {
			return err
		}
	}

	return nil
}
