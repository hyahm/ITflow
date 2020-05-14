package handle

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"itflow/app/bugconfig"
	"itflow/app/public"
	"itflow/db"
	"itflow/model"
	"itflow/network/bug"
	"itflow/network/response"
	"net/http"
	"strconv"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func SearchAllBugs(w http.ResponseWriter, r *http.Request) {

	countbasesql := "select count(id) from bugs where dustbin=0 and uid=? "
	bugsql := "select id,createtime,iid,sid,bugtitle,lid,pid,eid,spusers from bugs where dustbin=0 and uid=? "

	al, err := getbuglist(r, countbasesql, bugsql, false)
	if err != nil {
		w.Write(err)
		return
	}
	send, _ := json.Marshal(al)
	w.Write(send)
	return

}

func SearchMyBugs(w http.ResponseWriter, r *http.Request) {

	countbasesql := "select count(id) from bugs where dustbin=0 and uid=? "
	bugsql := "select id,createtime,iid,sid,bugtitle,lid,pid,eid,spusers from bugs where dustbin=0 and uid=? "

	al, err := getbuglist(r, countbasesql, bugsql, false)
	if err != nil {
		w.Write(err)
		return
	}
	send, _ := json.Marshal(al)
	w.Write(send)
	return

}

func SearchMyTasks(w http.ResponseWriter, r *http.Request) {
	countbasesql := "select count(id) from bugs where dustbin=0 "
	bugsql := "select id,createtime,iid,sid,bugtitle,lid,pid,eid,spusers from bugs where dustbin=0 "
	errorcode := &response.Response{}
	al := &model.AllArticleList{}
	// 获取状态

	nickname := xmux.GetData(r).Get("nickname").(string)
	showstatus := bugconfig.CacheUidFilter[bugconfig.CacheNickNameUid[nickname]]

	//更新缓存
	bugconfig.CacheUidFilter[bugconfig.CacheNickNameUid[nickname]] = showstatus

	searchparam := xmux.GetData(r).Data.(*bug.SearchParam)
	// 第二步， 检查level
	if searchparam.Level != "" {
		// 判断这个值是否存在
		if lid, ok := bugconfig.CacheLevelLid[searchparam.Level]; ok {
			bugsql += fmt.Sprintf("and lid=%d ", lid)
			countbasesql += fmt.Sprintf("and lid=%d ", lid)
		} else {
			golog.Error("没有搜索到什么")
			w.Write(errorcode.Error("没有搜索到什么"))
			return
		}
	}
	// 第三步， 检查Title
	if searchparam.Title != "" {
		bugsql += fmt.Sprintf("and bugtitle like '%s' ", searchparam.Title)
		countbasesql += fmt.Sprintf("and bugtitle like '%s' ", searchparam.Title)

	}
	// 第四步， 检查Project
	if searchparam.Project != "" {
		// 判断这个值是否存在
		if pid, ok := bugconfig.CacheProjectPid[searchparam.Project]; ok {
			bugsql += fmt.Sprintf("and pid=%d ", pid)
			countbasesql += fmt.Sprintf("and pid=%d ", pid)
		} else {
			golog.Error("没有搜索到什么")
			w.Write(errorcode.Error("没有搜索到什么"))
			return
		}
	}
	if showstatus != "" {
		bugsql += fmt.Sprintf("and sid in (%s)", showstatus)
	}
	rows, err := db.Mconn.GetRows(bugsql)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	for rows.Next() {
		one := &model.ArticleList{}
		var iid int64
		var sid int64
		var lid int64
		var pid int64
		var eid int64
		var userlist string
		rows.Scan(&one.ID, &one.Date, &iid, &sid, &one.Title, &lid, &pid, &eid, &userlist)
		// 如果不存在这么办， 添加修改的时候需要判断
		one.Importance = bugconfig.CacheIidImportant[iid]
		one.Status = bugconfig.CacheSidStatus[sid]
		one.Level = bugconfig.CacheLidLevel[lid]
		one.Projectname = bugconfig.CachePidName[pid]
		one.Env = bugconfig.CacheEidName[eid]
		// 显示realname

		// 判断是否是自己的任务
		var ismytask bool
		for _, v := range strings.Split(userlist, ",") {
			if v == strconv.FormatInt(bugconfig.CacheNickNameUid[nickname], 10) {
				ismytask = true
				break
			}
		}

		if ismytask {
			for _, v := range strings.Split(userlist, ",") {
				//判断用户是否存在，不存在就 删吗 ， 先不删
				userid32, _ := strconv.Atoi(v)
				if realname, ok := bugconfig.CacheUidRealName[int64(userid32)]; ok {
					one.Handle = append(one.Handle, realname)
				}
			}
			one.Author = bugconfig.CacheUidRealName[bugconfig.CacheNickNameUid[nickname]]
			al.Count++
			al.Al = append(al.Al, one)
		}

	}
	// 获取查询的开始位置
	start, end := public.GetPagingLimitAndPage(al.Count, searchparam.Page, searchparam.Limit)
	if len(al.Al) > end {
		al.Al = al.Al[start:end]
	}

	send, _ := json.Marshal(al)
	w.Write(send)
	return

}

