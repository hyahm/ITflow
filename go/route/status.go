package route

import (
	"itflow/app/handle"
	"itflow/midware"
	"itflow/model/bug"

	"github.com/hyahm/xmux"
)

var Status *xmux.GroupRoute

func init() {
	Status = xmux.NewGroupRoute("status").
		ApiCreateGroup("bugstatus", "bug 状态管理", "bug status").
		ApiReqHeader("X-Token", "asdfasdfasdfasdfsdf")

	Status.Pattern("/status/list").Post(handle.StatusList).
		ApiResStruct(&bug.ListStatus{}).ApiDescribe("获取bugstatus状态列表").
		ApiCodeMsg("3", "没有权限").
		ApiCodeMsg("10", "token 过期").ApiResponseTemplate(`{
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
		}`).
		ApiCodeMsg("0", "成功")

	Status.Pattern("/status/add").Post(handle.StatusAdd).Bind(&bug.Status{}).
		AddMidware(midware.CheckPermssion).End(midware.EndLog).AddMidware(midware.JsonToStruct)

	Status.Pattern("/status/remove").Get(handle.StatusRemove).
		End(midware.EndLog)

	Status.Pattern("/status/update").Post(handle.StatusUpdate).
		End(midware.EndLog)

	Status.Pattern("/status/groupname").Post(handle.StatusGroupName)
}
