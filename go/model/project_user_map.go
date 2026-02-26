package model

import (
	"errors"
	"itflow/cache"
	"itflow/db"
)

type ProjectUserMap struct {
	ProjectId int64 `json:"project_id" gorm:"column:project_id"`
	Uid       int64 `json:"uid" gorm:"column:uid"`
}

func (ProjectUserMap) TableName() string {
	return "project_user_map"
}

func (ug *ProjectUserMap) GetUserGroupIds() ([]int64, error) {
	ids := make([]int64, 0)
	if ug.Uid <= 0 {
		return make([]int64, 0), nil
	}
	query := db.Gorm.Table(ug.TableName()).Group("project_id").Select("project_id")
	if ug.Uid != cache.SUPERID {
		query = query.Where("uid=?", ug.Uid)
	}
	err := query.Find(&ids).Error
	return ids, err
}

func (p *ProjectUserMap) GetProjectIdsByUid(uid int64) ([]int64, error) {
	pids := make([]int64, 0)
	err := db.Gorm.Table(p.TableName()).Where("uid").Select("project_id").Find(&pids).Error
	return pids, err
}

func (p *ProjectUserMap) UpdateUsersByProjectId(uids []int64) error {
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
	pm := make([]ProjectUserMap, 0, len(uids))
	for _, uid := range uids {
		pm = append(pm, ProjectUserMap{
			ProjectId: p.ProjectId,
			Uid:       uid,
		})
	}
	return db.Gorm.Table(p.TableName()).Create(&pm).Error
}

func (p *ProjectUserMap) GetUidsByProjectId() ([]int64, error) {
	uids := make([]int64, 0)
	if p.ProjectId <= 0 {
		return uids, errors.New("pid not found")
	}

	err := db.Gorm.Table(p.TableName()).Where("project_id=?", p.ProjectId).Select("uid").Find(&uids).Error
	return uids, err
}

func (p *ProjectUserMap) DeleteByProjectId() error {
	if p.ProjectId <= 0 {
		return errors.New("pid not found")
	}

	return db.Gorm.Table(p.TableName()).Where("project_id=?", p.ProjectId).Delete(p).Error
}

func (p *ProjectUserMap) InsertMany(uids []int64) error {
	inserts := make([]ProjectUserMap, 0)
	for _, uid := range uids {
		inserts = append(inserts, ProjectUserMap{
			ProjectId: p.ProjectId,
			Uid:       uid,
		})
	}

	return db.Gorm.Table(p.TableName()).Create(&inserts).Error
}
