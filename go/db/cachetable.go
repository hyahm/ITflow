package db

import (
	"time"

	"github.com/hyahm/cachetable"
)

type Token struct {
	Token    string
	NickName string
}

var CT *cachetable.Table
var ct cachetable.CT

func InitCacheTable() {
	ct = cachetable.NewCT()
	ct.Load(".token.db", &Token{})
	ct.Add("token", &Token{})

	CT, _ = ct.Table("token")

	CT.SetKeys("Token")
	go ct.Clean(time.Second * 1)
}

func SaveCacheTable() error {
	return ct.Save(".token.db")
}
