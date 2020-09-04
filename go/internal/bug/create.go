package bug

import (
	"encoding/json"
	"html"
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

func BugById(id interface{}) []byte {
	reb := &EditBug{}

	alsql := `select i.name,title,
	l.name,p.name,e.name,spusers,v.name,content 
	from bugs as b 
	join importants as i 
	join level as l 
	join project as p  
	join environment as e  
	join version as v  
	on b.id=? and 
	i.id=b.iid and 
	l.id=b.lid and 
	p.id=b.pid and  
	e.id=b.eid and 
	v.id=b.vid`
	var spids string
	err := db.Mconn.GetOne(alsql, id).Scan(&reb.Important, &reb.Title, &reb.Level, &reb.Projectname,
		&reb.Envname, &spids, &reb.Version, &reb.Content)
	if err != nil {
		golog.Error(err)
		return reb.ErrorE(err)
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
			golog.Error(err)
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
