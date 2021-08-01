package project

import (
	"errors"
	"itflow/model"
	"itflow/response"

	"github.com/hyahm/golog"
)

type ReqProject struct {
	Id          int64  `json:"id"`
	ProjectName string `json:"projectname"`
	GroupName   string `json:"groupname"`
}

var ProjectNameIsEmpty = errors.New("Project name is empty")
var ParticipantIsEmpty = errors.New("Participant is empty")

func (rp *ReqProject) checkValid() error {
	if rp.ProjectName == "" {
		return ProjectNameIsEmpty
	}
	if rp.GroupName == "" {
		return ParticipantIsEmpty
	}
	return nil
}

func (rp *ReqProject) Add(userid int64) ([]byte, error) {
	// 处理数据， 返回所需
	resp := &response.Response{}
	err := rp.checkValid()
	if err != nil {
		golog.Error(err)
		return resp.ErrorE(err), err
	}
	golog.Infof("%+v", rp)

	project := &model.Project{
		Name: rp.ProjectName,
		Uid:  userid,
	}
	err = project.Insert(rp.GroupName)
	if err != nil {
		golog.Error(err)
		return resp.ErrorE(err), err
	}
	resp.ID = project.Id
	return resp.Success(), nil
}
