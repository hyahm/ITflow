package project

import (
	"itflow/db"
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

type ProjectListResponse struct {
	model.Project
	UserInfo []model.KeyName `json:"user_info"`
}

func Read(w http.ResponseWriter, r *http.Request) {
	// 拿到所有项目

	uid := xmux.GetInstance(r).Get("uid").(int64)
	project := model.Project{
		Uid: uid,
	}
	projects, err := project.GetAllProjects(uid)

	// projects, err := model.GetAllProjects(uid)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	golog.Info(projects)
	plrs := make([]ProjectListResponse, 0, len(projects))
	for _, v := range projects {
		var kvs []model.KeyName
		pum := model.ProjectUserMap{
			ProjectId: v.Id,
		}
		uids, err := pum.GetUidsByProjectId()
		if err != nil {
			golog.Error(err)
			xmux.GetInstance(r).Response.(*response.Response).Code = 1
			xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
			return
		}
		if len(uids) > 0 {
			user := model.User{}
			kvs, err = user.GetKeyNameByUids(uids)
			if err != nil {
				golog.Error(err)
				xmux.GetInstance(r).Response.(*response.Response).Code = 1
				xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
				return
			}
		}

		plr := ProjectListResponse{
			v, kvs,
		}
		plrs = append(plrs, plr)
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = plrs

}

type ProjectRequest struct {
	model.Project
	UIds []int64 `json:"uids"`
}

func Create(w http.ResponseWriter, r *http.Request) {

	pr := xmux.GetInstance(r).Data.(*ProjectRequest)
	uid := xmux.GetInstance(r).Get("uid").(int64)
	project := model.Project{
		Uid:  uid,
		Name: pr.Name,
	}
	err := project.Insert()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	// 更新用户
	pgm := model.ProjectUserMap{
		ProjectId: project.Id,
	}
	err = pgm.InsertMany(pr.UIds)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).ID = project.Id
}

func Update(w http.ResponseWriter, r *http.Request) {
	project := xmux.GetInstance(r).Data.(*ProjectRequest)
	uid := xmux.GetInstance(r).Get("uid").(int64)
	err := project.Update(uid)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	pum := model.ProjectUserMap{
		ProjectId: project.Id,
	}
	err = pum.UpdateUsersByProjectId(project.UIds)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
}

func ProjectKeys(w http.ResponseWriter, r *http.Request) {
	uid := xmux.GetInstance(r).Get("uid").(int64)
	pkn, err := model.GetProjectKeyName(uid)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = pkn

}

func Delete(w http.ResponseWriter, r *http.Request) {

	project := xmux.GetInstance(r).Data.(*model.Project)
	golog.Info(project.Id)
	// 判断有没有bug在使用这个
	var count int64

	err := db.Gorm.Table("bugs").Where("pid=?", project.Id).Count(&count).Error
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}

	if count > 0 {
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = "使用中，无法删除"
		return
	}

	err = db.Gorm.Table("project").Where("id=?", project.Id).Delete(project).Error
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}

}
