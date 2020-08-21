package status

import (
	"encoding/json"
	"itflow/model"
)

type RespGetNames struct {
	StatusList []string `json:"statuslist"`
	Code       int      `json:"code"`
	Msg        string   `json:"msg"`
}

func (rgn *RespGetNames) Marshal() []byte {
	if rgn == nil {
		return nil
	}
	send, _ := json.Marshal(rgn)
	return send
}

func GetNames() []byte {
	resp := &RespGetNames{}
	st := &model.Status{}
	var err error
	resp.StatusList, err = st.Names()
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = 1
		return resp.Marshal()
	}
	return resp.Marshal()
}
