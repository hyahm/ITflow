package db

import (
	_ "embed"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
	"github.com/hyahm/gomysql"
)

var Mconn *gomysql.Db

func InitMysql(bugsql string) {
	var err error
	conf := &gomysql.Sqlconfig{
		// DbName:          goconfig.ReadString("mysql.db", "itflow"),
		Host:            goconfig.ReadString("mysql.host", "127.0.0.1"),
		UserName:        goconfig.ReadString("mysql.user", "root"),
		Password:        goconfig.ReadPassword("mysql.pwd", "123456"),
		Port:            goconfig.ReadInt("mysql.port", 3306),
		Timeout:         time.Second * 5,
		ReadTimeout:     time.Second * 30,
		ConnMaxLifetime: time.Hour * 4,
		MaxOpenConns:    5,
		MaxIdleConns:    5,
		MultiStatements: true,
	}
	conn, err := conf.NewDb()
	if err != nil {
		golog.Error(err)
		panic(err)
	}

	Mconn, err = conn.Use(goconfig.ReadString("mysql.db", "itflow"))
	if err != nil {
		golog.Error(err)
		panic(err)
	}

	rows, err := Mconn.Query(bugsql)
	if err != nil {
		panic(err)
	}
	rows.Close()
}
