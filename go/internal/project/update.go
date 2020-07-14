package project

import (
	"itflow/cache"
	"itflow/internal/response"
	"itflow/model"
	"strconv"

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
	uids := make([]string, 0)
	needAdd := true

	if needAdd {
		uids = append(uids, strconv.FormatInt(userid, 10))
	}
	project := &model.Project{
		Id:   rp.Id,
		Name: rp.ProjectName,
		Gid:  cache.CacheUserGroupUGid[rp.GroupName].Ugid,
	}
	golog.Info("update")
	err = project.Update()
	if err != nil {
		golog.Error(err)
		return resp.ErrorE(err), err
	}
	resp.Id = project.Id.ToInt64()
	cache.CachePidProject[cache.ProjectId(resp.Id)] = rp.ProjectName
	return resp.Success(), nil
}
