package project

import (
	"itflow/internal/response"
	"itflow/model"

	"github.com/hyahm/golog"
)

func (rp *ReqProject) Update(userid int64) ([]byte, error) {
	// 处理数据， 返回所需
	resp := &response.Response{}
	err := rp.checkValid()
	if err != nil {
		golog.Error(err)
		return resp.ErrorE(err), err
	}

	project := &model.Project{
		Id:   rp.Id,
		Name: rp.ProjectName,
	}

	golog.Info("update")
	err = project.Update(rp.GroupName)
	if err != nil {
		golog.Error(err)
		return resp.ErrorE(err), err
	}
	return resp.Success(), nil
}
