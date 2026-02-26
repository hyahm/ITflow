package model

import "itflow/db"

type PagePerm struct {
	Id   int64  `json:"id,omitempty" gorm:"primaryKey"`
	Name string `json:"name" gorm:"column:name"`
	Info string `json:"info" gorm:"column:info"`
}

func (PagePerm) TableName() string {
	return "page_perm"
}

func (p PagePerm) List() ([]PagePerm, error) {
	var pps []PagePerm
	err := db.Gorm.Table(p.TableName()).Find(&pps).Error
	if err != nil {
		return pps, err
	}

	return pps, err
}
