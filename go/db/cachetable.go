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

func InitCacheTable() {
	ct := cachetable.NewCT()
	ct.Add("token", Token{})

	CT, _ = ct.Table("token")

	CT.SetKeys("Token")
	go ct.Clean(time.Second * 1)
}
