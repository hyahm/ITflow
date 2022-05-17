package handle

import (
	"fmt"
	"itflow/cache"
	"itflow/db"
	"itflow/encrypt"
	"itflow/internal/user"
	"itflow/model"
	"itflow/response"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func Create(w http.ResponseWriter, r *http.Request) {
	uid := xmux.GetInstance(r).Get("uid").(int64)

	getuser := xmux.GetInstance(r).Data.(*model.User)
	if getuser.Jobid == 0 {
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = "职位不能为空"
		return
	}
	if strings.Contains(getuser.NickName, "@") {
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = "昵称不能包含@符号"
		return
	}
	getuser.Password = encrypt.PwdEncrypt(getuser.Password, cache.Salt)
	getuser.CreateTime = time.Now().Unix()
	getuser.CreateUId = uid
	err := getuser.Create()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}

	model.CacheEmail.SendMail("成功创建用户",
		fmt.Sprintf(`<html><body><h1>已成功创建用户<h1>登录网址:<a href="%s">%s</a></br>用户名: %s</br> 密码: %s</br>邮箱: %s</body></html>`,
			r.Referer(), r.Referer(), getuser.NickName, getuser.Password, getuser.Email),
		getuser.Email)

	xmux.GetInstance(r).Response.(*response.Response).ID = getuser.ID

}

func Delete(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = "id类型错误"
		return
	}
	if id64 == cache.SUPERID {
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = "无法删除"
		return
	}
	err = model.DeleteUser(id)
	// 判断是否有bug
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}

}

func DisableUser(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")

	result := db.Mconn.Update("update user set disable=ABS(disable-1) where id=?", id)
	if result.Err != nil {
		golog.Error(result.Err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = result.Err.Error()
		return
	}
	result = db.Mconn.Update("update bugs set dustbin=ABS(dustbin-1) where uid=?", id)
	if result.Err != nil {
		golog.Error(result.Err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = result.Err.Error()
		return
	}

}

// 显示自己能管理的权限，不显示自己的
func Read(w http.ResponseWriter, r *http.Request) {
	uid := xmux.GetInstance(r).Get("uid").(int64)
	us, err := model.GetAllUsers(uid)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = us
}

func Update(w http.ResponseWriter, r *http.Request) {
	// 跟 userlist 一样，
	// 根据uid 查找
	uid := xmux.GetInstance(r).Get("uid").(int64)
	if cache.SUPERID != uid {
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = "没有权限"
		return
	}

	user := xmux.GetInstance(r).Data.(*model.User)

	err := user.Update()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
}

func GetRoles(w http.ResponseWriter, r *http.Request) {
	ar, err := model.AllRole()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = ar
}

func GetRoleGroupPerm(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	rg := model.RoleGroup{}
	data, err := rg.GetEditDataById(id)
	// ar, err := model.AllRole()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	// 通过permids 来获取详细权限
	xmux.GetInstance(r).Response.(*response.Response).Data = data
}

// func GetThisRoles(w http.ResponseWriter, r *http.Request) {

// 	errorcode := &response.Response{}

// 	rl := &getroles{}

// 	id := r.FormValue("id")

// 	var rolestring string
// 	err := db.Mconn.GetOne("select rolestring from user where id=?", id).Scan(&rolestring)
// 	if err != nil {
// 		golog.Error(err)
// 		w.Write(errorcode.ConnectMysqlFail())
// 		return
// 	}

// 	send, _ := json.Marshal(rl)
// 	w.Write(send)
// 	return

// }

func GetGroup(w http.ResponseWriter, r *http.Request) {

}

func GetTaskTyp(w http.ResponseWriter, r *http.Request) {

	ts := make(map[int]string, 0)

	rows, err := db.Mconn.GetRows("select id,name from typ")
	if err != nil {
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	for rows.Next() {
		var t string
		var id int
		err = rows.Scan(&id, &t)
		if err != nil {
			golog.Info(err)
			continue
		}
		ts[id] = t
	}
	golog.Error(ts)
	xmux.GetInstance(r).Response.(*response.Response).Data = ts

}

func ResetPwd(w http.ResponseWriter, r *http.Request) {

	rp := xmux.GetInstance(r).Data.(*user.ResetPassword)

	newpassword := encrypt.PwdEncrypt(rp.Password, cache.Salt)

	updatepwdsql := "update user set password=? where id=?"
	result := db.Mconn.Update(updatepwdsql, newpassword, rp.Id)
	if result.Err != nil {
		golog.Error(result.Err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = result.Err.Error()
		return
	}

}
