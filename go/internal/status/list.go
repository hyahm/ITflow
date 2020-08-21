package status

import (
	"encoding/json"
	"itflow/model"
)

type RespStatusList struct {
	StatusList []*model.Status `json:"statuslist"`
	Code       int             `json:"code"`
	Msg        string          `json:"msg,omitempty"`
}

func (rsl *RespStatusList) Marshal() []byte {
	if rsl == nil {
		return nil
	}
	send, _ := json.Marshal(rsl)
	return send
}

// 获取状态的id和名称列表
func StatusList() []byte {
	resp := &RespStatusList{}
	st := &model.Status{}
	var err error
	resp.StatusList, err = st.List()
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = 1

		return resp.Marshal()
	}
	return resp.Marshal()
}
