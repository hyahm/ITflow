package model

import (
	"itflow/db"

	"github.com/hyahm/golog"
)

type Information struct {
	ID   int64
	Uid  int64
	Bid  int64
	Info string
	Time int64
}

func NewInformationsByBid(bid interface{}) ([]*Information, error) {
	infors := make([]*Information, 0)
	getinfosql := "select id, bid,uid,info,time from informations where bid=?"
	rows, err := db.Mconn.GetRows(getinfosql, bid)
	if err != nil {
		golog.Error(err)
		return infors, err
	}
	for rows.Next() {
		im := &Information{}
		// var uid int64
		rows.Scan(&im.ID, &im.Bid, &im.Uid, &im.Info, &im.Time)
		// im.User = cache.CacheUidRealName[uid]
		infors = append(infors, im)
	}
	return infors, nil
}
