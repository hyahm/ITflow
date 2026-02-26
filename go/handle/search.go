package handle

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func SearchAllBugs(w http.ResponseWriter, r *http.Request) {

	uid := xmux.GetInstance(r).Get("uid").(int64)
	search := xmux.GetInstance(r).Data.(*model.ReqMyBugFilter)
	bug := model.Bug{}
	bugs, count, err := bug.GetUsefulCondition(*search, uid)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).Count = count
	xmux.GetInstance(r).Response.(*response.Response).Data = bugs
}

// func SearchMyBugs(w http.ResponseWriter, r *http.Request) {

// 	uid := xmux.GetInstance(r).Get("uid").(int64)
// 	search := xmux.GetInstance(r).Data.(*model.ReqMyBugFilter)
// 	bug := model.Bug{}

// 	bugs, count, err := bug.GetUsefulCondition(*search, uid)
// 	if err != nil {
// 		golog.Error(err)
// 		xmux.GetInstance(r).Response.(*response.Response).Code = 1
// 		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
// 		return
// 	}
// 	xmux.GetInstance(r).Response.(*response.Response).Count = count
// 	xmux.GetInstance(r).Response.(*response.Response).Data = bugs
// }
