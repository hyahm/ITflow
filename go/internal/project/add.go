package project

import (
	"errors"
	"itflow/cache"
	"itflow/internal/response"
	"itflow/model"
	"strconv"
	"strings"

	"github.com/hyahm/golog"
)

type ReqProject struct {
	Id          cache.ProjectId `json:"id"`
	ProjectName cache.Project   `json:"projectname"`
	SelectUser  []string        `json:"selectuser"`
}

var ProjectNameIsEmpty = errors.New("Project name is empty")
var ParticipantIsEmpty = errors.New("Participant is empty")

func (rp *ReqProject) checkValid() error {
	if rp.ProjectName == "" {
		return ProjectNameIsEmpty
	}
	if len(rp.SelectUser) == 0 {
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
	uids := make([]string, 0)
	needAdd := true
	for _, v := range rp.SelectUser {
		// 获取的真实用户名
		if uid, ok := cache.CacheRealNameUid[v]; ok {
			if uid == userid {
				needAdd = false
			}
			uids = append(uids, strconv.FormatInt(uid, 10))
		}
	}
	if needAdd {
		uids = append(uids, strconv.FormatInt(userid, 10))
	}
	project := &model.Project{
		Name:        rp.ProjectName,
		Participant: strings.Join(uids, ","),
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
