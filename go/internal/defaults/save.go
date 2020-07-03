package defaults

import (
	"errors"
	"itflow/cache"
	"itflow/model"
)

type ReqDefaultValue struct {
	DefaultStatus cache.Status `json:"defaultstatus"`
}

func (qdv *ReqDefaultValue) Save() error {
	sid := qdv.DefaultStatus.Id()
	if sid == 0 {
		return errors.New("没有找到status ")
	}

	dv := &model.DefaultValue{
		Status: qdv.DefaultStatus.Id(),
	}
	err := dv.Update()
	if err != nil {
		return err
	}
	// 	// 更新缓存
	cache.DefaultSid = sid
	return nil
}
