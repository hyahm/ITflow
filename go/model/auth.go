package model

import (
	"encoding/json"
	"errors"
	"itflow/db"
	"time"
)

type Auth struct {
	ID       int64  `json:"id,omitempty" db:"id,default"`
	Name     string `json:"name,omitempty" db:"name,default"`
	Uid      int64  `json:"uid,omitempty" db:"uid,default"`
	Created  int64  `json:"created,omitempty" db:"created,default"`
	Pri      string `json:"pri" db:"pri"` // git 的地址
	UpTime   int64  `json:"uptime" db:"uptime"`
	Pub      string `json:"pub" db:"pub,default"`           // docfiy 文档目录
	Typ      int    `json:"typ" db:"typ,default"`           // 认证  0： 不需要认证   1： 用户密码认证  2： 密钥认证
	User     string `json:"user" db:"user,default"`         // 用户名
	Password string `json:"password" db:"password,default"` // 密码
}

type Auths struct {
	Code  int     `json:"code"`
	Auths []*Auth `json:"auths"`
	Total int     `json:"total"`
	Msg   string  `json:"msg"`
}

func (as *Auths) Marshal() []byte {
	send, _ := json.Marshal(as)
	return send
}

func GetAllAuths(uid int64) ([]Auth, error) {
	auths := make([]Auth, 0)
	result := db.Mconn.Select(&auths, "select id,name,uid,created,typ,uptime,pub,pri,user,password from auth where uid=?", uid)
	return auths, result.Err
}

func (a *Auth) Insert(uid int64) (int64, error) {
	if err := ChecKeyName(a.Name); err != nil {
		return 0, err
	}
	a.Created = time.Now().Unix()
	return db.Mconn.Insert("insert into auth(name,uid,created,pri,pub,typ,user,password) values(?,?,?,?,?,?,?,?)",
		a.Name, uid, a.Created, a.Pri, a.Pub, a.Typ, a.User, a.Password)
}

func (a *Auth) Update(uid int64) (int64, error) {
	//
	a.UpTime = time.Now().Unix()
	return db.Mconn.Update("update auth set name=?,uptime=?,pri=?,pub=?,typ=?,user=?,password=? where id=? and uid=?",
		a.Name, a.UpTime, a.Pri, a.Pub, a.Typ, a.User, a.Password, a.ID, uid)
}

func ChecKeyName(name string) error {
	//
	var count int
	err := db.Mconn.GetOne("select count(id) from auth where name=?", name).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("name duplicate")
	}
	return nil
}

func GetKeyNamesByUid(uid int64) []byte {
	//
	auths := &Auths{
		Auths: make([]*Auth, 0),
	}
	rows, err := db.Mconn.GetRows("select id,name from auth where uid=?", uid)
	if err != nil {
		auths.Code = 1
		auths.Msg = err.Error()
		return auths.Marshal()
	}
	for rows.Next() {
		auth := &Auth{}
		if err := rows.Scan(&auth.ID, &auth.Name); err != nil {
			continue
		}
		auths.Auths = append(auths.Auths, auth)
	}
	return auths.Marshal()
}

func DeleteAuth(id int64) (err error) {
	//
	_, err = db.Mconn.Delete("delete from auth where id=?", id)
	return
}

func NewAuthById(id int64) (*Auth, error) {
	auth := &Auth{}
	// ID       int64  `json:"id,omitempty"`
	// Name     string `json:"name,omitempty"`
	// Uid      int64  `json:"uid,omitempty"`
	// Created  int64  `json:"created,omitempty"`
	// Pri      string `json:"pri"` // git 的地址
	// UpTime   int64  `json:"uptime"`
	// Pub      string `json:"pub"`      // docfiy 文档目录
	// Typ      int    `json:"typ"`      // 认证  0： 不需要认证   1： 用户密码认证  2： 密钥认证
	// User     string `json:"user"`     // 用户名
	// Password string `json:"password"` // 密码
	err := db.Mconn.GetOne("select name,uid,created,pri,uptime,pub,typ,user,password from auth where id=?", id).Scan(
		&auth.Name, &auth.Uid, &auth.Created, &auth.Pri, &auth.UpTime,
		&auth.Pub, &auth.Typ, &auth.User, &auth.Password,
	)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func CheckMyKey(uid, id interface{}) error {
	var count int
	err := db.Mconn.GetOne("select count(id) from auth where uid=? and id=?", uid, id).Scan(&count)
	if err != nil {
		return err
	}
	if count != 1 {
		return errors.New("key exsit")
	}
	return nil
}
