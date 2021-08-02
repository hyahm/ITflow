package response

import (
	"encoding/json"
	"fmt"

	"github.com/hyahm/golog"
)

type Response struct {
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	ID         int64       `json:"id"`
	UpdateTime int64       `json:"update_time"`
	Data       interface{} `json:"data,omitemtpy"`
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
