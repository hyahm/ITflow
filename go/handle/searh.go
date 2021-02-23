package handle

import (
	"itflow/db"
	"itflow/internal/bug"
	"itflow/internal/search"
	"itflow/model"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/hyahm/golog"
	"github.com/hyahm/gomysql"
	"github.com/hyahm/xmux"
)

func SearchAllBugs(w http.ResponseWriter, r *http.Request) {

	uid := xmux.GetData(r).Get("uid").(int64)
	uidStr := strconv.FormatInt(uid, 10)
	mybug := xmux.GetData(r).Data.(*search.ReqMyBugFilter)
	al := &model.AllArticleList{
		Al:   make([]*model.ArticleList, 0),
		Page: 1,
	}
	statuslist, err := model.GetMyStatusList(uid)
	if err != nil {
		golog.Error(err)
		w.Write(al.ErrorE(err))
		return
	}
	golog.Info(statuslist)
	// 找出所有跟自己有关的项目， 列出所有项目的bug
	prows, err := db.Mconn.GetRows("select p.id, u.ids from project as p join usergroup as u on p.ugid=u.id;")
	if err != nil {
		golog.Error(err)
		w.Write(al.ErrorE(err))
		return
	}

	myproject := make([]string, 0)
	for prows.Next() {
		var pid string
		var uids string
		err = prows.Scan(&pid, &uids)
		if err != nil {
			golog.Info(err)
			continue
		}
		for _, v := range strings.Split(uids, ",") {
			if uidStr == v {
				myproject = append(myproject, pid)
				break
			}
		}
	}
	prows.Close()
	golog.Info(myproject)
	conditionsql, args := mybug.GetUsefulCondition(uid)
	countArgs := make([]interface{}, 0)
	countArgs = append(countArgs, (gomysql.InArgs)(statuslist).ToInArgs())
	countArgs = append(countArgs, (gomysql.InArgs)(myproject).ToInArgs())
	countArgs = append(countArgs, args...)
	countsql := "select count(id) from bugs where dustbin=false and sid in (?) and pid in (?)"
	db.Mconn.OpenDebug()
	golog.Info(countsql + conditionsql)
	golog.Info(countArgs)
	err = db.Mconn.GetOneIn(countsql+conditionsql, countArgs...).Scan(&al.Count)
	golog.Info(db.Mconn.GetSql())
	if err != nil {
		golog.Error(err)
		w.Write(al.ErrorE(err))
		return
	}

	page, start, end := xmux.GetLimit(al.Count, mybug.Page, mybug.Limit)
	al.Page = page
	searchArgs := make([]interface{}, 0)
	searchArgs = append(searchArgs, (gomysql.InArgs)(statuslist).ToInArgs())
	searchArgs = append(searchArgs, (gomysql.InArgs)(myproject).ToInArgs())
	searchArgs = append(searchArgs, args...)
	searchArgs = append(searchArgs, start, end)
	// searchsql := "select id,createtime,iid,sid,title,lid,pid,eid,spusers from bugs join  on dustbin=true and uid=? "
	searchsql := `select b.id,b.createtime,i.name,s.name,title,l.name,p.name,e.name,spusers,u.realname from bugs as b
	join importants as i
	join status as s
	join level as l
	join project as p
	join environment as e
	join user as u
	on dustbin=false and b.iid = i.id and b.sid = s.id and b.lid = l.id and b.pid=p.id and b.eid = e.id and b.uid=u.id and sid in (?) and pid in (?)`
	rows, err := db.Mconn.GetRowsIn(searchsql+conditionsql+" order by id desc limit ?,?", searchArgs...)

	if err != nil {
		golog.Error(err)
		w.Write(al.ErrorE(err))
		return
	}

	for rows.Next() {
		bug := &model.ArticleList{
			Handle: make([]string, 0),
		}
		var ids string
		err = rows.Scan(&bug.ID,
			&bug.Date, &bug.Importance, &bug.Status, &bug.Title, &bug.Level, &bug.Projectname,
			&bug.Env, &ids, &bug.Author)
		if err != nil {
			golog.Info(err)
			continue
		}
		realnames, err := db.Mconn.GetRowsIn("select realname from user where id in (?)",
			(gomysql.InArgs)(strings.Split(ids, ",")).ToInArgs())
		if err != nil {
			golog.Error(err)
			w.Write(al.ErrorE(err))
			return
		}
		for realnames.Next() {
			var name string
			err = realnames.Scan(&name)
			if err != nil {
				golog.Error(err)
				continue
			}
			bug.Handle = append(bug.Handle, name)
		}
		realnames.Close()
		al.Al = append(al.Al, bug)
	}
	rows.Close()
	w.Write(al.Marshal())
	return

}

