package bug

import (
	"itflow/cache"
)

type PassBug struct {
	Id          int64         `json:"id"  type:"string" need:"是" default:"" information:"bug的id"`
	Date        int64         `json:"date"  type:"string" need:"是" default:"" information:"提交事件"`
	Remark      string        `json:"remark"  type:"string" need:"是" default:"" information:"说明"`
	SelectUsers []string      `json:"selectusers"  type:"string" need:"是" default:"" information:"转交的任务"`
	Status      cache.Status  `json:"status"  type:"string" need:"是" default:"" information:"转交后的状态"`
	User        string        `json:"user"  type:"string" need:"是" default:"" information:"处理人"`
	ProjectName cache.Project `json:"projectname"  type:"string" need:"是" default:"" information:"项目名"`
	Code        int           `json:"code"  type:"int" need:"是" default:"" information:"错误码"`
}