func SearchBugManager(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	al := &model.AllArticleList{}
	searchparam := xmux.GetData(r).Data.(*bug.BugManager)

	basesql, args := managertotal("select count(id) from bugs", searchparam)

	err := db.Mconn.GetOne(basesql, args...).Scan(&al.Count)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if al.Count == 0 {
		w.Write(errorcode.Error("没有找到bug"))
		return
	}
	alsql := "select id,createtime,iid,sid,bugtitle,uid,lid,pid,eid,spusers,dustbin from bugs"

	rows, err := managersearch(alsql, al.Count, searchparam)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	for rows.Next() {
		bl := &model.ArticleList{}
		var sid, lid int64
		var spusers string
		var uid int64
		var pid int64
		var eid int64
		var iid int64
		rows.Scan(&bl.ID, &bl.Date, &iid, &sid, &bl.Title, &uid, &lid, &pid, &eid, &spusers, &bl.Dustbin)
		bl.Level = bugconfig.CacheLidLevel[lid]
		bl.Importance = bugconfig.CacheIidImportant[iid]
		bl.Status = bugconfig.CacheSidStatus[sid]
		bl.Author = bugconfig.CacheUidRealName[uid]
		bl.Projectname = bugconfig.CachePidName[pid]
		bl.Handle = formatUserlistToRealname(spusers)
		bl.Env = bugconfig.CacheEidName[eid]
		al.Al = append(al.Al, bl)
	}

	send, _ := json.Marshal(al)
	w.Write(send)
	return

}

// 返回搜索的字符串 和 参数
func searchParamsSql(params *bug.SearchParam) (string, []interface{}) {
	basesql := ""
	args := make([]interface{}, 0)
	if params.Title != "" {
		basesql = basesql + " and bugtitle like ? "
		args = append(args, "%"+params.Title+"%")
	}
	if params.Level != "" {
		basesql = basesql + " and level=? "
		args = append(args, params.Level)
	}

	if params.Project != "" {
		pid := bugconfig.CacheProjectPid[params.Project]
		basesql = basesql + " and pid=? "
		args = append(args, pid)
	}
	return basesql, args
}

func managertotal(basesql string, params *bug.BugManager) (string, []interface{}) {
	basesql = basesql + " where 1=1 "
	args := make([]interface{}, 0)

	if params.Id > 0 {
		basesql = basesql + " and id=? "
		args = append(args, params.Id)
	}
	if params.Title != "" {
		basesql = basesql + " and title=? "
		args = append(args, params.Title)
	}
	if params.Author != "" {
		basesql = basesql + " and uid=? "
		args = append(args, bugconfig.CacheNickNameUid[params.Author])
	}

	return basesql, args
}

func managersearch(basesql string, count int, params *bug.BugManager) (*sql.Rows, error) {
	searchsql, args := managertotal(basesql, params)

	start, end := public.GetPagingLimitAndPage(count, params.Page, params.Limit)

	args = append(args, start)
	args = append(args, end)
	searchsql = searchsql + " order by id desc limit ?,? "

	return db.Mconn.GetRows(searchsql, args...)
}

