package search

import (
	"errors"
	"fmt"
	"itflow/cache"
	"itflow/db"
	"itflow/pkg/search"

	"github.com/hyahm/golog"
)

var ErrorNoStatus = errors.New("没选择状态，返回空数组")

type ReqMyBugFilter struct {
	Page        int              `json:"page"`
	Limit       int              `json:"limit"`
	Level       cache.Level      `json:"level"`
	Project     cache.Project    `json:"project"`
	Title       string           `json:"title"`
	ShowsStatus cache.StatusList // 这个应该从数据库获取
}

func (rmf *ReqMyBugFilter) GetUsefulCondition(uid int64, countsql, searchsql string) (*search.BugList, error) {
	// 搜索所有跟bug相关的处理方法
	golog.Info(rmf.Limit)
	if rmf.Level != "" {
		// 判断这个值是否存在

		if lid := rmf.Level.Id(); lid != 0 {
			countsql += fmt.Sprintf("and lid=%d ", lid)
			searchsql += fmt.Sprintf("and lid=%d ", lid)
		} else {
			rmf.Level = ""
		}
	}
	if rmf.Title != "" {

		searchsql += fmt.Sprintf("and title like '%s' ", rmf.Title)
		countsql += fmt.Sprintf("and title like '%s' ", rmf.Title)

	}

	if rmf.Project != "" {
		// 判断这个值是否存在
		if pid, ok := cache.CacheProjectPid[rmf.Project]; ok {
			searchsql += fmt.Sprintf("and pid=%d ", pid)
			countsql += fmt.Sprintf("and pid=%d ", pid)
		} else {
			rmf.Level = ""
		}
	}
	// 获取此用户能看到的状态
	var statuslist string
	err := db.Mconn.GetOne("select showstatus from user where id=?", uid).Scan(&statuslist)
	if err != nil {
		return nil, err
	}
	insql := fmt.Sprintf("and sid in ('%s')", statuslist)
	countsql += insql
	searchsql += insql

	sb := &search.BugList{
		CountSql: countsql,
		ListSql:  searchsql,
		Uid:      uid,
		Limit:    rmf.Limit,
		Page:     rmf.Page,
	}

	return sb, nil
}
