package midware

// 专门赋予管理层的用户权限
// func UserPerm(w http.ResponseWriter, r *http.Request) bool {
// 	errorcode := &response.Response{}
// 	uid := xmux.GetInstance(r).Get("uid").(int64)
// 	if uid == cache.SUPERID {
// 		return false
// 	}
// 	var manager_count int

// 	err := db.Mconn.GetOne("select count(id) from jobs where hypo=(select jid from user where id=?)", uid).Scan(&manager_count)
// 	if err == nil && manager_count > 0 {
// 		return false
// 	}
// 	golog.Error(err)
// 	errorcode.Msg = "没有权限"
// 	w.Write(errorcode.Success())
// 	return true
// }
