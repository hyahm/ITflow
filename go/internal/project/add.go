package project

import (
	"errors"
	"itflow/cache"
	"itflow/internal/response"
	"itflow/model"

	"github.com/hyahm/golog"
)

type ReqProject struct {
	Id          cache.ProjectId `json:"id"`
	ProjectName cache.Project   `json:"projectname"`
	GroupName   string          `json:"groupname"`
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
	// 处理数据， 返回所需, 如果没加自己， 就默认把自己也加上去
	resp := &response.Response{}
	err := rp.checkValid()
	if err != nil {
		golog.Error(err)
		return resp.ErrorE(err), err
	}
	golog.Infof("%+v", rp)
	ug, ok := cache.CacheUserGroupUGid[rp.GroupName]
	if !ok {
		return resp.ErrorE(ParticipantIsEmpty), ParticipantIsEmpty
	}
	project := &model.Project{
		Name: rp.ProjectName,
		Gid:  ug.Ugid,
		Uid:  userid,
	}
	err = project.Insert()
	if err != nil {
		golog.Error(err)
		return resp.ErrorE(err), err
	}
	resp.Id = project.Id.ToInt64()
	cache.CacheProjectPid[rp.ProjectName] = cache.ProjectId(resp.Id)
	cache.CachePidProject[cache.ProjectId(resp.Id)] = rp.ProjectName
	return resp.Success(), nil
}
