package search

import (
	"errors"
	"fmt"
	"itflow/cache"
	"itflow/pkg/search"

	"github.com/hyahm/golog"
)

type ReqMyBugFilter struct {
	Page        int              `json:"page"`
	Limit       int              `json:"limit"`
	Level       cache.Level      `json:"level"`
	Project     string           `json:"project"`
	Title       string           `json:"title"`
	ShowsStatus cache.StatusList // 这个应该从数据库获取
}

func (rmf *ReqMyBugFilter) GetUsefulCondition(uid int64, countsql, searchsql string) (*search.BugList, error) {
	// 搜索所有跟bug相关的处理方法

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
	golog.Info(uid)
	showstatus := cache.CacheUidFilter[uid].ToShow()
	if len(showstatus) == 0 {
		// 没选择状态，返回空数组

		golog.Info("没选择状态，返回空数组")
		return nil, errors.New("没选择状态，返回空数组")
	}

	countsql += fmt.Sprintf("and sid in (%s)", cache.CacheUidFilter[uid])
	searchsql += fmt.Sprintf("and sid in (%s) ", cache.CacheUidFilter[uid])
	sb := &search.BugList{
		CountSql: countsql,
		ListSql:  searchsql,
		Uid:      uid,
		Limit:    rmf.Limit,
		Page:     rmf.Page,
	}

	return sb, nil
}

// func (rmf *ReqMyBugFilter) GetUsefulConditionWithoutUid(countsql, searchsql string) (*search.BugList, error) {
// 	// 搜索所有跟bug相关的处理方法

// 	if rmf.Level != "" {
// 		// 判断这个值是否存在

// 		if lid := rmf.Level.Id(); lid != 0 {
// 			countsql += fmt.Sprintf("and lid=%d ", lid)
// 			searchsql += fmt.Sprintf("and lid=%d ", lid)
// 		} else {
// 			rmf.Level = ""
// 		}
// 	}
// 	if rmf.Title != "" {

// 		searchsql += fmt.Sprintf("and title like '%s' ", rmf.Title)
// 		countsql += fmt.Sprintf("and title like '%s' ", rmf.Title)

// 	}

// 	if rmf.Project != "" {
// 		// 判断这个值是否存在
// 		if pid, ok := cache.CacheProjectPid[rmf.Project]; ok {
// 			searchsql += fmt.Sprintf("and pid=%d ", pid)
// 			countsql += fmt.Sprintf("and pid=%d ", pid)
// 		} else {
// 			rmf.Level = ""
// 		}
// 	}
// 	// 获取此用户能看到的状态

// 	// countsql += fmt.Sprintf("and sid in (%s)", cache.CacheUidFilter[uid])
// 	// searchsql += fmt.Sprintf("and sid in (%s) ", cache.CacheUidFilter[uid])
// 	sb := &search.BugList{
// 		CountSql: countsql,
// 		ListSql:  searchsql,
// 		Limit:    rmf.Limit,
// 		Page:     rmf.Page,
// 	}

// 	return sb, nil
// }
