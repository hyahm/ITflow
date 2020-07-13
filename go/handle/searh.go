package handle

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/bug"
	"itflow/internal/response"
	"itflow/internal/search"
	"itflow/model"
	"itflow/pkg/pager"
	"net/http"
	"strconv"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func SearchAllBugs(w http.ResponseWriter, r *http.Request) {

	uid := xmux.GetData(r).Get("uid").(int64)
	mybug := xmux.GetData(r).Data.(*search.ReqMyBugFilter)
	// mybug.GetUsefulCondition(uid)
	countsql := "select count(id) from bugs where dustbin=0  "
	searchsql := "select id,createtime,iid,sid,title,lid,pid,eid,spusers from bugs where dustbin=0  "

	sch, err := mybug.GetUsefulCondition(uid, countsql, searchsql)
	if err != nil {
		if err == search.ErrorNoStatus {
			al := &model.AllArticleList{
				Al: make([]*model.ArticleList, 0),
			}
			w.Write(al.Marshal())
			return
		}
		golog.Error(err)
		al := &model.AllArticleList{
			Al:   make([]*model.ArticleList, 0),
			Code: 1,
			Msg:  err.Error(),
		}
		w.Write(al.Marshal())
		return
	}

	w.Write(sch.GetMyBugs())
	return

}

func SearchMyBugs(w http.ResponseWriter, r *http.Request) {
	// 搜索的条件，
	// page: 1,
	// limit: 10,
	// level: '',
	// project: '',
	// title: '',
	// showstatus: []

	uid := xmux.GetData(r).Get("uid").(int64)
	mybug := xmux.GetData(r).Data.(*search.ReqMyBugFilter)
	// mybug.GetUsefulCondition(uid)
	countsql := fmt.Sprintf("select count(id) from bugs where dustbin=0 and uid=%d ", uid)
	searchsql := fmt.Sprintf("select id,createtime,iid,sid,title,lid,pid,eid,spusers from bugs where dustbin=0 and uid=%d ", uid)

	sch, err := mybug.GetUsefulCondition(uid, countsql, searchsql)
	if err != nil {
		if err == search.ErrorNoStatus {
			al := &model.AllArticleList{
				Al: make([]*model.ArticleList, 0),
			}
			w.Write(al.Marshal())
			return
		}
		golog.Error(err)
		al := &model.AllArticleList{
			Al:   make([]*model.ArticleList, 0),
			Code: 1,
			Msg:  err.Error(),
		}
		w.Write(al.Marshal())
		return
	}
	w.Write(sch.GetMyBugs())
	return

}

func SearchMyTasks(w http.ResponseWriter, r *http.Request) {
	uid := xmux.GetData(r).Get("uid").(int64)
	mybug := xmux.GetData(r).Data.(*search.ReqMyBugFilter)
	countsql := "select spusers from bugs where dustbin=0 "
	searchsql := "select id,createtime,iid,sid,title,lid,pid,eid,spusers from bugs where dustbin=0 "
	sch, err := mybug.GetUsefulCondition(uid, countsql, searchsql)
	if err != nil {
		if err == search.ErrorNoStatus {
			al := &model.AllArticleList{
				Al: make([]*model.ArticleList, 0),
			}
			w.Write(al.Marshal())
			return
		}
		golog.Error(err)
		al := &model.AllArticleList{
			Al:   make([]*model.ArticleList, 0),
			Code: 1,
			Msg:  err.Error(),
		}
		w.Write(al.Marshal())
		return
	}
	// 获取状态
	w.Write(sch.GetMyTasks())
	return

}

func SearchBugManager(w http.ResponseWriter, r *http.Request) {

	// errorcode := &response.Response{}
	uid := xmux.GetData(r).Get("uid").(int64)
	mybug := xmux.GetData(r).Data.(*search.ReqMyBugFilter)
	// mybug.GetUsefulCondition(uid)
	countsql := "select count(id) from bugs where dustbin=1 "
	searchsql := "select id,createtime,iid,sid,title,lid,pid,eid,spusers from bugs where dustbin=1 "
	sch, err := mybug.GetUsefulCondition(uid, countsql, searchsql)
	if err != nil {
		if err == search.ErrorNoStatus {
			al := &model.AllArticleList{
				Al: make([]*model.ArticleList, 0),
			}
			w.Write(al.Marshal())
			return
		}
		golog.Error(err)
		al := &model.AllArticleList{
			Al:   make([]*model.ArticleList, 0),
			Code: 1,
			Msg:  err.Error(),
		}
		w.Write(al.Marshal())
		return
	}

	w.Write(sch.GetMyBugs())

}

