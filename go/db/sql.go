package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
	"github.com/hyahm/gomysql"
)

const SESSION = "session"

var Mconn *gomysql.Db

func InitMysql() {
	var err error
	conf := &gomysql.Sqlconfig{
		DbName:   goconfig.ReadString("mysql.db"),
		Host:     goconfig.ReadString("mysql.host"),
		UserName: goconfig.ReadString("mysql.user"),
		Password: goconfig.ReadString("mysql.pwd"),
		Port:     goconfig.ReadInt("mysql.port"),
		Timeout:  "5s",
	}
	Mconn, err = conf.NewDb()
	if err != nil {
		golog.Error(err.Error())
		panic(err)
	}

}
