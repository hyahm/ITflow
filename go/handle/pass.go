package handle

import (
	"itflow/model"
	"itflow/response"
	"net/http"
	"time"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

type RequestPass struct {
	Bid     int64   `json:"bid" `     // bugid
	Remark  string  `json:"remark" `  // bugid
	SpUsers []int64 `json:"spusers" ` // bugid
}

func PassBug(w http.ResponseWriter, r *http.Request) {
	rp := xmux.GetInstance(r).Data.(*RequestPass)
	uid := xmux.GetInstance(r).Get("uid").(int64)
	// 更新bug表
	bug := model.Bug{
		ID:   rp.Bid,
		Uids: rp.SpUsers,
		Sid:  model.Default.Pass, // 转交的默认状态
	}
	err := bug.Update()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	// 新增information表
	information := model.Information{
		Uid:  uid,
		Bid:  rp.Bid,
		Info: rp.Remark,
		Time: time.Now().Unix(),
	}
	err = information.Insert()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	w.Write(response.Success())
}

func TaskList(w http.ResponseWriter, r *http.Request) {

	// errorcode := &response.Response{}

	// al := &model.AllArticleList{}
	// uid := xmux.GetInstance(r).Get("uid").(int64)

	// getaritclesql := `select id,createtime,importent,s.name,title,u.realname,l.name,p.name,spusers from bugs as b
	// join user as u
	// join level as l
	// join project as p
	// join status as s
	// 		on b.id in (select bid from userandbug where b.uid=?)  order by id desc `

	// rows, err := db.Mconn.GetRows(getaritclesql, uid)

	// if err != nil {
	// 	golog.Error(err)
	// 	w.Write(errorcode.ErrorE(err))
	// 	return
	// }
	// for rows.Next() {
	// 	sendlist := &model.ArticleList{}
	// 	var spusers string
	// 	rows.Scan(&sendlist.ID, &sendlist.Date, &sendlist.Importance, &sendlist.Status,
	// 		&sendlist.Title, &sendlist.Author, &sendlist.Level, &sendlist.Projectname, &spusers)
	// 	sendlist.Handle = assist.FormatUserlistToShow(spusers)

	// 	al.Al = append(al.Al, sendlist)
	// }

	// rows.Close()
	// send, _ := json.Marshal(al)
	// w.Write(send)
	// return

}
