package search

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"itflow/db"
	"itflow/model"
	"strconv"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
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
	err := db.Mconn.GetOne(pl.CountSql).Scan(&pl.Count)
	if err != nil {
		golog.Error(err)
		return nil, err
	}
	// 增加显示的状态
	start, end := xmux.GetLimit(pl.Count, pl.Page, pl.Limit)
	golog.Info(start)
	golog.Info(pl.Page)
	golog.Info(end)
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

		var userlist string
		rows.Scan(&one.ID, &one.Date, &one.Importance, &one.Status, &one.Title, &one.Level,
			&one.Projectname, &one.Env, &userlist, &one.Author)
		// 如果不存在这么办， 添加修改的时候需要判断

		rows, err := db.Mconn.GetRows(fmt.Sprintf("select realname from user where id in ('%s')", userlist))
		if err != nil {
			golog.Error(err)
			return al.ErrorE(err)
		}
		for rows.Next() {
			realname := new(string)
			err = rows.Scan(realname)
			if err != nil {
				golog.Error(err)
				continue
			}
			one.Handle = append(one.Handle, *realname)
		}

		al.Al = append(al.Al, one)

	}
	al.Count = pl.Count
	al.Page = pl.Page
	return al.Marshal()
}

func (pl *BugList) myTaskRows() (*sql.Rows, error) {
	var err error
	golog.Info(pl.CountSql)
	countrows, err := db.Mconn.GetRows(pl.CountSql)
	if err != nil {
		golog.Error(err)
		return nil, err
	}
	for countrows.Next() {
		var userlist string
		countrows.Scan(&userlist)
		for _, v := range strings.Split(userlist, ",") {
			//判断用户是否存在，不存在就 删吗 ， 先不删
			userid64, _ := strconv.ParseInt(v, 10, 64)
			if userid64 == pl.Uid {
				pl.Count++
				break
			}
		}

	}
	// 增加显示的状态

	pl.ListSql += " order by id desc "
	var rows *sql.Rows
	golog.Info(pl.ListSql)
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
	start, end := xmux.GetLimit(pl.Count, pl.Page, pl.Limit)
	var timer int
	for rows.Next() {
		one := &model.ArticleList{}

		var userlist string
		rows.Scan(&one.ID, &one.Date, &one.Importance, &one.Status, &one.Title, &one.Level, &one.Projectname, &one.Env, &userlist, &one.Author)

		var isMyTask bool

		for _, v := range strings.Split(userlist, ",") {
			//判断用户是否存在，不存在就 删吗 ，
			userid64, _ := strconv.ParseInt(v, 10, 64)
			if userid64 == pl.Uid {
				isMyTask = true
			}

		}

		// }
		if isMyTask {
			// 超过了后面就不用做了
			if timer >= end {
				break
			}
			if timer >= start && timer < end {
				rows, err := db.Mconn.GetRows(fmt.Sprintf("select realname from user where id in ('%s')", userlist))
				if err != nil {
					golog.Error(err)
					return al.ErrorE(err)
				}
				for rows.Next() {
					realname := new(string)
					err = rows.Scan(realname)
					if err != nil {
						golog.Error(err)
						continue
					}
					one.Handle = append(one.Handle, *realname)
				}
			}

			timer++
		}

	}
	al.Count = pl.Count
	al.Page = pl.Page
	return al.Marshal()
}
