package search

import (
	"errors"
	"itflow/cache"
)

var ErrorNoStatus = errors.New("没选择状态，返回空数组")

type ReqMyBugFilter struct {
	Page        int           `json:"page"`
	Limit       int           `json:"limit"`
	Level       cache.Level   `json:"level"`
	Project     cache.Project `json:"project"`
	Title       string        `json:"title"`
	ShowsStatus []string      `json:"showstatus"`
}

func (rmf *ReqMyBugFilter) GetUsefulCondition(uid int64) (string, []interface{}) {
	// 获取需要拼接的sql条件语句, 参数
	// 确定前面已经有了where 或者on
	condition := ""
	args := make([]interface{}, 0)
	if rmf.Level != "" {
		// 判断这个值是否存在
		condition += " and lid=(select id from level where name=?) "
		args = append(args, rmf.Level)

	}
	if rmf.Title != "" {
		condition += " and title like ? "
		args = append(args, "%"+rmf.Title+"%")
	}

	if rmf.Project != "" {
		// 判断这个值是否存在
		condition += " and pid=(select id from project where name=?) "
		args = append(args, rmf.Project)
	}
	// 获取此用户能看到的状态

	return condition, args
}
