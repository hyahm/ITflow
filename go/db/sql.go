package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
	"github.com/hyahm/gomysql"
)

var Mconn *gomysql.Db

func InitMysql() {
	var err error
	conf := &gomysql.Sqlconfig{
		DbName:   goconfig.ReadString("mysql.db", "project"),
		Host:     goconfig.ReadString("mysql.host", "127.0.0.1"),
		UserName: goconfig.ReadString("mysql.user", "root"),
		Password: goconfig.ReadPassword("mysql.pwd", "123456"),
		Port:     goconfig.ReadInt("mysql.port", 3306),
		Timeout:  "5s",
	}
	Mconn, err = conf.NewDb()
	if err != nil {
		golog.Error(err)
		panic(err)
	}

}