func SearchMyBugs(w http.ResponseWriter, r *http.Request) {

	uid := xmux.GetData(r).Get("uid").(int64)
	mybug := xmux.GetData(r).Data.(*search.ReqMyBugFilter)
	al := &model.AllArticleList{
		Al:   make([]*model.ArticleList, 0),
		Page: 1,
	}
	statislist, err := model.GetMyStatusList(uid)
	if err != nil {
		golog.Error(err)
		w.Write(al.ErrorE(err))
		return
	}

	conditionsql, args := mybug.GetUsefulCondition(uid)
	countArgs := make([]interface{}, 0)
	countArgs = append(countArgs, uid, (gomysql.InArgs)(statislist).ToInArgs())
	countArgs = append(countArgs, args...)
	countsql := "select count(id) from bugs where dustbin=false and uid=? and sid in (?)"
	err = db.Mconn.GetOneIn(countsql+conditionsql, countArgs...).Scan(&al.Count)
	if err != nil {
		golog.Error(err)
		w.Write(al.ErrorE(err))
		return
	}

	page, start, end := xmux.GetLimit(al.Count, mybug.Page, mybug.Limit)
	al.Page = page
	searchArgs := make([]interface{}, 0)
	searchArgs = append(searchArgs, uid, (gomysql.InArgs)(statislist).ToInArgs())
	searchArgs = append(searchArgs, args...)
	searchArgs = append(searchArgs, start, end)
	// searchsql := "select id,createtime,iid,sid,title,lid,pid,eid,spusers from bugs join  on dustbin=true and uid=? "
	searchsql := `select b.id,b.createtime, ifnull(i.name, ''), ifnull(s.name, ''),title, 
	ifnull(l.name, ''), ifnull(p.name, ''), ifnull(e.name, ''),spusers, ifnull(u.realname,'') from bugs as b
	left join importants as i on b.iid = i.id 
	left join status as s on b.sid = s.id 
	left join level as l on b.lid = l.id 
	left join project as p on b.pid=p.id 
	left join environment as e on b.eid = e.id 
	left join user as u on  b.uid=u.id 
	where dustbin=false  and b.uid=? and sid in (?)`
	rows, err := db.Mconn.GetRowsIn(searchsql+conditionsql+" order by id desc limit ?,?", searchArgs...)

	if err != nil {
		golog.Error(err)
		w.Write(al.ErrorE(err))
		return
	}

	for rows.Next() {
		bug := &model.ArticleList{
			Handle: make([]string, 0),
		}
		var ids string
		err = rows.Scan(&bug.ID,
			&bug.Date, &bug.Importance, &bug.Status, &bug.Title, &bug.Level, &bug.Projectname,
			&bug.Env, &ids, &bug.Author)
		if err != nil {
			golog.Info(err)
			continue
		}
		realnames, err := db.Mconn.GetRowsIn("select realname from user where id in (?)",
			(gomysql.InArgs)(strings.Split(ids, ",")).ToInArgs())
		if err != nil {
			golog.Error(err)
			w.Write(al.ErrorE(err))
			return
		}
		for realnames.Next() {
			var name string
			err = realnames.Scan(&name)
			if err != nil {
				golog.Error(err)
				continue
			}
			bug.Handle = append(bug.Handle, name)
		}
		realnames.Close()
		al.Al = append(al.Al, bug)
	}

	rows.Close()
	w.Write(al.Marshal())

}

