package handle

import (
	"itflow/classify"
	"itflow/db"
	"itflow/internal/log"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func SearchLog(w http.ResponseWriter, r *http.Request) {

	alllog := xmux.GetInstance(r).Data.(*log.SearchLog)
	listlog := &log.Loglist{}
	args := make([]interface{}, 0)

	condition := ""
	// 如果搜索了类别
	if alllog.Classify != "" {
		//判断是否在类别数组中
		condition = " and classify=? "
		args = append(args, alllog.Classify)
	}
	// 如果有时间选择，并且不为0
	if alllog.StartTime != 0 {
		condition += " and exectime between ? and ? "
		args = append(args, alllog.StartTime, alllog.EndTime)
	}

	//获取总行数
	countsql := "select count(id) from log where 1=1" + condition
	err := db.Mconn.GetOne(countsql, args...).Scan(&listlog.Count)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	if listlog.Count == 0 {
		golog.Error("no rows")
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = "no rows"
		return
	}
	page, start, end := xmux.GetLimit(listlog.Count, alllog.Page, alllog.Limit)
	listlog.Page = page
	args = append(args, start, end)
	basesql := "select l.id,exectime,classify,action,ip,u.realname from log as l join user as u on l.uid=u.id "
	rows, err := db.Mconn.GetRows(basesql+condition+" order by id desc limit ?,?", args...)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}

	for rows.Next() {
		one := &log.LogRow{}
		//basesql := "select id,exectime,classify,content,ip from log "
		rows.Scan(&one.Id, &one.Exectime, &one.Classify, &one.Action, &one.Ip, &one.UserName)
		listlog.LogList = append(listlog.LogList, one)
	}
	rows.Close()
	xmux.GetInstance(r).Response.(*response.Response).Data = listlog

}

func LogClassify(w http.ResponseWriter, r *http.Request) {
	xmux.GetInstance(r).Response.(*response.Response).Data = classify.CLASSIFY

}

func LogList(w http.ResponseWriter, r *http.Request) {

	sl := xmux.GetInstance(r).Data.(*log.SearchLog)

	var count int
	countsql := "select count(id) from log"

	err := db.Mconn.GetOne(countsql).Scan(&count)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}

	page, start, end := xmux.GetLimit(count, sl.Page, sl.Limit)
	alllog := &log.Loglist{
		Count: count,
		Page:  page,
	}

	dsql := "select id,exectime,classify,action,ip,username from log order by id desc limit ?,?"
	rows, err := db.Mconn.GetRows(dsql, start, end)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	for rows.Next() {
		log := &log.LogRow{}
		rows.Scan(&log.Id, &log.Exectime, &log.Classify, &log.Action, &log.Ip, &log.UserName)
		alllog.LogList = append(alllog.LogList, log)
	}
	rows.Close()
	xmux.GetInstance(r).Response.(*response.Response).Data = alllog

}
