package handle

import (
	"itflow/bug/model"
	"encoding/json"
)

type errorstruct struct {
	Id         int64                     `json:"id"`
	AffectId   int64                     `json:"affectid"`
	Code       int                       `json:"statuscode"`
	Msg        string                    `json:"message"`
	Path       *onefd                    `json:"path"`
	Filename   string                    `json:"filename"`
	Data       []byte                    `json:"data"`
	UpdateTime int64                     `json:"updatetime"`
	Size       int64                     `json:"size"`
	HeaderList []*model.Table_headerlist `json:"headerlist"`
}

// can not connect mysql
func (es *errorstruct) ErrorConnentMysql() []byte {
	es.Code = 1
	send, _ := json.Marshal(es)
	return send
}

// connect redis error
func (es *errorstruct) ErrorConnentRedis() []byte {
	es.Code = 2
	send, _ := json.Marshal(es)
	return send
}

// token not found or token expirasion
func (es *errorstruct) ErrorNotFoundToken() []byte {
	es.Code = 400
	send, _ := json.Marshal(es)
	return send
}

// get post or get request data error
func (es *errorstruct) ErrorGetData() []byte {
	es.Code = 7
	send, _ := json.Marshal(es)
	return send
}

// marshaJson error
func (es *errorstruct) ErrorMarshalJson() []byte {
	es.Code = 5
	send, _ := json.Marshal(es)
	return send
}

//  params error
func (es *errorstruct) ErrorParams() []byte {
	es.Code = 4
	send, _ := json.Marshal(es)
	return send
}

//
func (es *errorstruct) ErrorImage() []byte {
	es.Code = 6
	send, _ := json.Marshal(es)
	return send
}

// decrypt error, or not found rsa key
func (es *errorstruct) ErrorRsa() []byte {
	es.Code = 8
	send, _ := json.Marshal(es)
	return send
}

// password or username error
func (es *errorstruct) ErrorUserNameOrPassword() []byte {
	es.Code = 10
	send, _ := json.Marshal(es)
	return send
}

// file not found
func (es *errorstruct) ErrorFileNotFound() []byte {
	es.Code = 12
	send, _ := json.Marshal(es)
	return send
}

func (es *errorstruct) ErrorCannotExec() []byte {
	es.Code = 11
	send, _ := json.Marshal(es)
	return send
}

// has bug can not be remove
func (es *errorstruct) ErrorHasBug() []byte {
	es.Code = 20
	send, _ := json.Marshal(es)
	return send
}

// has bug can not be remove
func (es *errorstruct) ErrorHasPosition() []byte {
	es.Code = 21
	send, _ := json.Marshal(es)
	return send
}

func (es *errorstruct) ErrorHasHeader() []byte {
	es.Code = 22
	send, _ := json.Marshal(es)
	return send
}

func (es *errorstruct) ErrorHasGroup() []byte {
	es.Code = 23
	send, _ := json.Marshal(es)
	return send
}

func (es *errorstruct) ErrorHasHypo() []byte {
	es.Code = 24
	send, _ := json.Marshal(es)
	return send
}

func (es *errorstruct) ErrorHasEnv() []byte {
	es.Code = 25
	send, _ := json.Marshal(es)
	return send
}

func (es *errorstruct) ErrorHasUser() []byte {
	es.Code = 26
	send, _ := json.Marshal(es)
	return send
}

func (es *errorstruct) ErrorIsDefault() []byte {
	es.Code = 40
	send, _ := json.Marshal(es)
	return send
}

func (es *errorstruct) ErrorNoPermission() []byte {
	es.Code = 14
	send, _ := json.Marshal(es)
	return send
}

func (es *errorstruct) ErrorOpenFile() []byte {
	es.Code = 13
	send, _ := json.Marshal(es)
	return send
}

func (es *errorstruct) ErrorKeyNotFound() []byte {
	es.Code = 30
	send, _ := json.Marshal(es)
	return send
}

func (es *errorstruct) ErrorType() []byte {
	es.Code = 31
	send, _ := json.Marshal(es)
	return send
}

func (es *errorstruct) ErrorSearch() []byte {
	es.Code = 40
	send, _ := json.Marshal(es)
	return send
}

// 唯一性验证从100开始

func (es *errorstruct) ErrorRepeatNickName() []byte {
	es.Code = 100
	send, _ := json.Marshal(es)
	return send
}

func (es *errorstruct) ErrorRepeatEmail() []byte {
	es.Code = 101
	send, _ := json.Marshal(es)
	return send
}

//func (es *errorstruct) ErrorOpenFile() []byte {
//	es.Code = 13
//	send, _ := json.Marshal(es)
//	return send
//}

func (es *errorstruct) ErrorNull() []byte {
	es.Code = 200
	send, _ := json.Marshal(es)
	return send
}

func (es *errorstruct) ErrorDefaultValue() []byte {
	es.Code = 201
	send, _ := json.Marshal(es)
	return send
}

func (es *errorstruct) ErrorInUser() []byte {
	es.Code = 202
	send, _ := json.Marshal(es)
	return send
}

func (es *errorstruct) ErrorNeedDefault() []byte {
	es.Code = 203
	send, _ := json.Marshal(es)
	return send
}