func SearchMyTasks(w http.ResponseWriter, r *http.Request) {
	// 查询任务者中是否有自己
	uid := xmux.GetData(r).Get("uid").(int64)
	strUid := strconv.FormatInt(uid, 10)
	mybug := xmux.GetData(r).Data.(*search.ReqMyBugFilter)
	al := &model.AllArticleList{
		Al: make([]*model.ArticleList, 0),
	}
	statislist, err := model.GetMyStatusList(uid)
	if err != nil {
		golog.Error(err)
		w.Write(al.ErrorE(err))
		return
	}
	golog.Info(statislist)
	conditionsql, args := mybug.GetUsefulCondition(uid)
	countArgs := make([]interface{}, 0)
	countArgs = append(countArgs, (gomysql.InArgs)(statislist).ToInArgs())
	countArgs = append(countArgs, args...)

	countsql := "select id,spusers from bugs where dustbin=false and sid in (?) order by id desc"
	countRows, err := db.Mconn.GetRowsIn(countsql+conditionsql, countArgs...)
	if err != nil {
		golog.Error(err)
		w.Write(al.ErrorE(err))
		return
	}
	myTaskId := make([]string, 0)
	for countRows.Next() {
		var id string
		var spusers string
		err = countRows.Scan(&id, &spusers)
		if err != nil {
			golog.Info(err)
			continue
		}
		for _, v := range strings.Split(spusers, ",") {
			if strUid == v {
				// 查询到自己的任务
				myTaskId = append(myTaskId, id)
				break
			}
		}
	}
	countRows.Close()
	al.Count = len(myTaskId)
	if al.Count == 0 {
		w.Write(al.Marshal())
		return
	}
	golog.Info(myTaskId)
	page, start, end := xmux.GetLimit(al.Count, mybug.Page, mybug.Limit)
	al.Page = page
	// searchsql := "select id,createtime,iid,sid,title,lid,pid,eid,spusers from bugs join  on dustbin=true and uid=? "
	searchsql := `select b.id,b.createtime,ifnull( i.name, ''),ifnull( s.name, ''),title, ifnull(l.name, ''),ifnull( p.name, ''), ifnull(e.name, ''),spusers,u.realname 
	from bugs as b 
	left join importants as i on b.iid = i.id 
	left join status as s on b.sid = s.id 
	left join level as l on b.lid = l.id 
	join project as p on b.pid=p.id 
	left join environment as e on  b.eid = e.id 
	join user as u
	 on b.uid=u.id and  dustbin=false and  b.id in (?)  order by id desc`
	rows, err := db.Mconn.GetRowsIn(searchsql,
		(gomysql.InArgs)(myTaskId[start:start+end]).ToInArgs())

	if err != nil {
		golog.Error(err)
		w.Write(al.ErrorE(err))
		return
	}
	wg := &sync.WaitGroup{}
	for rows.Next() {

		bug := &model.ArticleList{
			Handle: make([]string, 0),
		}
		var ids string
		err = rows.Scan(&bug.ID,
			&bug.Date, &bug.Importance, &bug.Status, &bug.Title, &bug.Level, &bug.Projectname,
			&bug.Env, &ids, &bug.Author)
		if err != nil {
			golog.Info(err)
			continue
		}
		wg.Add(1)
		go func() {
			realnames, err := db.Mconn.GetRowsIn("select realname from user where id in (?)",
				(gomysql.InArgs)(strings.Split(ids, ",")).ToInArgs())
			if err != nil {
				golog.Error(err)
				w.Write(al.ErrorE(err))
				return
			}
			for realnames.Next() {
				var name string
				err = realnames.Scan(&name)
				if err != nil {
					golog.Error(err)
					return
				}
				bug.Handle = append(bug.Handle, name)
			}
			realnames.Close()
			al.Al = append(al.Al, bug)
			wg.Done()
		}()

	}
	wg.Wait()
	rows.Close()
	w.Write(al.Marshal())
	return

}

