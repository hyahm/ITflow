package handle

import (
	"encoding/json"
	"fmt"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/log"
	"itflow/internal/response"
	"itflow/pkg/pager"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func SearchLog(w http.ResponseWriter, r *http.Request) {

	alllog := xmux.GetData(r).Data.(*log.Search_log)
	listlog := &log.Loglist{}

	basesql := "select id,exectime,classify,action,ip,username from log "
	endsql := ""
	// 如果搜索了类别
	if alllog.Classify != "" {
		//判断是否在类别数组中
		var realclassify bool
		for _, v := range cache.CLASSIFY {
			if v == alllog.Classify {
				realclassify = true
				break
			}
		}
		if !realclassify {
			golog.Debug("没有找到key")

			w.Write(listlog.Error("没有找到key"))
			return
		}
		endsql = fmt.Sprintf("where classify='%v' ", alllog.Classify)
	}
	// 如果有时间选择，并且不为0
	if alllog.StartTime != 0 {
		if len(endsql) == 0 {
			endsql = fmt.Sprintf("where exectime between %d and %d ", alllog.StartTime, alllog.EndTime)
		} else {
			endsql += fmt.Sprintf(" and exectime between %d and %d ", alllog.StartTime, alllog.EndTime)
		}
	}

	//获取总行数
	countsql := "select count(id) from log " + endsql
	err := db.Mconn.GetOne(countsql).Scan(&alllog.Count)
	if err != nil {
		golog.Error(err)
		w.Write(listlog.ErrorE(err))
		return
	}
	if alllog.Count == 0 {
		golog.Error("no rows")
		w.Write(listlog.NoRows())
		return
	}

	start, end := alllog.GetPagingLimitAndPage()
	rows, err := db.Mconn.GetRows(basesql+endsql+" order by id desc limit ?,?", start, end)
	if err != nil {
		golog.Error(err)
		w.Write(listlog.ErrorE(err))
		return
	}

	for rows.Next() {
		one := &log.LogRow{}
		//basesql := "select id,exectime,classify,content,ip from log "
		rows.Scan(&one.Id, &one.Exectime, &one.Classify, &one.Action, &one.Ip, &one.UserName)
		listlog.LogList = append(listlog.LogList, one)
	}
	listlog.Page = alllog.Page
	listlog.Count = alllog.Count
	send, _ := json.Marshal(listlog)
	w.Write(send)
	return

}

func LogClassify(w http.ResponseWriter, r *http.Request) {

	send, _ := json.Marshal(cache.CLASSIFY)
	w.Write(send)
	return

}

func LogList(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	sl := xmux.GetData(r).Data.(*log.Search_log)

	var count int
	countsql := "select count(id) from log"

	err := db.Mconn.GetOne(countsql).Scan(&count)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	start, end := pager.GetPagingLimitAndPage(count, sl.Page, sl.Limit)
	alllog := &log.Loglist{
		Count: count,
	}

	dsql := "select id,exectime,classify,action,ip,username from log order by id desc limit ?,?"
	rows, err := db.Mconn.GetRows(dsql, start, end)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	for rows.Next() {
		log := &log.LogRow{}
		rows.Scan(&log.Id, &log.Exectime, &log.Classify, &log.Action, &log.Ip, &log.UserName)
		alllog.LogList = append(alllog.LogList, log)
	}
	send, _ := json.Marshal(alllog)
	w.Write(send)
	return

}
