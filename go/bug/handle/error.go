package handle

import (
	"encoding/json"
	"fmt"
	"itflow/bug/model"
)

type errorstruct struct {
	Id         int64                     `json:"id,omitempty"`
	AffectId   int64                     `json:"affectid,omitempty"`
	Code       int                       `json:"code"`
	Msg        string                    `json:"message,omitempty"`
	Path       *onefd                    `json:"path,omitempty"`
	Filename   string                    `json:"filename,omitempty"`
	Data       []byte                    `json:"data,omitempty"`
	UpdateTime int64                     `json:"updatetime,omitempty"`
	Size       int64                     `json:"size,omitempty"`
	HeaderList []*model.Table_headerlist `json:"headerlist,omitempty"`
}

func (es *errorstruct) ErrorE(err error) []byte {
	es.Code = 1
	es.Msg = err.Error()
	send, _ := json.Marshal(es)
	return send
}

func (es *errorstruct) Errorf(format string, args ...interface{}) []byte {
	es.Code = 1
	es.Msg = fmt.Sprintf(format, args...)
	send, _ := json.Marshal(es)
	return send
}

func (es *errorstruct) Error(msg string) []byte {
	es.Code = 1
	es.Msg = msg
	send, _ := json.Marshal(es)
	return send
}

func (es *errorstruct) Success() []byte {
	es.Msg = "success"
	send, _ := json.Marshal(es)
	return send
}

func (es *errorstruct) ErrorNoPermission() []byte {
	es.Msg = "没有权限"
	send, _ := json.Marshal(es)
	return send
}

//
//// can not connect mysql
//func (es *errorstruct) ErrorConnentMysql() []byte {
//	es.Code = 1
//	es.Msg = "连接数据库出错"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//// connect redis error
//func (es *errorstruct) ErrorConnentRedis() []byte {
//	es.Code = 2
//	es.Msg = "连接redis出错"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//// token not found or token expirasion
//func (es *errorstruct) ErrorNotFoundToken() []byte {
//	es.Code = 400
//	es.Msg = "not found token"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//// get post or get request data error
//func (es *errorstruct) ErrorGetData() []byte {
//	es.Code = 7
//	es.Msg = "not get response date"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//// marshaJson error
//func (es *errorstruct) ErrorMarshalJson() []byte {
//	es.Code = 5
//	es.Msg = "解析json出错"
//	send, _ := json.Marshal(es)
//	return send
//}
//
////  params error
//func (es *errorstruct) ErrorParams() []byte {
//	es.Code = 4
//	es.Msg = "参数错误"
//	send, _ := json.Marshal(es)
//	return send
//}
//
////
//func (es *errorstruct) ErrorImage() []byte {
//	es.Code = 6
//	es.Msg = "图片出错"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//// decrypt error, or not found rsa key
//func (es *errorstruct) ErrorRsa() []byte {
//	es.Code = 8
//	es.Msg = "秘钥失败"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//// password or username error
//func (es *errorstruct) ErrorUserNameOrPassword() []byte {
//	es.Code = 10
//	es.Msg = "账号或密码错误"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//// file not found
//func (es *errorstruct) ErrorFileNotFound() []byte {
//	es.Code = 12
//	es.Msg = "没有找到文件"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *errorstruct) ErrorCannotExec() []byte {
//	es.Code = 11
//	es.Msg = "无法执行"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//// has bug can not be remove
//func (es *errorstruct) ErrorHasBug() []byte {
//	es.Code = 20
//	send, _ := json.Marshal(es)
//	return send
//}
//
//// has bug can not be remove
//func (es *errorstruct) ErrorHasPosition() []byte {
//	es.Code = 21
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *errorstruct) ErrorHasHeader() []byte {
//	es.Code = 22
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *errorstruct) ErrorHasGroup() []byte {
//	es.Code = 23
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *errorstruct) ErrorHasHypo() []byte {
//	es.Code = 24
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *errorstruct) ErrorHasEnv() []byte {
//	es.Code = 25
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *errorstruct) ErrorHasUser() []byte {
//	es.Code = 26
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *errorstruct) ErrorIsDefault() []byte {
//	es.Code = 40
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *errorstruct) ErrorNoPermission() []byte {
//	es.Code = 14
//	es.Msg = "没有权限"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *errorstruct) ErrorOpenFile() []byte {
//	es.Code = 13
//	es.Msg = "打开文件失败"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *errorstruct) ErrorKeyNotFound() []byte {
//	es.Code = 30
//	es.Msg = "key not found "
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *errorstruct) ErrorType() []byte {
//	es.Code = 31
//	es.Msg = "类型错误"
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *errorstruct) ErrorSearch() []byte {
//	es.Code = 40
//	send, _ := json.Marshal(es)
//	return send
//}
//
//// 唯一性验证从100开始
//
//func (es *errorstruct) ErrorRepeatNickName() []byte {
//	es.Code = 100
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *errorstruct) ErrorRepeatEmail() []byte {
//	es.Code = 101
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *errorstruct) ErrorNull() []byte {
//	es.Code = 200
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *errorstruct) ErrorDefaultValue() []byte {
//	es.Code = 201
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *errorstruct) ErrorInUser() []byte {
//	es.Code = 202
//	send, _ := json.Marshal(es)
//	return send
//}
//
//func (es *errorstruct) ErrorNeedDefault() []byte {
//	es.Code = 203
//	send, _ := json.Marshal(es)
//	return send
//}
