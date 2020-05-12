package db

import (
	"time"

	"github.com/hyahm/cachetable"
)

const (
	TOKEN    = "Token"
	NICKNAME = "NickName"
	ID       = "Id"
)

type Token struct {
	Token    string
	NickName string
	Id       int64
}

var CT *cachetable.Table
var ct cachetable.CT

func InitCacheTable() {
	ct = cachetable.NewCT()
	ct.Load(".token.db", &Token{})
	ct.Add(TOKEN, &Token{})

	CT, _ = ct.Table(TOKEN)

	CT.SetKeys(TOKEN)
	go ct.Clean(time.Second * 10)
}

func SaveCacheTable() error {
	return ct.Save(".token.db")
}
