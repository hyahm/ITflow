package response

import (
	"encoding/json"
	"fmt"
)

type Onefd struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	ModDate  int64  `json:"date"`
	IsFile   bool   `json:"isfile"`
	IsOwner  bool   `json:"isowner"`
	HasWrite bool   `json:"haswrite"`
	Ru       bool   `json:"readuser"`
	Rname    string `json:"readname"`
	Wu       bool   `json:"writeuser"`
	Wname    string `json:"writename"`
}

type Response struct {
	Id         int64  `json:"id,omitempty" type:"int" need:"否" default:"" information:"返回的id，某些端口需要"`
	AffectId   int64  `json:"affectid,omitempty" type:"int" need:"否" default:"" information:"返回的行数，某些端口需要"`
	Code       int    `json:"code" type:"int" need:"是" default:"" information:"返回的状态码， 0为成功，非0失败"`
	Msg        string `json:"msg,omitempty" type:"string" need:"否" default:"" information:"错误信息， 状态码非肯定有"`
	Path       *Onefd `json:"path,omitempty" type:"object" need:"否" default:"" information:"返回路径，共享文件接口用到"`
	Filename   string `json:"filename,omitempty" type:"string" need:"否" default:"" information:"文件名，共享文件接口用到"`
	Data       []byte `json:"data,omitempty" type:"bytes" need:"否" default:"" information:"返回数据，某些接口用到"`
	UpdateTime int64  `json:"updatetime,omitempty" type:"string" need:"否" default:"" information:"更新时间， 共享文件接口用到"`
	Size       int64  `json:"size,omitempty" type:"int" need:"否" default:"" information:"文件大小，共享文件接口用到"`
}

func (es *Response) ErrorE(err error) []byte {
	es.Msg = err.Error()
	es.Code = 1
	send, _ := json.Marshal(es)
	return send
}

func (es *Response) GetDataErr() []byte {
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
