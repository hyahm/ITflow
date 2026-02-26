package model

import (
	"errors"
	"html"
	"itflow/cache"
	"itflow/db"
	"time"

	"github.com/hyahm/golog"
)

type Bug struct {
	Id          int64     `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"column:title"`
	StatusId    int64     `json:"status_id" gorm:"column:status_id"` // bug状态id
	CreateId    int64     `json:"create_id" gorm:"column:create_id"` // 创建者
	Content     string    `json:"content" gorm:"column:content"`
	ImportantId int64     `json:"important_id" gorm:"column:important_id"` // import id
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
	HandleUid   int64     `json:"handle_uid" gorm:"column:handle_uid"` // 谁的bug
	LevelId     int64     `json:"level_id" gorm:"column:level_id"`     // level id
	EnvId       int64     `json:"env_id" gorm:"column:env_id"`         // env id
	TypeId      int64     `json:"type_id" gorm:"column:type_id"`       // type id
	ProjectId   int64     `json:"project_id" gorm:"column:project_id"` // project id
	UpdateTime  time.Time `json:"update_time" gorm:"column:update_time"`
	DeadLine    time.Time `json:"deadline" gorm:"column:deadline"`
	Dustbin     bool      `json:"dustbin" gorm:"column:dustbin"`
}

func (bug *Bug) TableName() string {
	return "bugs"
}

func GetBugById(id interface{}, uid int64) (*Bug, error) {
	bug := &Bug{}
	result := db.Mconn.Select(&bug, "select * from bugs where id=? and uid=?", id, uid)
	return bug, result.Err
}

func GetCreatedCountByTime(start, end int64) (int, error) {
	var count int
	err := db.Mconn.GetOne("select count(id) from bugs where dustbin=false and create_time between ? and ?", start, end).Scan(&count)
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
	result := db.Mconn.UpdateInterface(bug, getlistsql, bug.Id, bug.CreateId)
	return result.Err
}

func (bug *Bug) UpdateStatus(sids ...int64) error {
	getlistsql := "update bugs set $set where id=? and json_contains(spusers, json_array(?)) and sid not in (?)"
	result := db.Mconn.UpdateInterfaceIn(bug, getlistsql, bug.Id, bug.CreateId, sids)
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

func (bug *Bug) CreateBug() error {
	bug.CreateTime = time.Now()
	bug.UpdateTime = time.Now()
	return db.Gorm.Table(bug.TableName()).Omit("deadline").Create(bug).Error
	// insertsql := "insert into bugs($key) values($value)"
	// result := db.Mconn.InsertInterfaceWithID(bug, insertsql)
	// if result.Err != nil {
	// 	return result.Err
	// }
	// bug.ID = result.LastInsertId
	// return result.Err
}

func (bug *Bug) EditBug() (err error) {
	bug.Content = html.EscapeString(bug.Content)
	insertsql := "update bugs set  $set where id=?"
	result := db.Mconn.UpdateInterface(bug, insertsql, bug.Id)
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

var ErrorNoStatus = errors.New("没选择状态，返回空数组")

type ReqMyBugFilter struct {
	Page        int     `json:"page"`
	Limit       int     `json:"limit"`
	PageType    int     `json:"page_type"`
	LevelId     int64   `json:"level_id"`
	ProjectId   int64   `json:"project_id"`
	Title       string  `json:"title"`
	ShowsStatus []int64 `json:"showstatus"`
}

type BugResponse struct {
	Id         int64     `json:"id" gorm:"primaryKey"`
	Title      string    `json:"title" gorm:"column:title"`
	StatusName string    `json:"status_name" gorm:"column:status_name"` // bug状态id
	Create     string    `json:"create_name" gorm:"column:create_name"` // 创建者
	Content    string    `json:"content" gorm:"column:content"`
	Important  string    `json:"important_name" gorm:"column:important_name"` // import id
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	Handle     string    `json:"handle_name" gorm:"column:handle_name"`   // 谁的bug
	Level      string    `json:"level_name" gorm:"column:level_name"`     // level id
	Env        string    `json:"env_name" gorm:"column:env_name"`         // env id
	Type       string    `json:"type_name" gorm:"column:type_name"`       // type id
	Project    string    `json:"project_name" gorm:"column:project_name"` // project id
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
	DeadLine   time.Time `json:"deadline" gorm:"column:deadline "`
	Dustbin    bool      `json:"dustbin" gorm:"column:dustbin"`
}

func (bug *Bug) GetUsefulCondition(rmf ReqMyBugFilter, uid int64) ([]BugResponse, int64, error) {
	// 获取需要拼接的sql条件语句, 参数
	// 确定前面已经有了where 或者on
	golog.Infof("%+v", rmf)
	bugs := make([]BugResponse, 0)
	query := db.Gorm.Table(bug.TableName()).
		Joins("left join level on bugs.level_id = level.id").
		Joins("left join status on bugs.status_id = status.id").
		Joins("left join user AS a on bugs.handle_uid = a.id").
		Joins("left join user AS b on bugs.create_id = b.id").
		Joins("left join importants on bugs.important_id = importants.id").
		Joins("left join environment on bugs.env_id = environment.id").
		Joins("left join typ on bugs.type_id = typ.id").
		Joins("left join  project on bugs.project_id = project.id").
		Select(
			`bugs.id, bugs.title, status.name AS status_name,
			b.nickname AS create_name, bugs.content,
			importants.name AS important_name, bugs.create_time,
			a.nickname AS handle_name, level.name AS level_name,
			environment.name AS env_name, typ.name AS type_name,
			project.name AS project_name, bugs.update_time,
			bugs.deadline, bugs.dustbin`,
		)

	if rmf.LevelId > 0 {
		query = query.Where("bugs.level_id=?", rmf.LevelId)
	}

	if rmf.Title != "" {
		query = query.Where("bugs.title like ?", "%"+rmf.Title+"%")
	}

	if rmf.ProjectId > 0 {
		// 判断这个值是否存在
		query = query.Where("bugs.project_id=?", rmf.ProjectId)
	}
	switch rmf.PageType {
	case 1:
		query = query.Where("bugs.dustbin=1")
	case 2:
		query = query.Where("bugs.create_id=?", uid).
			Where("bugs.dustbin=false")
	case 4:
		query = query.Where("bugs.handle_uid=?", uid).Where("bugs.dustbin=false")
	default:
		query = query.Where("bugs.dustbin=false")
	}
	if len(rmf.ShowsStatus) > 0 {
		query = query.Where("bugs.status_id in ?", rmf.ShowsStatus)
	}
	var count int64
	err := query.Count(&count).Error
	if err != nil {
		golog.Error(err)
		return bugs, count, err
	}
	golog.Info(count)
	err = query.Offset((rmf.Page - 1) * rmf.Limit).Limit(rmf.Limit).Find(&bugs).Error
	// return strings.Join(rmf.Condition, " and"), rmf.Args
	// golog.Info(bugs[0].CreateTime)
	return bugs, count, err
}
