package handle

import (
	"encoding/json"
	"fmt"
	"itflow/cache"
	"itflow/db"
	"itflow/encrypt"
	"itflow/internal/response"
	"itflow/internal/role"
	"itflow/internal/user"
	"net/http"
	"strings"
	"time"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	// nickname := xmux.GetData(r).Get("nickname").(string)
	uid := xmux.GetData(r).Get("uid").(int64)

	createTime := time.Now().Unix()
	getuser := xmux.GetData(r).Data.(*user.GetAddUser)
	if strings.Contains(getuser.Nickname, "@") {
		w.Write(errorcode.Error("昵称不能包含@符号"))
		return
	}
	enpassword := encrypt.PwdEncrypt(getuser.Password, cache.Salt)
	var err error
	db.Mconn.OpenDebug()
	errorcode.Id, err = db.Mconn.Insert(`insert into user(nickname, password, email, createtime, createuid, realname, jid) values(
		?,?,?,?,?,?, 
		(select id from jobs where name=?))`, getuser.Nickname,
		enpassword, getuser.Email, createTime,
		uid, getuser.RealName, getuser.Position)
	golog.Info(db.Mconn.GetSql())
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	cache.CacheEmail.SendMail("成功创建用户",
		fmt.Sprintf(`<html><body><h1>已成功创建用户<h1>登录网址:<a href="%s">%s</a></br>用户名: %s</br> 密码: %s</br>邮箱: %s</body></html>`, r.Referer(), r.Referer(), getuser.Nickname, getuser.Password, getuser.Email),
		getuser.Email)
	// // 验证组和职位不能为空
	// if getuser.StatusGroup == "" || getuser.RoleGroup == "" || getuser.Position == "" {
	// 	w.Write(errorcode.Error("验证组和职位不能为空"))
	// 	return
	// }
	// //1，先要验证nickname 是否有重复的
	// if _, ok := cache.CacheNickNameUid[getuser.Nickname]; ok {
	// 	w.Write(errorcode.Error("nickname 重复"))
	// 	return
	// }

	// //验证邮箱 是否有重复的
	// var hasemail bool
	// for _, v := range cache.CacheUidEmail {
	// 	if v == getuser.Email {
	// 		hasemail = true
	// 	}
	// }
	// if hasemail {
	// 	w.Write(errorcode.Error("email 重复"))
	// 	return
	// }

	// ids := make([]string, 0)
	// for k := range cache.CacheSidStatus {
	// 	ids = append(ids, strconv.FormatInt(k.ToInt64(), 10))
	// }

	// var sgid int64
	// var hassggroup bool

	// for k, v := range cache.CacheSgidGroup {
	// 	if v == getuser.StatusGroup {
	// 		sgid = k
	// 		hassggroup = true
	// 		break
	// 	}
	// }

	// var rid int64
	// err := model.CheckRoleNameInGroup(getuser.RoleGroup, &rid)
	// if err != nil {
	// 	golog.Error(err)
	// 	w.Write(errorcode.ErrorE(err))
	// 	return
	// }
	// if !hassggroup {
	// 	w.Write(errorcode.Error("没有找到权限"))
	// 	return
	// }
	// // 获取级别,如果这个职位不存在，就返回错误
	// var jid int64
	// var ok bool
	// if jid, ok = cache.CacheJobnameJid[getuser.Position]; !ok {
	// 	w.Write(errorcode.Error("职位不存在"))
	// 	return
	// }

	// // 增加用户

	// user := model.User{
	// 	NickName:   getuser.Nickname,
	// 	RealName:   getuser.RealName,
	// 	Password:   enpassword,
	// 	Email:      getuser.Email,
	// 	CreateId:   uid,
	// 	ShowStatus: cache.StoreLevelId(strings.Join(ids, ",")),
	// 	BugGroupId: sgid,
	// 	Roleid:     rid,
	// 	Jobid:      jid,
	// }
	// err = user.Create()
	// if err != nil {
	// 	golog.Error(err)
	// 	w.Write(errorcode.ErrorE(err))
	// 	return
	// }
	// //更新缓存

	// send, _ := json.Marshal(errorcode)
	w.Write(errorcode.Success())
	return

}

