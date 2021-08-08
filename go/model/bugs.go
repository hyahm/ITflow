package model

import (
	"html"
	"itflow/db"
)

type Bug struct {
	ID         int64   `json:"id" db:"id,default"`
	UID        int64   `json:"uid" db:"uid"`
	Title      string  `json:"title" db:"title"`
	Sid        int64   `json:"sid" db:"sid"` // bug状态id
	Content    string  `json:"content" db:"content"`
	OwnerId    int64   `json:"ownerid" db:"ownerid"` // 创建者
	Iid        int64   `json:"iid" db:"iid"`         // import id
	CreateTime int64   `json:"createtime" db:"createtime"`
	Uids       []int64 `json:"spusers" db:"spusers"` // users id
	Lid        int64   `json:"lid" db:"lid"`         // level id
	Eid        int64   `json:"eid" db:"eid"`         // env id
	Tid        int64   `json:"tid" db:"tid"`         // type id
	Pid        int64   `json:"pid" db:"pid"`         // project id
	UpdateTime int64   `json:"updatetime" db:"updatetime"`
	Dustbin    bool    `json:"dustbin" db:"dustbin"`
}

func GetBugById(id interface{}, uid int64) (*Bug, error) {
	bug := &Bug{}
	err := db.Mconn.Select(&bug, "select * from bugs where id=? and ownerid=?", id, uid)
	if err != nil {
		return nil, err
	}
	return bug, nil
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

	_, err := db.Mconn.Update(getlistsql, id)

	return err
}

func (bug *Bug) Delete(id interface{}) error {
	getlistsql := "delete from bugs  where id=?"

	_, err := db.Mconn.Update(getlistsql, id)

	return err
}

func (bug *Bug) CreateBug() (err error) {
	insertsql := "insert into bugs($key) values($value)"
	ids, err := db.Mconn.InsertInterfaceWithID(bug, insertsql)
	if err != nil {
		return
	}
	bug.ID = ids[0]
	return
}

func (bug *Bug) EditBug() (err error) {
	bug.Content = html.EscapeString(bug.Content)
	insertsql := "update bugs set  $set where id=?"
	_, err = db.Mconn.UpdateInterface(bug, insertsql, bug.ID)
	return
}

func GetCount(sql string, args ...interface{}) (int, error) {
	var count int
	err := db.Mconn.GetOneIn(sql, args...).Scan(&count)
	return count, err
}

func GetAllBug(sql string, args ...interface{}) ([]Bug, error) {
	bugs := make([]Bug, 0)
	err := db.Mconn.SelectIn(&bugs, sql, args...)
	return bugs, err
}
