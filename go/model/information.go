package model

import (
	"itflow/db"
	"itflow/internal/comment"

	"github.com/hyahm/golog"
)

type Information struct {
	ID   int64
	Uid  int64
	Bid  int64
	Info string
	Time int64
}

func NewInformationsByBid(bid interface{}, cms []*comment.Informations) error {
	// sl.Comments = make([]*comment.Informations, len(cc))
	getinfosql := "select u.realname,info,time from informations as i join user as u on bid=? and u.id=i.uid"
	rows, err := db.Mconn.GetRows(getinfosql, bid)
	if err != nil {
		golog.Error(err)
		return err
	}
	for rows.Next() {
		im := &comment.Informations{}
		// var uid int64
		rows.Scan(&im.User, &im.Info, &im.Date)
		// im.User = cache.CacheUidRealName[uid]
		cms = append(cms, im)
	}
	return nil
}