func SearchBugManager(w http.ResponseWriter, r *http.Request) {

	uid := xmux.GetData(r).Get("uid").(int64)
	mybug := xmux.GetData(r).Data.(*search.ReqMyBugFilter)
	al := &model.AllArticleList{
		Al: make([]*model.ArticleList, 0),
	}
	statislist, err := model.GetMyStatusList(uid)
	if err != nil {
		golog.Error(err)
		w.Write(al.ErrorE(err))
		return
	}

	conditionsql, args := mybug.GetUsefulCondition(uid)
	args = append(args, (gomysql.InArgs)(statislist).ToInArgs())
	countsql := "select count(id) from bugs where dustbin=true and sid in (?)"
	err = db.Mconn.GetOneIn(countsql+conditionsql, args...).Scan(&al.Count)
	if err != nil {
		golog.Error(err)
		w.Write(al.ErrorE(err))
		return
	}

	page, start, end := xmux.GetLimit(al.Count, mybug.Page, mybug.Limit)
	al.Page = page
	args = append(args, start, end)
	// searchsql := "select id,createtime,iid,sid,title,lid,pid,eid,spusers from bugs join  on dustbin=true and uid=? "
	searchsql := `select b.id,b.createtime,i.name,s.name,title,l.name,p.name,e.name,spusers,u.realname from bugs as b
	join importants as i
	join status as s
	join level as l
	join project as p
	join environment as e
	join user as u
	on dustbin=true and b.iid = i.id and b.sid = s.id and b.lid = l.id and b.pid=p.id and b.eid = e.id and b.uid=u.id and sid in (?)`
	rows, err := db.Mconn.GetRowsIn(searchsql+conditionsql+" limit ?,?", args...)

	if err != nil {
		golog.Error(err)
		w.Write(al.ErrorE(err))
		return
	}

	for rows.Next() {
		bug := &model.ArticleList{
			Handle: make([]string, 0),
		}
		var ids string
		err = rows.Scan(&bug.ID,
			&bug.Date, &bug.Importance, &bug.Status, &bug.Title, &bug.Level, &bug.Projectname,
			&bug.Env, &ids, &bug.Author)
		if err != nil {
			golog.Info(err)
			continue
		}
		realnames, err := db.Mconn.GetRowsIn("select realname from user where id in (?)",
			(gomysql.InArgs)(strings.Split(ids, ",")).ToInArgs())
		if err != nil {
			golog.Error(err)
			w.Write(al.ErrorE(err))
			return
		}
		for realnames.Next() {
			var name string
			err = realnames.Scan(&name)
			if err != nil {
				golog.Error(err)
				continue
			}
			bug.Handle = append(bug.Handle, name)
		}
		realnames.Close()
		al.Al = append(al.Al, bug)
	}

	rows.Close()
	w.Write(al.Marshal())
	return

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

	// if params.Project != "" {
	// 	pid := cache.CacheProjectPid[params.Project]
	// 	basesql = basesql + " and pid=? "
	// 	args = append(args, pid)
	// }
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
		// args = append(args, cache.CacheNickNameUid[params.Author])
	}

	return basesql, args
}

