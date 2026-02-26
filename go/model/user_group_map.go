package model

import (
	"errors"
	"itflow/cache"
	"itflow/db"
)

type UserGroupMap struct {
	ProjectId int64 `json:"project_id" gorm:"column:project_id"`
	Uid       int64 `json:"uid" gorm:"column:uid"`
}

func (UserGroupMap) TableName() string {
	return "user_group_map"
}

func (ug *UserGroupMap) GetUserGroupIds() ([]int64, error) {

	ids := make([]int64, 0)
	if ug.Uid <= 0 {
		return make([]int64, 0), nil
	}
	query := db.Gorm.Table(ug.TableName()).Group("project_id desc").Select("project_id")
	if ug.Uid != cache.SUPERID {
		query = query.Where("uid=?", ug.Uid)
	}
	err := query.Find(&ids).Error
	return ids, err
}

func (p *UserGroupMap) GetProjectIdsByUid(uid int64) ([]int64, error) {
	pids := make([]int64, 0)
	err := db.Gorm.Table(p.TableName()).Where("uid").Select("project_id").Scan(&pids).Error
	return pids, err
}

func (p *UserGroupMap) UpdateUsersByProjectId(uids []int64) error {
	if p.ProjectId <= 0 {
		return errors.New("pid not found")
	}
	err := p.DeleteByProjectId()
	if err != nil {
		return err
	}
	if len(uids) == 0 {
		return nil
	}
	pm := make([]UserGroupMap, 0, len(uids))
	for _, uid := range uids {
		pm = append(pm, UserGroupMap{
			ProjectId: p.ProjectId,
			Uid:       uid,
		})
	}
	return db.Gorm.Table(p.TableName()).Create(&pm).Error
}

func (p *UserGroupMap) GetUidsByProjectId() ([]int64, error) {
	uids := make([]int64, 0)
	if p.ProjectId <= 0 {
		return uids, errors.New("pid not found")
	}

	err := db.Gorm.Table(p.TableName()).Where("project_id=?", p.ProjectId).Select("uid").Find(&uids).Error
	return uids, err
}

func (p *UserGroupMap) DeleteByProjectId() error {
	if p.ProjectId <= 0 {
		return errors.New("pid not found")
	}

	return db.Gorm.Table(p.TableName()).Where("project_id=?", p.ProjectId).Delete(p).Error
}

func (p *UserGroupMap) InsertMany(uids []int64) error {
	inserts := make([]UserGroupMap, 0)
	for _, uid := range uids {
		inserts = append(inserts, UserGroupMap{
			ProjectId: p.ProjectId,
			Uid:       uid,
		})
	}

	return db.Gorm.Table(p.TableName()).Create(&inserts).Error
}