func getbuglist(r *http.Request, countbasesql string, bugsql string, mytask bool) (*model.AllArticleList, []byte) {

	errorcode := &response.Response{}
	nickname := xmux.GetData(r).Get("nickname").(string)
	searchparam := &bug.SearchParam{} // 接收的参数
	searchq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err)
		return nil, errorcode.ErrorE(err)
	}

	err = json.Unmarshal(searchq, searchparam)
	if err != nil {
		golog.Error(err)
		return nil, errorcode.ErrorE(err)
	}
	al := &model.AllArticleList{}
	// 获取状态
	showstatus := bugconfig.CacheUidFilter[bugconfig.CacheNickNameUid[nickname]]

	//更新缓存
	bugconfig.CacheUidFilter[bugconfig.CacheNickNameUid[nickname]] = showstatus

	// 第二步， 检查level
	if searchparam.Level != "" {
		// 判断这个值是否存在
		if lid, ok := bugconfig.CacheLevelLid[searchparam.Level]; ok {
			bugsql += fmt.Sprintf("and lid=%d ", lid)
			countbasesql += fmt.Sprintf("and lid=%d ", lid)
		} else {
			golog.Error(err)
			return nil, errorcode.Error("没有搜索到")
		}
	}
	// 第三步， 检查Title
	if searchparam.Title != "" {

		bugsql += fmt.Sprintf("and bugtitle like '%s' ", searchparam.Title)
		countbasesql += fmt.Sprintf("and bugtitle like '%s' ", searchparam.Title)

	}
	// 第四步， 检查Project
	if searchparam.Project != "" {
		// 判断这个值是否存在
		if pid, ok := bugconfig.CacheProjectPid[searchparam.Project]; ok {
			bugsql += fmt.Sprintf("and pid=%d ", pid)
			countbasesql += fmt.Sprintf("and pid=%d ", pid)
		} else {
			golog.Error(err)
			return nil, errorcode.Error("没有搜索到")
		}
	}

	if showstatus != "" {
		countbasesql += fmt.Sprintf("and sid in (%s)", showstatus)
		bugsql += fmt.Sprintf("and sid in (%s) ", showstatus)
	}

	err = db.Mconn.GetOne(countbasesql, bugconfig.CacheNickNameUid[nickname]).Scan(&al.Count)
	if err != nil {
		golog.Error(err)
		return nil, errorcode.ErrorE(err)
	}

	// 获取查询的总个数
	start, end := public.GetPagingLimitAndPage(al.Count, searchparam.Page, searchparam.Limit)

	rows, err := db.Mconn.GetRows(bugsql+" limit ?,?", bugconfig.CacheNickNameUid[nickname], start, end)
	if err != nil {
		golog.Error(err)
		return nil, errorcode.ErrorE(err)
	}

	for rows.Next() {
		one := &model.ArticleList{}
		var iid int64
		var sid int64
		var lid int64
		var pid int64
		var eid int64
		var userlist string
		rows.Scan(&one.ID, &one.Date, &iid, &sid, &one.Title, &lid, &pid, &eid, &userlist)
		// 如果不存在这么办， 添加修改的时候需要判断
		one.Importance = bugconfig.CacheIidImportant[iid]
		one.Status = bugconfig.CacheSidStatus[sid]
		one.Level = bugconfig.CacheLidLevel[lid]
		one.Projectname = bugconfig.CachePidName[pid]
		one.Env = bugconfig.CacheEidName[eid]
		// 显示realname

		//如果是我的任务

		for _, v := range strings.Split(userlist, ",") {
			//判断用户是否存在，不存在就 删吗 ， 先不删
			userid32, _ := strconv.Atoi(v)
			if realname, ok := bugconfig.CacheUidRealName[int64(userid32)]; ok {
				one.Handle = append(one.Handle, realname)
			}
		}

		if mytask {
			// 判断是否是自己的任务，先要过滤查询条件，然后查询spusers
			var ismytask bool
			for _, v := range strings.Split(userlist, ",") {
				if v == strconv.FormatInt(bugconfig.CacheNickNameUid[nickname], 10) {
					ismytask = true
					break
				}
			}
			if ismytask {
				for _, v := range strings.Split(userlist, ",") {
					//判断用户是否存在，不存在就 删吗 ， 先不删
					userid32, _ := strconv.Atoi(v)
					if realname, ok := bugconfig.CacheUidRealName[int64(userid32)]; ok {
						one.Handle = append(one.Handle, realname)
					}
				}
			} else {
				continue
			}
		}

		one.Author = bugconfig.CacheUidRealName[bugconfig.CacheNickNameUid[nickname]]
		al.Al = append(al.Al, one)
	}
	return al, nil
}
