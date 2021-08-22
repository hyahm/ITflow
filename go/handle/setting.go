package handle

import (
	"encoding/json"
	"fmt"
	"itflow/cache"
	"itflow/db"
	"itflow/encrypt"
	"itflow/internal/user"
	"itflow/model"
	"itflow/response"
	"net/http"
	"strings"
	"time"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	uid := xmux.GetInstance(r).Get("uid").(int64)

	getuser := xmux.GetInstance(r).Data.(*model.User)
	if strings.Contains(getuser.NickName, "@") {
		w.Write(response.Error("昵称不能包含@符号"))
		return
	}
	getuser.Password = encrypt.PwdEncrypt(getuser.Password, cache.Salt)
	getuser.CreateTime = time.Now().Unix()
	getuser.CreateUId = uid
	err := getuser.Create()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	cache.CacheEmail.SendMail("成功创建用户",
		fmt.Sprintf(`<html><body><h1>已成功创建用户<h1>登录网址:<a href="%s">%s</a></br>用户名: %s</br> 密码: %s</br>邮箱: %s</body></html>`,
			r.Referer(), r.Referer(), getuser.NickName, getuser.Password, getuser.Email),
		getuser.Email)

	res := response.Response{
		ID: getuser.ID,
	}
	w.Write(res.Marshal())
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	err := model.DeleteUser(id)
	// 判断是否有bug
	// var count int
	// err := db.Mconn.GetOne("select count(id) from bugs where uid=?", id).Scan(&count)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	w.Write(response.Success())
	// if count > 0 {
	// 	golog.Error("uid:%v,has bugs,can not remove")
	// 	w.Write(errorcode.IsUse())
	// 	return
	// }
	// // 查看用户组是否存在此用户
	// userrows, err := db.Mconn.GetRows("select ids from usergroup")
	// if err != nil {
	// 	golog.Error(err)
	// 	w.Write(errorcode.ErrorE(err))
	// 	return
	// }
	// var hasgroup bool
	// for userrows.Next() {
	// 	var ids string
	// 	userrows.Scan(&ids)
	// 	for _, v := range strings.Split(ids, ",") {
	// 		if v == id {
	// 			hasgroup = true
	// 			break
	// 		}
	// 	}
	// 	if hasgroup {
	// 		w.Write(errorcode.Error("还有group"))
	// 		return
	// 	}
	// }
	// userrows.Close()
	// _, err = db.Mconn.Update("delete from user where id=?", id)
	// if err != nil {
	// 	golog.Error(err)
	// 	w.Write(errorcode.ErrorE(err))
	// 	return
	// }

	// send, _ := json.Marshal(errorcode)
	// w.Write(send)
	// return

}

func DisableUser(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	id := r.FormValue("id")
	var err error
	_, err = db.Mconn.Update("update user set disable=ABS(disable-1) where id=?", id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	_, err = db.Mconn.Update("update bugs set dustbin=ABS(dustbin-1) where uid=?", id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)

}

// 显示自己能管理的权限，不显示自己的
func UserList(w http.ResponseWriter, r *http.Request) {
	uid := xmux.GetInstance(r).Get("uid").(int64)
	us, err := model.GetAllUsers(uid)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := response.Response{
		Data: us,
	}
	w.Write(res.Marshal())
}

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	// 跟 userlist 一样，
	// 根据uid 查找
	errorcode := &response.Response{}
	uid := xmux.GetInstance(r).Get("uid").(int64)
	if cache.SUPERID != uid {
		w.Write(errorcode.ErrorNoPermission())
		return
	}

	user := xmux.GetInstance(r).Data.(*model.User)

	err := user.Update()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	w.Write(response.Success())
}

func GetRoles(w http.ResponseWriter, r *http.Request) {
	res := response.Response{}
	ar, err := model.AllRole()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res.Data = ar
	w.Write(res.Marshal())
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

type sendGroup struct {
	Groups []string `json:"groups"`
	Code   int      `json:"code"`
}

func GetGroup(w http.ResponseWriter, r *http.Request) {

	sg := &sendGroup{}

	send, _ := json.Marshal(sg)
	w.Write(send)
	return

}

type sty struct {
	Ts   map[int]string `json:"ts"`
	Code int            `json:"code"`
}

func GetTaskTyp(w http.ResponseWriter, r *http.Request) {

	ts := &sty{
		Ts: make(map[int]string, 0),
	}
	rows, err := db.Mconn.GetRows("select id,name from typ")
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"code": 2, "msg": "%s"}`, err.Error())))
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
		ts.Ts[id] = t
	}
	send, _ := json.Marshal(ts)
	w.Write(send)
	return

}

func ResetPwd(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	rp := xmux.GetInstance(r).Data.(*user.ResetPassword)

	newpassword := encrypt.PwdEncrypt(rp.Password, cache.Salt)

	updatepwdsql := "update user set password=? where id=?"
	_, err := db.Mconn.Update(updatepwdsql, newpassword, rp.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}
