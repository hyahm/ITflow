package model

import (
	"html"
	"itflow/db"

	"github.com/hyahm/golog"
)

type ArticleList struct {
	ID          int      `json:"id"`
	Date        int64    `json:"date"`
	Author      string   `json:"author"`
	Importance  string   `json:"importance"`
	Status      string   `json:"status"`
	Title       string   `json:"title"`
	Action      string   `json:"action"`
	Dustbin     int      `json:"dustbin"`
	Level       string   `json:"level"`
	Projectname string   `json:"projectname"`
	Env         string   `json:"env"`
	Handle      []string `json:"handle"`
}

type AllArticleList struct {
	Al    []*ArticleList `json:"articlelist"`
	Code  int            `json:"code"`
	Count int            `json:"total"`
	Page  int            `json:"page"`
}

type Bug struct {
	ID           int64
	Uid          int64
	StatusId     int64 // sid
	Title        string
	Content      string
	ImportanceId int64
	CreateTime   int64
	VersionId    int64
	OprateUsers  string
	LevelId      int64
	EnvId        int64
	ProjectId    int64
	UpdateTime   int64
	Dustbin      bool
}

func (bug *Bug) CreateBug() (err error) {
	insertsql := "insert into bugs(uid,title,sid,content,iid,createtime,lid,pid,eid,spusers,vid) values(?,?,?,?,?,?,?,?,?,?,?)"
	bug.ID, err = db.Mconn.Insert(insertsql,
		bug.Uid, bug.Title, bug.StatusId, html.EscapeString(bug.Content),
		bug.ImportanceId, bug.CreateTime, bug.LevelId,
		bug.ProjectId, bug.EnvId, bug.OprateUsers, bug.VersionId)

	return
}

func (bug *Bug) EditBug() (err error) {
	insertsql := "update bugs set title=?,content=?,iid=?,updatetime=?,lid=?,pid=?,eid=?,spusers=?,vid=? where id=?"

	_, err = db.Mconn.Update(insertsql, bug.Title, html.EscapeString(bug.Content), bug.ImportanceId,
		bug.UpdateTime, bug.LevelId, bug.ProjectId, bug.EnvId, bug.OprateUsers, bug.VersionId, bug.ID)

	return
}

func NewBugById(id interface{}) (*Bug, error) {
	bug := &Bug{}
	alsql := "select iid,title,lid,pid,eid,spusers,vid,content from bugs where id=?"
	err := db.Mconn.GetOne(alsql, id).Scan(&bug.ImportanceId, &bug.Title, &bug.LevelId, &bug.ProjectId,
		&bug.EnvId, &bug.OprateUsers, &bug.VersionId, &bug.Content)
	if err != nil {
		golog.Error(err)
		return nil, err
	}

	return bug, nil
}
