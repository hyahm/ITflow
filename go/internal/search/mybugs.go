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
	Condition   []string
	Args        []interface{}
}

func (rmf *ReqMyBugFilter) GetUsefulCondition(uid int64) (string, []interface{}) {
	// 获取需要拼接的sql条件语句, 参数
	// 确定前面已经有了where 或者on
	if rmf.LevelId > 0 {
		// 判断这个值是否存在
		rmf.Condition = append(rmf.Condition, " lid=? ")
		rmf.Args = append(rmf.Args, rmf.LevelId)

	}
	if rmf.Title != "" {
		rmf.Condition = append(rmf.Condition, " title like ? ")
		rmf.Args = append(rmf.Args, "%"+rmf.Title+"%")
	}

	if rmf.ProjectId > 0 {
		// 判断这个值是否存在
		rmf.Condition = append(rmf.Condition, " pid=? ")
		rmf.Args = append(rmf.Args, rmf.ProjectId)
	}
	if len(rmf.ShowsStatus) > 0 {
		ss := make([]string, 0)
		for _, v := range rmf.ShowsStatus {
			ss = append(ss, fmt.Sprintf("%d", v))
		}
		rmf.Condition = append(rmf.Condition, fmt.Sprintf(" sid in (%s)", strings.Join(ss, ",")))
	}

	return strings.Join(rmf.Condition, " and"), rmf.Args
}
