package defaults

import (
	"errors"
	"itflow/cache"
	"itflow/model"
)

type ReqDefaultValue struct {
	Created   cache.Status `json:"created"`
	Completed cache.Status `json:"completed"`
}

func (qdv *ReqDefaultValue) Save() error {
	sid := qdv.Created.Id()
	if sid == 0 {
		return errors.New("没有找到status ")
	}

	cid := qdv.Completed.Id()
	if cid == 0 {
		return errors.New("没有找到completestatus ")
	}

	dv := &model.DefaultValue{
		Created:   qdv.Created.Id(),
		Completed: qdv.Completed.Id(),
	}
	err := dv.Update()
	if err != nil {
		return err
	}
	// 	// 更新缓存
	cache.DefaultCreateSid = sid
	cache.DefaultCompleteSid = sid
	return nil
}
