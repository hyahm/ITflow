package model

import (
	"html"
	"itflow/cache"
	"itflow/db"
)

type Bug struct {
	ID         int64   `json:"id" db:"id,default"`
	Title      string  `json:"title" db:"title"`
	Sid        int64   `json:"sid" db:"sid"` // bug状态id
	Uid        int64   `json:"uid" db:"uid"` // 创建者
	Content    string  `json:"content" db:"content"`
	Iid        int64   `json:"iid" db:"iid"` // import id
	CreateTime int64   `json:"createtime" db:"createtime"`
	Uids       []int64 `json:"spusers" db:"spusers"` // users id
	Vid        int64   `json:"vid" db:"vid"`         // version id
	Lid        int64   `json:"lid" db:"lid"`         // level id
	Eid        int64   `json:"eid" db:"eid"`         // env id
	Tid        int64   `json:"tid" db:"tid"`         // type id
	Pid        int64   `json:"pid" db:"pid"`         // project id
	UpdateTime int64   `json:"updatetime" db:"updatetime"`
	DeadLine   int64   `json:"deadline" db:"deadline"`
	Dustbin    bool    `json:"dustbin" db:"dustbin"`
}

func GetBugById(id interface{}, uid int64) (*Bug, error) {
	bug := &Bug{}
	result := db.Mconn.Select(&bug, "select * from bugs where id=? and uid=?", id, uid)
	return bug, result.Err
}

func GetCreatedCountByTime(start, end int64) (int, error) {
	var count int
	err := db.Mconn.GetOne("select count(id) from bugs where dustbin=false and createtime between ? and ?", start, end).Scan(&count)
	return count, err
}

func GetCompletedCountByTime(start, end, statusid int64) (int, error) {
	var count int
	err := db.Mconn.GetOne("select count(id) from bugs where dustbin=false and updatetime between ? and ? and sid=?", start, end, statusid).Scan(&count)
	return count, err
}

func (bug *Bug) Resume(id interface{}) error {
	getlistsql := "update bugs set dustbin=false where id=?"

	result := db.Mconn.Update(getlistsql, id)

	return result.Err
}

func (bug *Bug) Update() error {
	getlistsql := "update bugs set $set where id=? and uid=?"
	result := db.Mconn.UpdateInterface(bug, getlistsql, bug.ID, bug.Uid)
	return result.Err
}

func (bug *Bug) UpdateStatus(sids ...int64) error {
	getlistsql := "update bugs set $set where id=? and json_contains(spusers, json_array(?)) and sid not in (?)"
	result := db.Mconn.UpdateInterfaceIn(bug, getlistsql, bug.ID, bug.Uid, sids)
	return result.Err
}

func (bug *Bug) Delete(uid, id interface{}) error {
	if uid == cache.SUPERID {
		getlistsql := "delete from bugs  where id=?"
		result := db.Mconn.Update(getlistsql, id)
		return result.Err
	} else {
		getlistsql := "delete from bugs  where id=? and uid=?"
		result := db.Mconn.Update(getlistsql, id, uid)
		return result.Err
	}

}

func (bug *Bug) CreateBug() (err error) {
	insertsql := "insert into bugs($key) values($value)"
	result := db.Mconn.InsertInterfaceWithID(bug, insertsql)
	if result.Err != nil {
		return result.Err
	}
	bug.ID = result.LastInsertId
	return result.Err
}

func (bug *Bug) EditBug() (err error) {
	bug.Content = html.EscapeString(bug.Content)
	insertsql := "update bugs set  $set where id=?"
	result := db.Mconn.UpdateInterface(bug, insertsql, bug.ID)
	return result.Err
}

func GetCount(sql string, args ...interface{}) (int, error) {
	var count int
	err := db.Mconn.GetOneIn(sql, args...).Scan(&count)
	return count, err
}

func GetAllBug(sql string, args ...interface{}) ([]Bug, error) {
	bugs := make([]Bug, 0)
	result := db.Mconn.SelectIn(&bugs, sql, args...)
	return bugs, result.Err
}
