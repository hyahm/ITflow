package handle

import (
	"encoding/json"
	"fmt"
	"itflow/cache"
	"itflow/db"
	"itflow/encrypt"
	"itflow/internal/datalog"
	"itflow/internal/response"
	"itflow/internal/role"
	"itflow/internal/user"
	"itflow/mail"
	"itflow/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	nickname := xmux.GetData(r).Get("nickname").(string)
	uid := xmux.GetData(r).Get("uid").(int64)
	getuser := xmux.GetData(r).Data.(*user.GetAddUser)
	// 验证组和职位不能为空
	if getuser.StatusGroup == "" || getuser.RoleGroup == "" || getuser.Position == "" {
		w.Write(errorcode.Error("验证组和职位不能为空"))
		return
	}
	//1，先要验证nickname 是否有重复的
	if _, ok := cache.CacheNickNameUid[getuser.Nickname]; ok {
		w.Write(errorcode.Error("nickname 重复"))
		return
	}

	//验证邮箱 是否有重复的
	var hasemail bool
	for _, v := range cache.CacheUidEmail {
		if v == getuser.Email {
			hasemail = true
		}
	}
	if hasemail {
		w.Write(errorcode.Error("email 重复"))
		return
	}

	ids := make([]string, 0)
	for k := range cache.CacheSidStatus {
		ids = append(ids, strconv.FormatInt(k.ToInt64(), 10))
	}

	var sgid int64
	var hassggroup bool
	var hasrolegroup bool
	for k, v := range cache.CacheSgidGroup {
		if v == getuser.StatusGroup {
			sgid = k
			hassggroup = true
			break
		}
	}

	var rid int64
	for k, v := range cache.CacheRidGroup {
		if v == getuser.RoleGroup {
			rid = k
			hasrolegroup = true
			break
		}
	}

	if !hasrolegroup || !hassggroup {
		w.Write(errorcode.Error("没有找到权限"))
		return
	}
	// 获取级别,如果这个职位不存在，就返回错误
	var jid int64
	var ok bool
	if jid, ok = cache.CacheJobnameJid[getuser.Position]; !ok {
		w.Write(errorcode.Error("职位不存在"))
		return
	}

	// 增加用户
	enpassword := encrypt.PwdEncrypt(getuser.Password, cache.Salt)
	user := model.User{
		NickName:   getuser.Nickname,
		RealName:   getuser.RealName,
		Password:   enpassword,
		Email:      getuser.Email,
		CreateId:   uid,
		ShowStatus: cache.StoreLevelId(strings.Join(ids, ",")),
		BugGroupId: sgid,
		Roleid:     rid,
		Jobid:      jid,
	}
	err := user.Create()
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	//更新缓存

	cache.CacheUidSgid[errorcode.Id] = sgid
	cache.CacheUidNickName[errorcode.Id] = getuser.Nickname
	cache.CacheUidRealName[errorcode.Id] = getuser.RealName
	cache.CacheNickNameUid[getuser.Nickname] = errorcode.Id
	cache.CacheRealNameUid[getuser.RealName] = errorcode.Id
	cache.CacheUidRid[errorcode.Id] = rid
	cache.CacheUidRid[errorcode.Id] = jid
	cache.CacheUidEmail[cache.CacheNickNameUid[nickname]] = getuser.Email

	// 邮件通知

	if cache.CacheEmail.CreateUser {
		content := fmt.Sprintf("你的用户名: %v;<br> 密码: %v", getuser.Email, getuser.Password)
		mail.SendMail("创建用户成功", content, getuser.Email)
	}

	// 更新日志
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "setting",
		Action:   "createuser",
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func RemoveUser(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	id := r.FormValue("id")
	id32, err := strconv.Atoi(id)
	if err != nil {
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 判断是否有bug
	var count int
	err = db.Mconn.GetOne("select count(id) from bugs where uid=?", id).Scan(&count)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if count > 0 {
		golog.Error("uid:%v,has bugs,can not remove")
		w.Write(errorcode.Error("has bugs,can not remove"))
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
	_, err = db.Mconn.Update("delete from user where id=?", id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	nickname := xmux.GetData(r).Get("nickname").(string)
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "setting",
		Action:   "deluser",
	}

	delete(cache.CacheNickNameUid, cache.CacheUidNickName[int64(id32)])
	delete(cache.CacheRealNameUid, cache.CacheUidRealName[int64(id32)])
	delete(cache.CacheUidEmail, int64(id32))
	delete(cache.CacheUidRealName, int64(id32))
	delete(cache.CacheUidNickName, int64(id32))
	delete(cache.CacheUidFilter, int64(id32))
	delete(cache.CacheUidSgid, int64(id32))
	delete(cache.CacheUidRid, int64(id32))
	delete(cache.CacheUidJid, int64(id32))
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func DisableUser(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	nickname := xmux.GetData(r).Get("nickname").(string)

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

	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "setting",
		Action:   "disableuser",
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

// 显示自己能管理的权限，不显示自己的
func UserList(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	uid := xmux.GetData(r).Get("uid").(int64)
	if cache.SUPERID != uid {
		w.Write(errorcode.ErrorNoPermission())
		return
	}

	uls := &user.UserList{}

	// 0是系统管理员， 1是管理层， 2是普通用户
	//switch level {
	//case 0:
	getallsql := "select id,createtime,realname,nickname,email,disable,rid,bugsid,jid from user where level<>0"
	adminrows, err := db.Mconn.GetRows(getallsql)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	for adminrows.Next() {
		ul := &user.User{}
		var rid int64
		var jid int64
		var bugsid int64
		adminrows.Scan(&ul.Id, &ul.Createtime, &ul.Realname, &ul.Nickname, &ul.Email, &ul.Disable, &rid, &bugsid, &jid)
		ul.StatusGroup = cache.CacheSgidGroup[bugsid]
		ul.RoleGroup = cache.CacheRidGroup[rid]
		ul.Position = cache.CacheJidJobname[jid]
		uls.Userlist = append(uls.Userlist, ul)
	}

	send, _ := json.Marshal(uls)
	w.Write(send)
	return

}

func UserUpdate(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	nickname := xmux.GetData(r).Get("nickname").(string)
	if cache.SUPERID != cache.CacheNickNameUid[nickname] {
		w.Write(errorcode.ErrorNoPermission())
		return
	}

	uls := xmux.GetData(r).Data.(*user.User)

	// 0是系统管理员， 1是管理层， 2是普通用户
	//switch level {
	//case 0:
	var hasrolegroup bool
	var hasstatusgroup bool
	var rid int64
	var bsid int64
	for k, v := range cache.CacheRidGroup {
		if v == uls.RoleGroup {
			rid = k
			hasrolegroup = true
			break
		}
	}

	for k, v := range cache.CacheRidGroup {
		if v == uls.RoleGroup {
			rid = k
			hasrolegroup = true
			break
		}
	}

	for k, v := range cache.CacheSgidGroup {
		if v == uls.StatusGroup {
			bsid = k
			hasstatusgroup = true
			break
		}
	}
	if _, ok := cache.CacheJobnameJid[uls.Position]; !ok {
		w.Write(errorcode.Error("没有找到职位"))
		return
	}
	if !hasrolegroup || !hasstatusgroup {
		w.Write(errorcode.Error("没有找到status"))
		return
	}

	getallsql := "update user set realname=?,nickname=?,email=?,rid=?,bugsid=?,jid=? where id=?"
	_, err := db.Mconn.Update(getallsql,
		uls.Realname, uls.Nickname, uls.Email, rid, bsid, cache.CacheJobnameJid[uls.Position],
		uls.Id,
	)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "setting",
		Action:   "updateuser",
	}

	//更新缓存
	delete(cache.CacheNickNameUid, cache.CacheUidNickName[int64(uls.Id)])
	delete(cache.CacheRealNameUid, cache.CacheUidNickName[int64(uls.Id)])
	cache.CacheRealNameUid[uls.Realname] = int64(uls.Id)
	cache.CacheUidSgid[int64(uls.Id)] = bsid
	cache.CacheUidNickName[int64(uls.Id)] = uls.Nickname
	cache.CacheUidRealName[int64(uls.Id)] = uls.Realname
	cache.CacheNickNameUid[uls.Nickname] = int64(uls.Id)
	cache.CacheRealNameUid[uls.Realname] = int64(uls.Id)
	cache.CacheUidRid[int64(uls.Id)] = rid
	cache.CacheUidRid[int64(uls.Id)] = cache.CacheJobnameJid[uls.Position]
	cache.CacheUidEmail[cache.CacheNickNameUid[nickname]] = uls.Email
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	errorcode := &response.Response{}

	getuser := xmux.GetData(r).Data.(*user.ChangePasswod)

	nickname := xmux.GetData(r).Get("nickname").(string)
	uid := cache.CacheNickNameUid[nickname]

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