func getbuglist(r *http.Request, countbasesql string, bugsql string, mytask bool) (*model.AllArticleList, []byte) {

	// 	errorcode := &response.Response{}
	// 	nickname := xmux.GetData(r).Get("nickname").(string)
	// 	searchparam := &bug.SearchParam{} // 接收的参数
	// 	searchq, err := ioutil.ReadAll(r.Body)
	// 	if err != nil {
	// 		golog.Error(err)
	// 		return nil, errorcode.ErrorE(err)
	// 	}

	// 	err = json.Unmarshal(searchq, searchparam)
	// 	if err != nil {
	// 		golog.Error(err)
	// 		return nil, errorcode.ErrorE(err)
	// 	}
	// 	al := &model.AllArticleList{}
	// 	// 获取状态
	// 	// showstatus := cache.CacheUidFilter[cache.CacheNickNameUid[nickname]]

	// 	//更新缓存

	// 	// 第二步， 检查level
	// 	if searchparam.Level != "" {
	// 		// 判断这个值是否存在
	// 		if lid, ok := cache.CacheLevelLid[searchparam.Level]; ok {
	// 			bugsql += fmt.Sprintf("and lid=%d ", lid)
	// 			countbasesql += fmt.Sprintf("and lid=%d ", lid)
	// 		} else {
	// 			golog.Error(err)
	// 			return nil, errorcode.Error("没有搜索到")
	// 		}
	// 	}
	// 	// 第三步， 检查Title
	// 	if searchparam.Title != "" {

	// 		bugsql += fmt.Sprintf("and title like '%s' ", searchparam.Title)
	// 		countbasesql += fmt.Sprintf("and title like '%s' ", searchparam.Title)

	// 	}
	// 	// 第四步， 检查Project
	// 	if searchparam.Project != "" {
	// 		// 判断这个值是否存在
	// 		if pid, ok := cache.CacheProjectPid[searchparam.Project]; ok {
	// 			bugsql += fmt.Sprintf("and pid=%d ", pid)
	// 			countbasesql += fmt.Sprintf("and pid=%d ", pid)
	// 		} else {
	// 			golog.Error(err)
	// 			return nil, errorcode.Error("没有搜索到")
	// 		}
	// 	}

	// 	if showstatus != "" {
	// 		countbasesql += fmt.Sprintf("and sid in (%s)", showstatus)
	// 		bugsql += fmt.Sprintf("and sid in (%s) ", showstatus)
	// 	}

	// 	err = db.Mconn.GetOne(countbasesql, cache.CacheNickNameUid[nickname]).Scan(&al.Count)
	// 	if err != nil {
	// 		golog.Error(err)
	// 		return nil, errorcode.ErrorE(err)
	// 	}

	// 	// 获取查询的总个数
	// 	start, end := xmux.GetLimit(al.Count, searchparam.Page, searchparam.Limit)

	// 	rows, err := db.Mconn.GetRows(bugsql+" limit ?,?", cache.CacheNickNameUid[nickname], start, end)
	// 	if err != nil {
	// 		golog.Error(err)
	// 		return nil, errorcode.ErrorE(err)
	// 	}

	// 	for rows.Next() {
	// 		one := &model.ArticleList{}
	// 		var iid cache.ImportantId
	// 		var sid cache.StatusId
	// 		var lid cache.LevelId
	// 		var pid cache.ProjectId
	// 		var eid cache.EnvId
	// 		var userlist string
	// 		rows.Scan(&one.ID, &one.Date, &iid, &sid, &one.Title, &lid, &pid, &eid, &userlist)
	// 		// 如果不存在这么办， 添加修改的时候需要判断
	// 		one.Importance = cache.CacheIidImportant[iid]
	// 		one.Status = cache.CacheSidStatus[sid]
	// 		one.Level = cache.CacheLidLevel[lid]
	// 		one.Projectname = cache.CachePidProject[pid]
	// 		one.Env = cache.CacheEidEnv[eid]
	// 		// 显示realname

	// 		//如果是我的任务

	// 		for _, v := range strings.Split(userlist, ",") {
	// 			//判断用户是否存在，不存在就 删吗 ， 先不删
	// 			// userid32, _ := strconv.Atoi(v)
	// 			// if realname, ok := cache.CacheUidRealName[int64(userid32)]; ok {
	// 			// 	one.Handle = append(one.Handle, realname)
	// 			// }
	// 		}

	// 		if mytask {
	// 			// 判断是否是自己的任务，先要过滤查询条件，然后查询spusers
	// 			var ismytask bool
	// 			for _, v := range strings.Split(userlist, ",") {
	// 				if v == strconv.FormatInt(cache.CacheNickNameUid[nickname], 10) {
	// 					ismytask = true
	// 					break
	// 				}
	// 			}
	// 			if ismytask {
	// 				for _, v := range strings.Split(userlist, ",") {
	// 					//判断用户是否存在，不存在就 删吗 ， 先不删
	// 					// userid32, _ := strconv.Atoi(v)
	// 					// if realname, ok := cache.CacheUidRealName[int64(userid32)]; ok {
	// 					// 	one.Handle = append(one.Handle, realname)
	// 					// }
	// 				}
	// 			} else {
	// 				continue
	// 			}
	// 		}

	// 		one.Author = cache.CacheUidRealName[cache.CacheNickNameUid[nickname]]
	// 		al.Al = append(al.Al, one)
	// 	}
	return nil, nil
}
