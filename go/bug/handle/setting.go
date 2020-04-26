package handle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"itflow/bug/bugconfig"
	"itflow/bug/buglog"
	"itflow/bug/mail"
	"itflow/db"
	"itflow/gaencrypt"
	"itflow/model/response"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hyahm/golog"
)

type getAddUser struct {
	Nickname    string   `json:"nickname"`
	Email       string   `json:"email"`
	Password    string   `json:"password"`
	Role        []string `json:"role"`
	RealName    string   `json:"realname"`
	RoleGroup   string   `json:"rolegroup"`
	StatusGroup string   `json:"statusgroup"`
	Position    string   `json:"position"` // 普通用户就是真，管理员就假
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &response.Response{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	getuser := &getAddUser{}

	gu, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = json.Unmarshal(gu, getuser)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	//1，先要验证nickname 是否有重复的
	if _, ok := bugconfig.CacheNickNameUid[getuser.Nickname]; ok {
		w.Write(errorcode.Error("nickname 重复"))
		return
	}
	var hasemail bool
	for _, v := range bugconfig.CacheUidEmail {
		if v == getuser.Email {
			hasemail = true
		}
	}
	if hasemail {
		w.Write(errorcode.Error("email 重复"))
		return
	}
	// 验证组和职位不能为空
	if getuser.StatusGroup == "" || getuser.RoleGroup == "" || getuser.Position == "" {
		w.Write(errorcode.Error("验证组和职位不能为空"))
		return
	}
	ids := make([]string, 0)
	for k, _ := range bugconfig.CacheSidStatus {
		ids = append(ids, strconv.FormatInt(k, 10))
	}
	var sgid int64
	var hassggroup bool
	var hasrolegroup bool
	for k, v := range bugconfig.CacheSgidGroup {
		if v == getuser.StatusGroup {
			sgid = k
			hassggroup = true
			break
		}
	}

	var rid int64
	for k, v := range bugconfig.CacheRidGroup {
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
	if jid, ok = bugconfig.CacheJobnameJid[getuser.Position]; !ok {
		w.Write(errorcode.Error("职位不存在"))
		return
	}
	var level int64
	row, err := db.Mconn.GetOne("select level from jobs where id=?", jid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = row.Scan(&level)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 增加用户
	showstatus := strings.Join(ids, ",")
	enpassword := gaencrypt.PwdEncrypt(getuser.Password, bugconfig.Salt)
	createusersql := "insert into user(nickname,password,email,headimg,createtime,createuid,realname,showstatus,disable,bugsid,level,rid,jid) values(?,?,?,?,?,?,?,?,?,?,?,?,?)"
	errorcode.Id, err = db.Mconn.Insert(createusersql,
		getuser.Nickname, enpassword, getuser.Email,
		"", time.Now().Unix(), bugconfig.CacheNickNameUid[nickname],
		getuser.RealName, showstatus, false,
		sgid, level, rid, jid,
	)

	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	//更新缓存

	bugconfig.CacheUidSgid[errorcode.Id] = sgid
	bugconfig.CacheUidNickName[errorcode.Id] = getuser.Nickname
	bugconfig.CacheUidRealName[errorcode.Id] = getuser.RealName
	bugconfig.CacheNickNameUid[getuser.Nickname] = errorcode.Id
	bugconfig.CacheRealNameUid[getuser.RealName] = errorcode.Id
	bugconfig.CacheUidRid[errorcode.Id] = rid
	bugconfig.CacheUidRid[errorcode.Id] = jid
	bugconfig.CacheUidEmail[bugconfig.CacheNickNameUid[nickname]] = getuser.Email

	// 邮件通知

	if bugconfig.CacheEmail.CreateUser {
		content := fmt.Sprintf("你的用户名: %v;<br> 密码: %v", getuser.Email, getuser.Password)
		mail.SendMail("创建用户成功", content, []string{getuser.Email})
	}
	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "user",
	}
	err = il.Add(
		getuser.RealName, getuser.Nickname, getuser.Email)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func RemoveUser(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &response.Response{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	id := r.FormValue("id")
	id32, err := strconv.Atoi(id)
	if err != nil {
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 判断是否有bug
	var count int
	row, err := db.Mconn.GetOne("select count(id) from bugs where uid=?", id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = row.Scan(&count)
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

	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "user",
	}
	err = il.Del(
		nickname, id, bugconfig.CacheUidRealName[int64(id32)])
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	delete(bugconfig.CacheNickNameUid, bugconfig.CacheUidNickName[int64(id32)])
	delete(bugconfig.CacheRealNameUid, bugconfig.CacheUidRealName[int64(id32)])
	delete(bugconfig.CacheUidEmail, int64(id32))
	delete(bugconfig.CacheUidRealName, int64(id32))
	delete(bugconfig.CacheUidNickName, int64(id32))
	delete(bugconfig.CacheUidFilter, int64(id32))
	delete(bugconfig.CacheUidSgid, int64(id32))
	delete(bugconfig.CacheUidRid, int64(id32))
	delete(bugconfig.CacheUidJid, int64(id32))
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func DisableUser(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &response.Response{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	id := r.FormValue("id")

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

	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "user",
	}
	err = il.Del(
		nickname, id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

type userlist struct {
	Id          int    `json:"id"`
	Createtime  int64  `json:"createtime"`
	Realname    string `json:"realname"`
	Nickname    string `json:"nickname"`
	Email       string `json:"email"`
	Disable     int    `json:"disable"`
	StatusGroup string `json:"statusgroup"`
	RoleGroup   string `json:"rolegroup"`
	Position    string `json:"position"`
}

type sendUserList struct {
	Userlist []*userlist `json:"userlist"`
	Code     int         `json:"code"`
}

// 显示自己能管理的权限，不显示自己的
func UserList(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &response.Response{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if bugconfig.SUPERID != bugconfig.CacheNickNameUid[nickname] {
		w.Write(errorcode.ErrorNoPermission())
		return
	}

	uls := &sendUserList{}

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
		ul := &userlist{}
		var rid int64
		var jid int64
		var bugsid int64
		adminrows.Scan(&ul.Id, &ul.Createtime, &ul.Realname, &ul.Nickname, &ul.Email, &ul.Disable, &rid, &bugsid, &jid)
		ul.StatusGroup = bugconfig.CacheSgidGroup[bugsid]
		ul.RoleGroup = bugconfig.CacheRidGroup[rid]
		ul.Position = bugconfig.CacheJidJobname[jid]
		uls.Userlist = append(uls.Userlist, ul)
	}
	//case 1:
	//	getusersql := "select id,createtime,realname,nickname,rolestring,email,disable,rid,bugsid from user where level=1 and nickname<>?"
	//	adminrows, err := conn.GetRows(getusersql, nickname)
	//	if err != nil {
	//		golog.Error(err)
	//		w.Write(errorcode.ErrorConnentMysql())
	//		return
	//	}
	//	for adminrows.Next() {
	//		ul := &userlist{}
	//		var rid int64
	//		var bugsid int64
	//		adminrows.Scan(&ul.Id, &ul.Createtime, &ul.Realname, &ul.Nickname, &ul.Role, &ul.Email, &ul.Disable, &rid, &bugsid)
	//		ul.Role = bugconfig.CacheRidRole[rid]
	//		ul.BugStatusGroup = bugconfig.CacheSgidGroup[bugsid]
	//		uls.Userlist = append(uls.Userlist, ul)
	//	}
	//default:
	//	ul := &userlist{}
	//	var rid int64
	//	getusersql := "select id,createtime,realname,nickname,rolestring,email,disable from user where level=2 and nickname=?"
	//	err := conn.GetOne(getusersql, nickname).Scan(&ul.Id, &ul.Createtime, &ul.Realname, &ul.Nickname, &rid, &ul.Email, &ul.Disable)
	//	if err != nil {
	//		golog.Error(err)
	//		w.Write(errorcode.ErrorConnentMysql())
	//		return
	//	}
	//	ul.Role = bugconfig.CacheRidRole[rid]
	//	uls.Userlist = append(uls.Userlist, ul)
	//}

	send, _ := json.Marshal(uls)
	w.Write(send)
	return

}

func UserUpdate(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &response.Response{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if bugconfig.SUPERID != bugconfig.CacheNickNameUid[nickname] {
		w.Write(errorcode.ErrorNoPermission())
		return
	}

	uls := &userlist{}
	bytedata, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	err = json.Unmarshal(bytedata, uls)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 0是系统管理员， 1是管理层， 2是普通用户
	//switch level {
	//case 0:
	var hasrolegroup bool
	var hasstatusgroup bool
	var rid int64
	var bsid int64
	for k, v := range bugconfig.CacheRidGroup {
		if v == uls.RoleGroup {
			rid = k
			hasrolegroup = true
			break
		}
	}

	for k, v := range bugconfig.CacheRidGroup {
		if v == uls.RoleGroup {
			rid = k
			hasrolegroup = true
			break
		}
	}

	for k, v := range bugconfig.CacheSgidGroup {
		if v == uls.StatusGroup {
			bsid = k
			hasstatusgroup = true
			break
		}
	}
	if _, ok := bugconfig.CacheJobnameJid[uls.Position]; !ok {
		w.Write(errorcode.Error("没有找到职位"))
		return
	}
	if !hasrolegroup || !hasstatusgroup {
		w.Write(errorcode.Error("没有找到status"))
		return
	}

	getallsql := "update user set realname=?,nickname=?,email=?,rid=?,bugsid=?,jid=? where id=?"
	_, err = db.Mconn.Update(getallsql,
		uls.Realname, uls.Nickname, uls.Email, rid, bsid, bugconfig.CacheJobnameJid[uls.Position],
		uls.Id,
	)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "user",
	}
	err = il.Update("updateuser : changeuser:%s, operator: %s  ",
		uls.Realname, nickname)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	//更新缓存
	delete(bugconfig.CacheNickNameUid, bugconfig.CacheUidNickName[int64(uls.Id)])
	delete(bugconfig.CacheRealNameUid, bugconfig.CacheUidNickName[int64(uls.Id)])
	bugconfig.CacheRealNameUid[uls.Realname] = int64(uls.Id)
	bugconfig.CacheUidSgid[int64(uls.Id)] = bsid
	bugconfig.CacheUidNickName[int64(uls.Id)] = uls.Nickname
	bugconfig.CacheUidRealName[int64(uls.Id)] = uls.Realname
	bugconfig.CacheNickNameUid[uls.Nickname] = int64(uls.Id)
	bugconfig.CacheRealNameUid[uls.Realname] = int64(uls.Id)
	bugconfig.CacheUidRid[int64(uls.Id)] = rid
	bugconfig.CacheUidRid[int64(uls.Id)] = bugconfig.CacheJobnameJid[uls.Position]
	bugconfig.CacheUidEmail[bugconfig.CacheNickNameUid[nickname]] = uls.Email
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

type pwd struct {
	Oldpassword string `json:"oldpassword"`
	Newpassword string `json:"newpassword"`
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	errorcode := &response.Response{}
	name, err := logtokenmysql(r)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	getuser := &pwd{}
	gu, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	fmt.Println(string(gu))
	uid := bugconfig.CacheNickNameUid[name]

	err = json.Unmarshal(gu, getuser)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	getaritclesql := "select count(id) from user where id=? and password=?"
	oldpassword := gaencrypt.PwdEncrypt(getuser.Oldpassword, bugconfig.Salt)
	var n int
	row, err := db.Mconn.GetOne(getaritclesql, uid, oldpassword)
	if err != nil || n != 1 {
		golog.Error(err)
		w.Write(errorcode.ErrorNoPermission())
		return
	}
	err = row.Scan(&n)
	if err != nil || n != 1 {
		golog.Error(err)
		w.Write(errorcode.ErrorNoPermission())
		return
	}
	newpassword := gaencrypt.PwdEncrypt(getuser.Newpassword, bugconfig.Salt)
	chpwdsql := "update user set password=? where id=?"

	_, err = db.Mconn.Update(chpwdsql, newpassword, uid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = insertlog("resetpassword", "用户"+name+"修改了密码", r)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return
	return

}

type getroles struct {
	Rolelist []string `json:"rolelist"`
	Code     int      `json:"code"`
}

func GetRoles(w http.ResponseWriter, r *http.Request) {

	_, err := logtokenmysql(r)
	errorcode := &response.Response{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	rl := &getroles{}
	for _, v := range bugconfig.CacheRidRole {
		rl.Rolelist = append(rl.Rolelist, v)
	}
	send, _ := json.Marshal(rl)
	w.Write(send)
	return

}

func GetThisRoles(w http.ResponseWriter, r *http.Request) {

	_, err := logtokenmysql(r)
	errorcode := &response.Response{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	rl := &getroles{}

	id := r.FormValue("id")

	var rolestring string
	row, err := db.Mconn.GetOne("select rolestring from user where id=?", id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = row.Scan(&rolestring)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(rl)
	w.Write(send)
	return

}

type sendGroup struct {
	Groups []string `json:"groups"`
	Code   int      `json:"code"`
}

func GetGroup(w http.ResponseWriter, r *http.Request) {

	_, err := logtokenmysql(r)
	errorcode := &response.Response{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	sg := &sendGroup{}

	send, _ := json.Marshal(sg)
	w.Write(send)
	return

}

type resetPassword struct {
	Id       int    `json:"id"`
	Password string `json:"newpassword"`
}

func ResetPwd(w http.ResponseWriter, r *http.Request) {

	_, err := logtokenmysql(r)
	errorcode := &response.Response{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	rp := &resetPassword{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = json.Unmarshal(body, rp)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	newpassword := gaencrypt.PwdEncrypt(rp.Password, bugconfig.Salt)

	updatepwdsql := "update user set password=? where id=?"
	_, err = db.Mconn.Update(updatepwdsql, newpassword, rp.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}
