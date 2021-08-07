package setting

import (
	"itflow/handle"
	"itflow/internal/bug"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

// Status 状态路由组
var Status *xmux.GroupRoute

func init() {
	Status = xmux.NewGroupRoute().
		ApiCreateGroup("bugstatus", "bug 状态管理", "bug状态").
		ApiReqHeader("X-Token", "asdfasdfasdfasdfsdf").AddModule(midware.StatusPermModule)
	Status.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")
	Status.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Status.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Status.ApiCodeField("code").ApiCodeMsg("", "其他错误,请查看返回的msg")

	Status.Post("/status/list", handle.StatusList).
		ApiResStruct(&bug.ListStatus{}).ApiDescribe("获取bugstatus状态列表").
		ApiResponseTemplate(`{
			"statuslist": [
				{
					"id": 3,
					"name": "ToDoList"
				},
				{
					"id": 6,
					"name": "测试"
				},
				{
					"id": 7,
					"name": "need13"
				},
				{
					"id": 5,
					"name": "react"
				}
			],
			"code": 0
		}`)

	Status.Post("/status/add", handle.StatusAdd).Bind(&bug.ReqStatus{}).
		AddModule(midware.JsonToStruct).
		ApiDescribe("添加bug 状态").
		ApiReqStruct(&bug.ReqStatus{}).ApiRequestTemplate(`{"id": 0, "name": "普通"}`).
		ApiResStruct(&bug.ResponeStatus{}).ApiResponseTemplate(`{"id": 8,"code": 0}`)

	Status.Get("/status/remove", handle.StatusRemove).
		ApiDescribe("删除bug 状态").ApiSupplement("当此状态有bug在使用时， 无法删除")

	Status.Post("/status/update", handle.StatusUpdate).Bind(&bug.ReqStatus{}).AddModule(midware.JsonToStruct).
		ApiDescribe("修改状态").
		ApiReqStruct(&bug.ReqStatus{}).ApiRequestTemplate(`{"id": 0, "name": "普通"}`).
		ApiResStruct(&bug.ResponeStatus{}).ApiResponseTemplate(`{"code": 0}`)

}
