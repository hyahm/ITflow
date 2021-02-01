package bug

import (
	"encoding/json"
	"html"
	"itflow/cache"
	"itflow/db"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/gomysql"
)

// 前端编辑时需要的数据结构
type EditBug struct {
	Title       string   `json:"title,omitempty"`
	Content     string   `json:"content,omitempty"`
	Id          int64    `json:"id"`
	Selectusers []string `json:"selectuser,omitempty"`
	Important   string   `json:"important,omitempty"`
	Level       string   `json:"level,omitempty"`
	Projectname string   `json:"projectname,omitempty"`
	Envname     string   `json:"envname,omitempty"`
	Version     string   `json:"version,omitempty"`
	Code        int      `json:"code"`
	Msg         string   `json:"msg,omitempty"`
	Typ         int      `json:"typ"`
}

func (reb *EditBug) Marshal() []byte {
	send, err := json.Marshal(reb)
	if err != nil {
		golog.Error(err)
	}
	return send
}

func (reb *EditBug) Error(msg string) []byte {
	reb.Code = 1
	reb.Msg = msg
	return reb.Marshal()
}

func (reb *EditBug) ErrorE(err error) []byte {
	return reb.Error(err.Error())
}

func BugById(id string, uid int64) []byte {
	reb := &EditBug{}
	golog.Info(id)
	alsql := `select b.tid,ifnull(i.name, ''),title,
	ifnull(l.name, ''), ifnull(p.name, ''), ifnull(e.name, ''),spusers,ifnull(v.name, ''),content, b.uid 
	from bug.bugs as b 
	left join bug.importants as i on i.id=b.iid
	left join bug.level as l on l.id=b.lid
	left join bug.project as p  on p.id=b.pid
	left join bug.environment as e  on e.id=b.eid
	left join version as v  on v.id=b.vid
	where b.id=? `
	var spids string
	var bugUid int64
	err := db.Mconn.GetOne(alsql, id).Scan(&reb.Typ, &reb.Important, &reb.Title, &reb.Level, &reb.Projectname,
		&reb.Envname, &spids, &reb.Version, &reb.Content, &bugUid)
	if err != nil {
		golog.Error(err)
		return reb.ErrorE(err)
	}
	golog.Info(reb.Typ)
	if bugUid != uid && uid != cache.SUPERID {
		return reb.Error("没有权限")
	}
	rows, err := db.Mconn.GetRowsIn("select realname from user where id in (?)",
		(gomysql.InArgs)(strings.Split(spids, ",")).ToInArgs())
	if err != nil {
		golog.Error(err)
		return reb.ErrorE(err)
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			golog.Info(err)
			continue
		}
		reb.Selectusers = append(reb.Selectusers, name)
	}
	rows.Close()
	reb.Content = html.UnescapeString(reb.Content)
	return reb.Marshal()
}

// func (reb *EditBug) CheckBug() (*model.Bug, error) {
// 	// 将获取的数据转为可以存表的数据
// 	bug := &model.Bug{}
// 	bug.ID = reb.Id
// 	reb.Envname = reb.Envname.Trim()
// 	reb.Important = reb.Important.Trim()
// 	bug.Content = reb.Content
// 	bug.Title = reb.Title
// 	reb.Level = reb.Level.Trim()
// 	reb.Projectname = reb.Projectname.Trim()
// 	reb.Version = reb.Version.Trim()
// 	if reb.Envname == "" || reb.Important == "" || reb.Level == "" ||
// 		reb.Projectname == "" || reb.Version == "" {
// 		return nil, errors.New("all name not by empty")
// 	}
// 	var ok bool
// 	bug.LevelId = reb.Level.Id()
// 	if bug.LevelId == 0 {
// 		return nil, errors.New("没有找到level key")
// 	}
// 	bug.ProjectId = reb.Projectname.Id()
// 	if bug.ProjectId == 0 {
// 		return nil, errors.New("没有找到project key")
// 	}
// 	//
// 	if bug.EnvId, ok = cache.CacheEnvEid[reb.Envname]; !ok {
// 		return nil, errors.New("没有找到env key")
// 	}
// 	//
// 	bug.ImportanceId = reb.Important.Id()
// 	if bug.ImportanceId == 0 {
// 		return nil, errors.New("没有找到important key")
// 	}
// 	if bug.VersionId, ok = cache.CacheVersionVid[reb.Version]; !ok {
// 		return nil, errors.New("没有找到version key")
// 	}

// 	bug.OprateUsers = reb.Selectusers.RealNameToUsersId()
// 	return bug, nil
// }
