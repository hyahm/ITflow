package db

import (
	"fmt"
	"time"

	"github.com/hyahm/cachetable"
	"github.com/hyahm/golog"
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

var Table *cachetable.Table
var ct cachetable.CT

func InitCacheTable() {
	ct = cachetable.NewCT()
	// ct.Load(".token.db", &Token{})
	err := ct.CreateTable(TOKEN, &Token{})
	if err != nil {
		if err != cachetable.ExsitErr {
			panic(err)
		}

	}

	Table, _ = ct.Use(TOKEN)
	golog.Infof("%+v", Table)
	Table.SetKeys(TOKEN)

	go ct.Clean(time.Second * 10)
}

func SaveCacheTable() error {
	fmt.Println("save db")
	return ct.Save(".token.db")
}
