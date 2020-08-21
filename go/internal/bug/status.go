package bug

import (
	"encoding/json"
	"itflow/cache"
)

type ReqStatus struct {
	Id   cache.StatusId `json:"id" type:"int" need:"是" default:"0" information:"无效"`
	Name cache.Status   `json:"name" type:"string" need:"是" default:"" information:"状态的名称"`
}

type ListStatus struct {
	StatusList cache.StatusList `json:"statuslist,omitempty" type:"array" need:"否" information:"状态列表"`
	Code       int              `json:"code" type:"int" need:"是" information:"状态码"`
}

// 添加状态管理返回的错误
type ResponeStatus struct {
	Id   int64  `json:"id" type:"int" need:"是" default:"0" information:"返回插入的id"`
	Code int    `json:"code" type:"int" need:"是" default:"" information:"状态码"`
	Msg  string `json:"msg" type:"string" need:"否" default:"" information:"错误信息"`
}

func (rs *ResponeStatus) Success() []byte {

	send, _ := json.Marshal(rs)
	return send
}

type ChangeStatus struct {
	Id     int          `json:"id"`
	Status cache.Status `json:"status"`
	Code   int          `json:"code"`
}
