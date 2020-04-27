package response

import (
	"encoding/json"
	"fmt"
	"itflow/bug/model"
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
	Id         int64                     `json:"id,omitempty" type:"int" need:"否" default:"" information:"返回的id，某些端口需要"`
	AffectId   int64                     `json:"affectid,omitempty" type:"int" need:"否" default:"" information:"返回的行数，某些端口需要"`
	Code       int                       `json:"code" type:"int" need:"是" default:"" information:"返回的状态码， 0为成功，非0失败"`
	Msg        string                    `json:"message,omitempty" type:"string" need:"否" default:"" information:"错误信息， 状态码非肯定有"`
	Path       *Onefd                    `json:"path,omitempty" type:"object" need:"否" default:"" information:"返回路径，共享文件接口用到"`
	Filename   string                    `json:"filename,omitempty" type:"string" need:"否" default:"" information:"文件名，共享文件接口用到"`
	Data       []byte                    `json:"data,omitempty" type:"bytes" need:"否" default:"" information:"返回数据，某些接口用到"`
	UpdateTime int64                     `json:"updatetime,omitempty" type:"string" need:"否" default:"" information:"更新时间， 共享文件接口用到"`
	Size       int64                     `json:"size,omitempty" type:"int" need:"否" default:"" information:"文件大小，共享文件接口用到"`
	HeaderList []*model.Table_headerlist `json:"headerlist,omitempty" type:"array" need:"否" default:"" information:"用到的时候再标识， 一下想不起来"`
}

func (es *Response) ErrorE(err error) []byte {
	es.Code = 1
	es.Msg = err.Error()
	send, _ := json.Marshal(es)
	return send
}

func (es *Response) Errorf(format string, args ...interface{}) []byte {
	es.Code = 1
	es.Msg = fmt.Sprintf(format, args...)
	send, _ := json.Marshal(es)
	return send
}

func (es *Response) Error(msg string) []byte {
	es.Code = 1
	es.Msg = msg
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
	send, _ := json.Marshal(es)
	return send
}

//
//// can not connect mysql
//func (es *response.Response) ErrorConnentMysql() []byte {
//	es.Code = 1
//	es.Msg = "连接数据库出错"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//// connect redis error
//func (es *response.Response) ErrorConnentRedis() []byte {
//	es.Code = 2
//	es.Msg = "连接redis出错"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//// token not found or token expirasion
//func (es *response.Response) ErrorNotFoundToken() []byte {
//	es.Code = 400
//	es.Msg = "not found token"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//// get post or get request data error
//func (es *response.Response) ErrorGetData() []byte {
//	es.Code = 7
//	es.Msg = "not get response date"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//// marshaJson error
//func (es *response.Response) ErrorMarshalJson() []byte {
//	es.Code = 5
//	es.Msg = "解析json出错"
//	send, _ := json.Marshal(es)
//	return send
//}
//
////  params error
//func (es *response.Response) ErrorParams() []byte {
//	es.Code = 4
//	es.Msg = "参数错误"
//	send, _ := json.Marshal(es)
//	return send
//}
//
////
//func (es *response.Response) ErrorImage() []byte {
//	es.Code = 6
//	es.Msg = "图片出错"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//// decrypt error, or not found rsa key
//func (es *response.Response) ErrorRsa() []byte {
//	es.Code = 8
//	es.Msg = "秘钥失败"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//// password or username error
//func (es *response.Response) ErrorUserNameOrPassword() []byte {
//	es.Code = 10
//	es.Msg = "账号或密码错误"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//// file not found
//func (es *response.Response) ErrorFileNotFound() []byte {
//	es.Code = 12
//	es.Msg = "没有找到文件"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *response.Response) ErrorCannotExec() []byte {
//	es.Code = 11
//	es.Msg = "无法执行"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//// has bug can not be remove
//func (es *response.Response) ErrorHasBug() []byte {
//	es.Code = 20
//	send, _ := json.Marshal(es)
//	return send
//}
//
//// has bug can not be remove
//func (es *response.Response) ErrorHasPosition() []byte {
//	es.Code = 21
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *response.Response) ErrorHasHeader() []byte {
//	es.Code = 22
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *response.Response) ErrorHasGroup() []byte {
//	es.Code = 23
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *response.Response) ErrorHasHypo() []byte {
//	es.Code = 24
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *response.Response) ErrorHasEnv() []byte {
//	es.Code = 25
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *response.Response) ErrorHasUser() []byte {
//	es.Code = 26
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *response.Response) ErrorIsDefault() []byte {
//	es.Code = 40
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *response.Response) ErrorNoPermission() []byte {
//	es.Code = 14
//	es.Msg = "没有权限"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *response.Response) ErrorOpenFile() []byte {
//	es.Code = 13
//	es.Msg = "打开文件失败"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *response.Response) ErrorKeyNotFound() []byte {
//	es.Code = 30
//	es.Msg = "key not found "
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *response.Response) ErrorType() []byte {
//	es.Code = 31
//	es.Msg = "类型错误"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *response.Response) ErrorSearch() []byte {
//	es.Code = 40
//	send, _ := json.Marshal(es)
//	return send
//}
//
//// 唯一性验证从100开始
//
//func (es *response.Response) ErrorRepeatNickName() []byte {
//	es.Code = 100
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *response.Response) ErrorRepeatEmail() []byte {
//	es.Code = 101
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *response.Response) ErrorNull() []byte {
//	es.Code = 200
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *response.Response) ErrorDefaultValue() []byte {
//	es.Code = 201
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *response.Response) ErrorInUser() []byte {
//	es.Code = 202
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *response.Response) ErrorNeedDefault() []byte {
//	es.Code = 203
//	send, _ := json.Marshal(es)
//	return send
//}
