package search

import (
	"errors"
	"fmt"
	"strings"
)

var ErrorNoStatus = errors.New("没选择状态，返回空数组")

type ReqMyBugFilter struct {
	Page        int     `json:"page"`
	Limit       int     `json:"limit"`
	LevelId     int64   `json:"level_id"`
	ProjectId   int64   `json:"project_id"`
	Title       string  `json:"title"`
	ShowsStatus []int64 `json:"showstatus"`
}

func (rmf *ReqMyBugFilter) GetUsefulCondition(uid int64) (string, []interface{}) {
	// 获取需要拼接的sql条件语句, 参数
	// 确定前面已经有了where 或者on
	condition := make([]string, 0)
	args := make([]interface{}, 0)
	if rmf.LevelId > 0 {
		// 判断这个值是否存在
		condition = append(condition, " lid=? ")
		args = append(args, rmf.LevelId)

	}
	if rmf.Title != "" {
		condition = append(condition, " title like ? ")
		args = append(args, "%"+rmf.Title+"%")
	}

	if rmf.ProjectId > 0 {
		// 判断这个值是否存在
		condition = append(condition, " pid=? ")
		args = append(args, rmf.ProjectId)
	}
	if len(rmf.ShowsStatus) > 0 {
		ss := make([]string, 0)
		for _, v := range rmf.ShowsStatus {
			ss = append(ss, fmt.Sprintf("%d", v))
		}
		condition = append(condition, fmt.Sprintf(" sid in (%s)", strings.Join(ss, ",")))
	}
	// 获取此用户能看到的状态

	return strings.Join(condition, " and"), args
}
