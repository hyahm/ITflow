package search

import (
	"database/sql"
	"encoding/json"
	"itflow/cache"
	"itflow/db"
	"itflow/model"
	"strconv"
	"strings"

	"github.com/hyahm/golog"
)

type BugList struct {
	CountSql string
	ListSql  string
	Page     int
	Uid      int64
	Limit    int
	Count    int
}

func (pl *BugList) GetPagingLimitAndPage() (int, int) {
	// 都小于1了
	if pl.Limit == 0 {
		return 0, 0
	}
	if pl.Page < 1 {
		pl.Page = 1
	}
	// 超出了，返回最大的页码
	if pl.Page*pl.Limit > pl.Count+pl.Limit {
		if pl.Count%pl.Limit == 0 {
			return ((pl.Count / pl.Limit) - 1) * pl.Limit, pl.Limit
		} else {
			return (pl.Count/pl.Limit + 1) * pl.Limit, pl.Count % pl.Limit
		}
	} else {
		// if count%limit == 0 {

		start := (pl.Page - 1) * pl.Limit
		if pl.Count-start < pl.Limit {
			return start, pl.Count - start
		} else {
			return start, pl.Limit
		}

	}
}

func (pl *BugList) rows() (*sql.Rows, error) {
	golog.Info(pl.CountSql)
	err := db.Mconn.GetOne(pl.CountSql).Scan(&pl.Count)
	if err != nil {
		golog.Error(err)
		return nil, err
	}
	golog.Info(pl.Count)
	// 增加显示的状态
	start, end := pl.GetPagingLimitAndPage()
	pl.ListSql += " order by id desc limit ?,? "
	var rows *sql.Rows
	golog.Info(pl.ListSql)
	golog.Info(start)
	golog.Info(end)
	rows, err = db.Mconn.GetRows(pl.ListSql, start, end)

	if err != nil {
		golog.Error(err)
		return nil, err
	}
	return rows, nil
}

func (pl *BugList) GetMyBugs() []byte {

	// 获取所有数据的行

	al := &model.AllArticleList{
		Al: make([]*model.ArticleList, 0),
	}
	rows, err := pl.rows()
	if err != nil {
		al.Msg = err.Error()
		al.Code = 1
		send, _ := json.Marshal(al)
		return send
	}
	for rows.Next() {
		one := &model.ArticleList{}
		var iid cache.ImportantId
		var sid cache.StatusId
		var lid cache.LevelId
		var pid int64
		var eid int64
		var userlist string
		rows.Scan(&one.ID, &one.Date, &iid, &sid, &one.Title, &lid, &pid, &eid, &userlist)
		// 如果不存在这么办， 添加修改的时候需要判断

		for _, v := range strings.Split(userlist, ",") {
			//判断用户是否存在，不存在就 删吗 ， 先不删
			userid64, _ := strconv.ParseInt(v, 10, 64)
			if realname, ok := cache.CacheUidRealName[userid64]; ok {
				one.Handle = append(one.Handle, realname)
			}
		}

		// }
		one.Importance = cache.CacheIidImportant[iid]
		one.Status = cache.CacheSidStatus[sid]
		one.Level = cache.CacheLidLevel[lid]
		one.Projectname = cache.CachePidName[pid]
		one.Env = cache.CacheEidName[eid]
		one.Author = cache.CacheUidRealName[pl.Uid]
		al.Al = append(al.Al, one)

	}
	al.Count = pl.Count
	al.Page = pl.Page
	return al.Marshal()
}

func (pl *BugList) myTaskRows() (*sql.Rows, error) {
	var err error

	countrows, err := db.Mconn.GetRows(pl.CountSql)
	if err != nil {
		golog.Error(err)
		return nil, err
	}
	for countrows.Next() {
		var isMyTask bool
		var userlist string
		countrows.Scan(&userlist)
		for _, v := range strings.Split(userlist, ",") {
			//判断用户是否存在，不存在就 删吗 ， 先不删
			userid64, _ := strconv.ParseInt(v, 10, 64)
			if userid64 == pl.Uid {
				isMyTask = true
			}
		}
		if isMyTask {
			pl.Count++
		}
	}
	// 增加显示的状态

	pl.ListSql += " order by id desc "
	var rows *sql.Rows

	rows, err = db.Mconn.GetRows(pl.ListSql)

	if err != nil {
		golog.Error(err)
		return nil, err
	}
	return rows, nil
}

func (pl *BugList) GetMyTasks() []byte {

	// 获取所有数据的行
	al := &model.AllArticleList{
		Al: make([]*model.ArticleList, 0),
	}
	rows, err := pl.myTaskRows()
	if err != nil {
		al.Msg = err.Error()
		al.Code = 1
		send, _ := json.Marshal(al)
		return send
	}
	start, end := pl.GetPagingLimitAndPage()
	var timer int
	for rows.Next() {
		one := &model.ArticleList{}
		var iid cache.ImportantId
		var sid cache.StatusId
		var lid cache.LevelId
		var pid int64
		var eid int64
		var userlist string
		rows.Scan(&one.ID, &one.Date, &iid, &sid, &one.Title, &lid, &pid, &eid, &userlist)
		// 如果不存在这么办， 添加修改的时候需要判断
		one.Importance = cache.CacheIidImportant[iid]
		one.Status = cache.CacheSidStatus[sid]
		one.Level = cache.CacheLidLevel[lid]
		one.Projectname = cache.CachePidName[pid]
		one.Env = cache.CacheEidName[eid]
		// 显示realname

		//如果是我的任务

		// for _, v := range strings.Split(userlist, ",") {
		// 	//判断用户是否存在，不存在就 删吗 ， 先不删
		// 	userid32, _ := strconv.Atoi(v)
		// 	if realname, ok := cache.CacheUidRealName[int64(userid32)]; ok {
		// 		one.Handle = append(one.Handle, realname)
		// 	}
		// }

		var isMyTask bool
		for _, v := range strings.Split(userlist, ",") {
			//判断用户是否存在，不存在就 删吗 ， 先不删
			userid64, _ := strconv.ParseInt(v, 10, 64)
			if userid64 == pl.Uid {
				isMyTask = true
			}
			if realname, ok := cache.CacheUidRealName[userid64]; ok {
				one.Handle = append(one.Handle, realname)
			}
		}

		// }
		if isMyTask {
			if timer >= start && timer < end {
				one.Importance = cache.CacheIidImportant[iid]
				one.Status = cache.CacheSidStatus[sid]
				one.Level = cache.CacheLidLevel[lid]
				one.Projectname = cache.CachePidName[pid]
				one.Env = cache.CacheEidName[eid]
				one.Author = cache.CacheUidRealName[pl.Uid]
				al.Al = append(al.Al, one)
			}
			timer++
		}

	}
	al.Count = pl.Count
	al.Page = pl.Page
	return al.Marshal()
}