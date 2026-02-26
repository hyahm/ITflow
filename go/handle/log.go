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
	listlog := &log.Loglist{
		LogList: make([]*log.LogRow, 0),
	}

	query := db.Gorm.Table("log as l").Joins("inner join user as u on l.uid=u.id")

	if alllog.StartTime != 0 {
		query = query.Where("l.exectime between ? and ? ", alllog.StartTime, alllog.EndTime)
	}
	if alllog.Classify != "" {
		//判断是否在类别数组中
		query = query.Where("l.classify=?", alllog.Classify)
	}
	var count int64
	err := query.Count(&count).Error
	if err != nil {
		return
	}
	if count > 0 {
		err = query.Select("l.id,l.create_time,l.classify,l.action,ip,u.realname as user_name").Offset((alllog.Page - 1) * alllog.Limit).Limit(alllog.Limit).Find(&listlog.LogList).Error
	}
	listlog.Count = count
	listlog.Page = alllog.Page

	golog.Info(listlog)
	xmux.GetInstance(r).Response.(*response.Response).Data = listlog

}

func LogClassify(w http.ResponseWriter, r *http.Request) {
	xmux.GetInstance(r).Response.(*response.Response).Data = classify.CLASSIFY

}

// func LogList(w http.ResponseWriter, r *http.Request) {

// 	sl := xmux.GetInstance(r).Data.(*log.SearchLog)

// 	var count int
// 	countsql := "select count(id) from log"

// 	err := db.Mconn.GetOne(countsql).Scan(&count)
// 	if err != nil {
// 		golog.Error(err)
// 		xmux.GetInstance(r).Response.(*response.Response).Code = 1
// 		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
// 		return
// 	}

// 	page, start, end := xmux.GetLimit(count, sl.Page, sl.Limit)
// 	alllog := &log.Loglist{
// 		Count: count,
// 		Page:  page,
// 	}

// 	dsql := "select id,exectime,classify,action,ip,username from log order by id desc limit ?,?"
// 	rows, err := db.Mconn.GetRows(dsql, start, end)
// 	if err != nil {
// 		golog.Error(err)
// 		xmux.GetInstance(r).Response.(*response.Response).Code = 1
// 		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
// 		return
// 	}
// 	for rows.Next() {
// 		log := &log.LogRow{}
// 		rows.Scan(&log.Id, &log.Exectime, &log.Classify, &log.Action, &log.Ip, &log.UserName)
// 		alllog.LogList = append(alllog.LogList, log)
// 	}
// 	rows.Close()
// 	xmux.GetInstance(r).Response.(*response.Response).Data = alllog

// }
