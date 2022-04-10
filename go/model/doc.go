package model

import (
	"errors"
	"itflow/db"
	"itflow/dockercompose"
	"os"
	"path/filepath"
	"time"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
)

const BASEPORT = 4000

type Doc struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	UID      int64  `json:"uid"`
	Created  int64  `json:"created"`
	GitUrl   string `json:"giturl"`
	Updated  int64  `json:"uptime"`
	Dir      string `json:"dir"`
	Port     int    `json:"port"`
	KID      int64  `json:"kid"`
	AuthName string `json:"authname"`
}

func CheckName(name string) error {
	var count int
	err := db.Mconn.GetOne("select count(id) from doc where name=?", name).Scan(&count)
	if err != nil {
		return err
	}
	if count != 0 {
		return errors.New("name duplicate")
	}
	return nil
}

func GetAllDocs(uid int64) ([]*Doc, error) {
	docs := make([]*Doc, 0)
	getcategorysql := "select d.id,d.name,d.uid,d.created,giturl,d.uptime,kid,ifnull(a.name,'')  from doc as d left join auth as a on d.kid = a.id "
	rows, err := db.Mconn.GetRows(getcategorysql)
	if err != nil {
		golog.Error(err)
		return nil, err
	}

	for rows.Next() {
		doc := &Doc{}
		rows.Scan(&doc.ID, &doc.Name, &doc.UID, &doc.Created, &doc.GitUrl, &doc.Updated, &doc.KID, &doc.AuthName)
		if doc.UID == uid {
			docs = append(docs, doc)
			continue
		}

	}
	return docs, nil
}

func GetPort(name string) int {
	// 如果port 为0, name就是没找到
	var port int
	db.Mconn.GetOne("select port from doc where name=?", name).Scan(&port)
	return port
}

func GetPortByName(name string, uid int64) int {
	// 如果port 为0, name就是没找到
	var port int
	db.Mconn.GetOne("select port from doc where name=? and uid=?", name, uid).Scan(&port)
	return port
}

func ChecDocName(name string) error {
	//
	var count int
	err := db.Mconn.GetOne("select count(id) from doc where name=?", name).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("file duplicate")
	}
	return nil
}

func (d *Doc) NewGit() (*dockercompose.Git, error) {
	_git := &dockercompose.Git{}
	_git.Url = d.GitUrl

	// 创建根文档目录
	if _, err := os.Stat(goconfig.ReadPath("scs.path")); err == os.ErrNotExist {
		e := os.MkdirAll(goconfig.ReadPath("scs.path"), 0755)
		if e != nil {
			return nil, e
		}
	}

	_git.Path = filepath.Join(goconfig.ReadPath("scs.path"), d.Name)

	if d.KID > 0 {
		auth, err := NewAuthById(d.KID)
		if err != nil {
			return nil, err
		}
		_git.AuthType = auth.Typ
		_git.AuthType = auth.Typ
		_git.User = auth.User
		_git.Password = auth.Password
		_git.Pri = auth.Pri
	}

	return _git, nil
}

func (d *Doc) Insert(uid int64) (err error) {

	base := goconfig.ReadPath("scs.path")
	tmp := filepath.Clean(filepath.Join(base, d.Dir))
	golog.Info(tmp[:len(base)])
	golog.Info(base)
	if base != "." && tmp[:len(base)] != base {
		err = errors.New("please input valid dir")
		return
	}
	d.Created = time.Now().Unix()
	d.ID, err = db.Mconn.Insert("insert into doc(name, uid,created,giturl,dir,port,kid) values(?,?,?,?,?,?,?)",
		d.Name, uid, d.Created, d.GitUrl, d.Dir, d.Port, d.KID,
	)
	return
}

func (d *Doc) MaxPort() (err error) {
	return db.Mconn.GetOne("select port from doc order by port desc limit 1").Scan(&d.Port)
}

// func (d *Doc) Startup() error {
// 	cli := client.NewClient()
// 	cli.Domain = goconfig.ReadWithoutEndSlash("scs.domain", "https://127.0.0.1:11111")
// 	// 获取最大使用的端口号
// 	err := d.MaxPort()
// 	if err != nil || d.Port == 0 {
// 		d.Port = BASEPORT
// 	}
// 	golog.Info(d.Port)
// 	for {
// 		if err := utils.CheckPort(d.Port); err == nil {
// 			break
// 		}
// 		d.Port++
// 	}

// 	script := &server.Script{
// 		Name:    d.Name,
// 		Dir:     filepath.Join(goconfig.ReadPath("scs.path"), d.Name, d.Dir),
// 		Command: fmt.Sprintf(" docsify serve . --port %d ", d.Port),
// 		Env:     map[string]string{"Path": goconfig.ReadString("scs.docsify")},
// 	}
// 	out, err := cli.AddScript(script)
// 	golog.Info(string(out))
// 	return err
// }

// func (d *Doc) StopServer() error {
// 	cli := client.NewClient()
// 	cli.Domain = goconfig.ReadWithoutEndSlash("scs.domain", "https://127.0.0.1:11111")
// 	// 获取最大使用的端口号
// 	cli.Pname = d.Name
// 	_, err := cli.RemovePnameScrip()
// 	return err
// }

func CheckKid(kid interface{}) error {
	var docCount int
	err := db.Mconn.GetOne("select count(id) from doc where kid=?", kid).Scan(&docCount)
	if err != nil {
		return err
	}
	if docCount > 0 {
		return errors.New("key exsit")
	}
	return nil
}

func NewDocById(id interface{}, uid int64) (*Doc, error) {
	doc := &Doc{}

	err := db.Mconn.GetOne("select id, name, uid,created,giturl,uptime,dir,port,kid from doc where id=? and uid=?", id, uid).Scan(
		&doc.ID, &doc.Name, &doc.UID, &doc.Created, &doc.GitUrl,
		&doc.Updated, &doc.Dir, &doc.Port, &doc.KID,
	)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
