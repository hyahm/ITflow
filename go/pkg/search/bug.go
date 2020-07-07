package search

import (
	"database/sql"
	"encoding/json"
	"itflow/cache"
	"itflow/db"
	"itflow/model"
	"itflow/pkg/pager"
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

func (pl *BugList) rows() (*sql.Rows, error) {
	var err error
	err = db.Mconn.GetOne(pl.CountSql).Scan(&pl.Count)

	if err != nil {
		golog.Error(err)

		return nil, err
	}
	golog.Infof("----%+v-------\n", pl)
	// 增加显示的状态
	start, end := pager.GetPagingLimitAndPage(pl.Count, pl.Page, pl.Limit)
	pl.ListSql += " order by id desc limit ?,? "
	var rows *sql.Rows

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
		one.Importance = cache.CacheIidImportant[iid]
		one.Status = cache.CacheSidStatus[sid]
		one.Level = cache.CacheLidLevel[lid]
		one.Projectname = cache.CachePidName[pid]
		one.Env = cache.CacheEidName[eid]

		for _, v := range strings.Split(userlist, ",") {
			//判断用户是否存在，不存在就 删吗 ， 先不删
			userid32, _ := strconv.Atoi(v)
			if realname, ok := cache.CacheUidRealName[int64(userid32)]; ok {
				one.Handle = append(one.Handle, realname)
			}
		}

		// }

		one.Author = cache.CacheUidRealName[pl.Uid]
		al.Al = append(al.Al, one)
	}
	return al.Marshal()
	// return rows, nil
}

func (pl *BugList) GetMyTasks() []byte {
	al := &model.AllArticleList{
		Al: make([]*model.ArticleList, 0),
	}
	// 获取所有数据的行
	var count int
	err := db.Mconn.GetOne(pl.CountSql, pl.Uid).Scan(&count)
	if err != nil {
		golog.Error(err)
		al.Msg = err.Error()
		al.Code = 1
		return al.Marshal()
	}
	start, end := pager.GetPagingLimitAndPage(count, pl.Page, pl.Limit)
	pl.ListSql += " limit ?,?"
	rows, err := db.Mconn.GetRows(pl.ListSql, pl.Uid, start, end)
	if err != nil {
		golog.Error(err)
		al.Msg = err.Error()
		al.Code = 1
		return al.Marshal()
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

		// if mytask {
		// 判断是否是自己的任务，先要过滤查询条件，然后查询spusers
		var ismytask bool
		for _, v := range strings.Split(userlist, ",") {
			if v == strconv.FormatInt(pl.Uid, 10) {
				ismytask = true
				break
			}
		}
		if ismytask {
			for _, v := range strings.Split(userlist, ",") {
				//判断用户是否存在，不存在就 删吗 ， 先不删
				userid32, _ := strconv.Atoi(v)
				if realname, ok := cache.CacheUidRealName[int64(userid32)]; ok {
					one.Handle = append(one.Handle, realname)
				}
			}
		} else {
			continue
		}
		// }

		one.Author = cache.CacheUidRealName[pl.Uid]
		al.Al = append(al.Al, one)
	}
	return al.Marshal()
	// return rows, nil
}