// 返回搜索的字符串 和 参数
func searchParamsSql(params *bug.SearchParam) (string, []interface{}) {
	basesql := ""
	args := make([]interface{}, 0)
	if params.Title != "" {
		basesql = basesql + " and title like ? "
		args = append(args, "%"+params.Title+"%")
	}
	if params.Level != "" {
		basesql = basesql + " and level=? "
		args = append(args, params.Level)
	}

	if params.Project != "" {
		pid := cache.CacheProjectPid[params.Project]
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
		args = append(args, cache.CacheNickNameUid[params.Author])
	}

	return basesql, args
}

func managersearch(basesql string, count int, params *bug.BugManager) (*sql.Rows, error) {
	searchsql, args := managertotal(basesql, params)

	start, end := pager.GetPagingLimitAndPage(count, params.Page, params.Limit)

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
	showstatus := cache.CacheUidFilter[cache.CacheNickNameUid[nickname]]

	//更新缓存
	cache.CacheUidFilter[cache.CacheNickNameUid[nickname]] = showstatus

	// 第二步， 检查level
	if searchparam.Level != "" {
		// 判断这个值是否存在
		if lid, ok := cache.CacheLevelLid[searchparam.Level]; ok {
			bugsql += fmt.Sprintf("and lid=%d ", lid)
			countbasesql += fmt.Sprintf("and lid=%d ", lid)
		} else {
			golog.Error(err)
			return nil, errorcode.Error("没有搜索到")
		}
	}
	// 第三步， 检查Title
	if searchparam.Title != "" {

		bugsql += fmt.Sprintf("and title like '%s' ", searchparam.Title)
		countbasesql += fmt.Sprintf("and title like '%s' ", searchparam.Title)

	}
	// 第四步， 检查Project
	if searchparam.Project != "" {
		// 判断这个值是否存在
		if pid, ok := cache.CacheProjectPid[searchparam.Project]; ok {
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

	err = db.Mconn.GetOne(countbasesql, cache.CacheNickNameUid[nickname]).Scan(&al.Count)
	if err != nil {
		golog.Error(err)
		return nil, errorcode.ErrorE(err)
	}

	// 获取查询的总个数
	start, end := pager.GetPagingLimitAndPage(al.Count, searchparam.Page, searchparam.Limit)

	rows, err := db.Mconn.GetRows(bugsql+" limit ?,?", cache.CacheNickNameUid[nickname], start, end)
	if err != nil {
		golog.Error(err)
		return nil, errorcode.ErrorE(err)
	}

	for rows.Next() {
		one := &model.ArticleList{}
		var iid cache.ImportantId
		var sid cache.StatusId
		var lid cache.LevelId
		var pid cache.ProjectId
		var eid cache.EnvId
		var userlist string
		rows.Scan(&one.ID, &one.Date, &iid, &sid, &one.Title, &lid, &pid, &eid, &userlist)
		// 如果不存在这么办， 添加修改的时候需要判断
		one.Importance = cache.CacheIidImportant[iid]
		one.Status = cache.CacheSidStatus[sid]
		one.Level = cache.CacheLidLevel[lid]
		one.Projectname = cache.CachePidProject[pid]
		one.Env = cache.CacheEidEnv[eid]
		// 显示realname

		//如果是我的任务

		for _, v := range strings.Split(userlist, ",") {
			//判断用户是否存在，不存在就 删吗 ， 先不删
			userid32, _ := strconv.Atoi(v)
			if realname, ok := cache.CacheUidRealName[int64(userid32)]; ok {
				one.Handle = append(one.Handle, realname)
			}
		}

		if mytask {
			// 判断是否是自己的任务，先要过滤查询条件，然后查询spusers
			var ismytask bool
			for _, v := range strings.Split(userlist, ",") {
				if v == strconv.FormatInt(cache.CacheNickNameUid[nickname], 10) {
					ismytask = true
					break
				}
			}
			if ismytask {
				for _, v := range strings.Split(userlist, ",") {
					//判断用户是否存在，不存在就 删吗 ， 先不删
					userid32, _ := strconv.Atoi(v)
					if realname, ok := cache.CacheUidRealName[int64(userid32)]; ok {
						one.Handle = append(one.Handle, realname)
					}
				}
			} else {
				continue
			}
		}

		one.Author = cache.CacheUidRealName[cache.CacheNickNameUid[nickname]]
		al.Al = append(al.Al, one)
	}
	return al, nil
}