func RemoveUser(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	id := r.FormValue("id")
	// 判断是否有bug
	var count int
	err := db.Mconn.GetOne("select count(id) from bugs where uid=?", id).Scan(&count)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if count > 0 {
		golog.Error("uid:%v,has bugs,can not remove")
		w.Write(errorcode.IsUse())
		return
	}
	// 查看用户组是否存在此用户
	userrows, err := db.Mconn.GetRows("select ids from usergroup")
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	var hasgroup bool
	for userrows.Next() {
		var ids string
		userrows.Scan(&ids)
		for _, v := range strings.Split(ids, ",") {
			if v == id {
				hasgroup = true
				break
			}
		}
		if hasgroup {
			w.Write(errorcode.Error("还有group"))
			return
		}
	}
	userrows.Close()
	_, err = db.Mconn.Update("delete from user where id=?", id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

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
	return

}

// 显示自己能管理的权限，不显示自己的
func UserList(w http.ResponseWriter, r *http.Request) {
	uid := xmux.GetData(r).Get("uid").(int64)
	errorcode := &response.Response{}
	uls := &user.UserList{}
	if uid == cache.SUPERID {
		getallsql := `select u.id,createtime,realname,nickname,email,disable,j.name from 
		user as u 
		join jobs as j 
		on u.jid = j.id and u.id<>?`
		adminrows, err := db.Mconn.GetRows(getallsql, cache.SUPERID)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
		for adminrows.Next() {
			ul := &user.User{}
			err = adminrows.Scan(&ul.Id, &ul.Createtime, &ul.Realname, &ul.Nickname, &ul.Email,
				&ul.Disable, &ul.Position)
			if err != nil {
				golog.Info(err)
				continue
			}
			uls.Userlist = append(uls.Userlist, ul)
		}
		adminrows.Close()
		send, _ := json.Marshal(uls)
		w.Write(send)
		return
	} else {
		getallsql := `select u.id,createtime,realname,nickname,email,disable,r.name,s.name,j.name from 
		user as u  join rolegroup as r 
		join statusgroup as s 
		join jobs as j 
		on u.rid = r.id and u.bugsid = s.id and u.jid = j.id and u.jid in (select id from jobs where hypo=(select jid from user where id=?))`
		adminrows, err := db.Mconn.GetRows(getallsql, uid)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
		for adminrows.Next() {
			ul := &user.User{}
			err = adminrows.Scan(&ul.Id, &ul.Createtime, &ul.Realname, &ul.Nickname, &ul.Email,
				&ul.Disable, &ul.RoleGroup, &ul.StatusGroup, &ul.Position)
			if err != nil {
				golog.Info(err)
				continue
			}
			uls.Userlist = append(uls.Userlist, ul)
		}
		adminrows.Close()
		send, _ := json.Marshal(uls)
		w.Write(send)
		return
	}

}

func UserUpdate(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	uid := xmux.GetData(r).Get("uid").(int64)
	if cache.SUPERID != uid {
		w.Write(errorcode.ErrorNoPermission())
		return
	}

	uls := xmux.GetData(r).Data.(*user.User)

	// 0是系统管理员， 1是管理层， 2是普通用户
	//switch level {
	//case 0:

	// var hasstatusgroup bool
	// var rid int64
	// var bsid int64

	// err := model.CheckRoleNameInGroup(uls.RoleGroup, &rid)
	// if err != nil {
	// 	golog.Error(err)
	// 	w.Write(errorcode.ErrorE(err))
	// 	return
	// }

	// for k, v := range cache.CacheSgidGroup {
	// 	if v == uls.StatusGroup {
	// 		bsid = k
	// 		hasstatusgroup = true
	// 		break
	// 	}
	// }
	// if _, ok := cache.CacheJobnameJid[uls.Position]; !ok {
	// 	w.Write(errorcode.Error("没有找到职位"))
	// 	return
	// }
	// if !hasstatusgroup {
	// 	w.Write(errorcode.Error("没有找到status"))
	// 	return
	// }
	if strings.Contains(uls.Nickname, "@") {
		w.Write(errorcode.Error("昵称不能包含@符号"))
		return
	}
	getallsql := `update user set 
	 realname=?,	nickname=?,	email=?,
	 jid=(select coalesce(min(id),0) from jobs where name=?) 
	 where id=?`
	_, err := db.Mconn.Update(getallsql,
		uls.Realname, uls.Nickname, uls.Email, uls.Position,
		uls.Id,
	)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	errorcode := &response.Response{}

	getuser := xmux.GetData(r).Data.(*user.ChangePasswod)

	uid := xmux.GetData(r).Get("uid").(int64)

	getaritclesql := "select count(id) from user where id=? and password=?"
	oldpassword := encrypt.PwdEncrypt(getuser.Oldpassword, cache.Salt)
	var n int
	err := db.Mconn.GetOne(getaritclesql, uid, oldpassword).Scan(&n)
	if err != nil || n != 1 {
		golog.Error(err)
		w.Write(errorcode.ErrorNoPermission())
		return
	}

	newpassword := encrypt.PwdEncrypt(getuser.Newpassword, cache.Salt)
	chpwdsql := "update user set password=? where id=?"

	_, err = db.Mconn.Update(chpwdsql, newpassword, uid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return
}

func GetRoles(w http.ResponseWriter, r *http.Request) {

	rl := &role.RespRoles{}

	w.Write(rl.List())
	return

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

func ResetPwd(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	rp := xmux.GetData(r).Data.(*user.ResetPassword)

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
