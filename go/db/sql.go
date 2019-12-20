package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/hyahm/goconfig"
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
	}
	Mconn, err = conf.NewDb()
	if err != nil {
		panic(err)
	}

}