package model

import (
	"encoding/json"
	"html"
	"itflow/db"

	"github.com/hyahm/golog"
)

type ArticleList struct {
	ID          int      `json:"id"`
	Date        int64    `json:"date"`
	Author      string   `json:"author"`
	Importance  string   `json:"important"`
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
	Msg   string         `json:"msg"`
}

func (al *AllArticleList) Marshal() []byte {
	send, err := json.Marshal(al)
	if err != nil {
		golog.Error(err)
	}
	return send
}

func (al *AllArticleList) Error(msg string) []byte {
	al.Code = 1
	al.Msg = msg
	return al.Marshal()
}

func (al *AllArticleList) ErrorE(err error) []byte {
	return al.Error(err.Error())
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
	OprateUsers  int64
	LevelId      int64
	EnvId        int64
	ProjectId    int64
	UpdateTime   int64
	Dustbin      bool
}

func NewBugById(id interface{}) (*Bug, error) {
	bug := &Bug{}

	err := db.Mconn.GetOne("select id, uid,title,sid,content,iid,createtime,vid,spusers,lid,eid,pid,updatetime from bugs where dustbin=true and id=?", id).Scan(
		&bug.ID, &bug.Uid, &bug.Title, &bug.StatusId, &bug.Content, &bug.ImportanceId, &bug.CreateTime, &bug.VersionId, &bug.OprateUsers,
		&bug.LevelId, &bug.EnvId, &bug.ProjectId, &bug.UpdateTime,
	)
	bug.Content = html.UnescapeString(bug.Content)
	return bug, err
}

func GetCreatedCountByTime(start, end, statusid int64) (int, error) {
	var count int
	err := db.Mconn.GetOne("select count(id) from bugs where dustbin=true and createtime between ? and ? and sid=?", start, end, statusid).Scan(&count)
	return count, err
}

func GetCompletedCountByTime(start, end, statusid int64) (int, error) {
	var count int
	err := db.Mconn.GetOne("select count(id) from bugs where dustbin=true and updatetime between ? and ? and sid=?", start, end, statusid).Scan(&count)
	return count, err
}

func (bug *Bug) Resume(id interface{}) error {
	getlistsql := "update bugs set dustbin=false where id=?"

	_, err := db.Mconn.Update(getlistsql, id)

	return err
}

func (bug *Bug) CreateBug() (err error) {
	insertsql := "insert into bugs(uid,title,sid,content,iid,createtime,lid,pid,eid,spusers,vid,dustbin) values(?,?,?,?,?,?,?,?,?,?,?,true)"
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
