package response

import (
	"encoding/json"
	"fmt"

	"github.com/hyahm/golog"
)

type Response struct {
	Code       int         `json:"code"`
	Msg        string      `json:"msg,omitempty"`
	ID         int64       `json:"id,omitempty"`
	UpdateTime int64       `json:"update_time,omitempty"`
	CreateTime int64       `json:"create_time,omitempty"`
	UserIds    []int64     `json:"user_ids,omitempty"`
	VersionIds []int64     `json:"version_ids,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Count      int         `json:"count,omitempty"`
	Page       int         `json:"page,omitempty"`
	IsAdmin    bool        `json:"is_admin,omitempty"`
}

func (r *Response) Marshal() []byte {
	send, err := json.Marshal(r)
	if err != nil {
		golog.Error(err)
	}
	return send
}

func (es *Response) ErrorE(err error) []byte {
	es.Msg = err.Error()
	es.Code = 1
	send, _ := json.Marshal(es)
	return send
}

func (es *Response) GetInstanceErr() []byte {
	es.Msg = "获取请求数据失败"
	es.Code = 1
	send, _ := json.Marshal(es)
	return send
}

func (es *Response) IsUse() []byte {
	es.Msg = "使用中，无法删除"
	es.Code = 10
	send, _ := json.Marshal(es)
	return send
}

func (es *Response) NotJsonFormat() []byte {
	es.Msg = "数据无法被解析"
	es.Code = 1
	send, _ := json.Marshal(es)
	return send
}

func (es *Response) Errorf(format string, args ...interface{}) []byte {
	es.Msg = fmt.Sprintf(format, args...)
	send, _ := json.Marshal(es)
	return send
}

func (es *Response) Error(msg string) []byte {
	es.Msg = msg
	es.Code = 1
	send, _ := json.Marshal(es)
	return send
}

func (es *Response) Success() []byte {
	es.Msg = "success"
	send, _ := json.Marshal(es)
	return send
}

func (es *Response) ErrorNoPermission() []byte {
	es.Msg = "没有权限"
	es.Code = 3
	send, _ := json.Marshal(es)
	return send
}

func (es *Response) TokenNotFound() []byte {
	es.Code = -1
	es.Msg = "token过期"
	send, _ := json.Marshal(es)
	return send
}

// func (es *Response) ConnectRedisFail() []byte {
// 	es.Code = 11
// 	es.Msg = "系统错误"
// 	send, _ := json.Marshal(es)
// 	return send
// }

func (es *Response) ConnectMysqlFail() []byte {
	// 连接mysql失败
	es.Code = 1
	es.Msg = "系统错误"
	send, _ := json.Marshal(es)
	return send
}

func (es *Response) LoginFailed() []byte {
	es.Code = 2
	es.Msg = "用户或密码错误"
	send, _ := json.Marshal(es)
	return send
}

func Error(msg string) []byte {
	return []byte(fmt.Sprintf(`{"code": 1, "msg": "%s"}`, msg))
}

func ErrorE(err error) []byte {
	return Error(err.Error())
}

func Success() []byte {
	return []byte(`{"code": 0}`)
}
