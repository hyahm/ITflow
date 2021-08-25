package db

import (
	_ "embed"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
	"github.com/hyahm/gomysql"
)

var Mconn *gomysql.Db

func InitMysql(bugsql string) {
	var err error
	port, err := strconv.Atoi(goconfig.ReadEnv("MYSQL_PORT"))
	if err != nil {
		port = goconfig.ReadInt("mysql.port", 3306)
	}

	conf := &gomysql.Sqlconfig{
		// DbName:          goconfig.ReadString("mysql.db", "itflow"),
		Host:            goconfig.ReadEnv("MYSQL_HOST", goconfig.ReadString("mysql.host", "127.0.0.1")),
		UserName:        goconfig.ReadEnv("MYSQL_USER", goconfig.ReadString("mysql.user", "root")),
		Password:        goconfig.ReadEnv("MYSQL_PASSWORD", goconfig.ReadPassword("mysql.pwd", "123456")),
		Port:            port,
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
	Mconn, err = conn.Use(goconfig.ReadEnv("MYSQL_DB", goconfig.ReadString("mysql.db", "itflow")))
	if err != nil {
		if err.(*mysql.MySQLError).Number != 1050 {
			panic(err)
		}
	}
	rows, err := Mconn.Query(bugsql)
	if err != nil {
		if err.(*mysql.MySQLError).Number != 1050 {
			panic(err)
		}
	} else {
		rows.Close()
	}

}
