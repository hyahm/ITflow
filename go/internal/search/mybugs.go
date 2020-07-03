package search

import (
	"fmt"
	"itflow/cache"

	"github.com/hyahm/golog"
)

type ReqMyBugFilter struct {
	page        int              `json:"page"`
	Limit       int              `json:"limit"`
	Level       cache.Level      `json:"level"`
	Project     string           `json:"project"`
	Title       string           `json:"title"`
	ShowsStatus cache.StatusList // 这个应该从数据库获取
}

func GetFilterBugs() {
	// 搜索的条件，
	// page: 1,
	// limit: 10,
	// level: '',
	// project: '',
	// title: '',
	// showstatus: []
	// 分2部， 先获取符合条件的

	// 获取页码
}

func (rmf *ReqMyBugFilter) GetUsefulCondition(uid int64) ([]byte, error) {
	// 获取有用的条件
	countbasesql := "select count(id) from bugs where dustbin=0 and uid=? "
	bugsql := "select id,createtime,iid,sid,title,lid,pid,eid,spusers from bugs where dustbin=0 and uid=? "
	if rmf.Level != "" {
		// 判断这个值是否存在

		if lid := rmf.Level.Id(); lid != 0 {
			bugsql += fmt.Sprintf("and lid=%d ", lid)
			countbasesql += fmt.Sprintf("and lid=%d ", lid)
		} else {
			rmf.Level = ""
		}
	}
	if rmf.Title != "" {

		bugsql += fmt.Sprintf("and title like '%s' ", rmf.Title)
		countbasesql += fmt.Sprintf("and title like '%s' ", rmf.Title)

	}

	if rmf.Project != "" {
		// 判断这个值是否存在
		if pid, ok := cache.CacheProjectPid[rmf.Project]; ok {
			bugsql += fmt.Sprintf("and pid=%d ", pid)
			countbasesql += fmt.Sprintf("and pid=%d ", pid)
		} else {
			rmf.Level = ""
		}
	}
	// 获取此用户能看到的状态
	golog.Info(uid)
	showstatus := cache.CacheUidFilter[uid].ToShow()
	if len(showstatus) == 0 {
		// 没选择状态，返回空数组
		golog.Info("没选择状态，返回空数组")
		// return nil
	}

	countbasesql += fmt.Sprintf("and sid in (%s)", cache.CacheUidFilter[uid])
	bugsql += fmt.Sprintf("and sid in (%s) ", cache.CacheUidFilter[uid])

	golog.Info(countbasesql)
	golog.Info(bugsql)
	return nil, nil
}

// func getbuglist(r *http.Request, countbasesql string, bugsql string, mytask bool) (*model.AllArticleList, []byte) {

// 	// errorcode := &response.Response{}
// 	// nickname := xmux.GetData(r).Get("nickname").(string)
// 	// searchparam := &bug.SearchParam{} // 接收的参数
// 	// searchq, err := ioutil.ReadAll(r.Body)
// 	// if err != nil {
// 	// 	golog.Error(err)
// 	// 	return nil, errorcode.ErrorE(err)
// 	// }

// 	// err = json.Unmarshal(searchq, searchparam)
// 	// if err != nil {
// 	// 	golog.Error(err)
// 	// 	return nil, errorcode.ErrorE(err)
// 	// }
// 	al := &model.AllArticleList{}
// 	// 获取状态
// 	showstatus := cache.CacheUidFilter[cache.CacheNickNameUid[nickname]]

// 	//更新缓存
// 	cache.CacheUidFilter[cache.CacheNickNameUid[nickname]] = showstatus

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
// 	start, end := pager.GetPagingLimitAndPage(al.Count, searchparam.Page, searchparam.Limit)

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
// 		var pid int64
// 		var eid int64
// 		var userlist string
// 		rows.Scan(&one.ID, &one.Date, &iid, &sid, &one.Title, &lid, &pid, &eid, &userlist)
// 		// 如果不存在这么办， 添加修改的时候需要判断
// 		one.Importance = cache.CacheIidImportant[iid]
// 		one.Status = cache.CacheSidStatus[sid]
// 		one.Level = cache.CacheLidLevel[lid]
// 		one.Projectname = cache.CachePidName[pid]
// 		one.Env = cache.CacheEidName[eid]
// 		// 显示realname

// 		//如果是我的任务

// 		for _, v := range strings.Split(userlist, ",") {
// 			//判断用户是否存在，不存在就 删吗 ， 先不删
// 			userid32, _ := strconv.Atoi(v)
// 			if realname, ok := cache.CacheUidRealName[int64(userid32)]; ok {
// 				one.Handle = append(one.Handle, realname)
// 			}
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
// 					userid32, _ := strconv.Atoi(v)
// 					if realname, ok := cache.CacheUidRealName[int64(userid32)]; ok {
// 						one.Handle = append(one.Handle, realname)
// 					}
// 				}
// 			} else {
// 				continue
// 			}
// 		}

// 		one.Author = cache.CacheUidRealName[cache.CacheNickNameUid[nickname]]
// 		al.Al = append(al.Al, one)
// 	}
// 	return al, nil
// }
