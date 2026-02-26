package db

import (
	_ "embed"
	"fmt"
	"strconv"
	"time"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
	"github.com/hyahm/gosql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Mconn *gosql.Db
var Gorm *gorm.DB

func InitMysqlDatabase(bugsql string) {
	var err error
	port, err := strconv.Atoi(goconfig.ReadEnv("MYSQL_PORT"))
	if err != nil {
		port = goconfig.ReadInt("db.port", 3306)
	}
	conf := &gosql.Sqlconfig{
		// DbName:          goconfig.ReadString("mysql.db", "itflow"),
		Host:            goconfig.ReadEnv("MYSQL_HOST", goconfig.ReadString("db.host", "127.0.0.1")),
		UserName:        goconfig.ReadEnv("MYSQL_USER", goconfig.ReadString("db.user", "root")),
		Password:        goconfig.ReadEnv("MYSQL_PASSWORD", goconfig.ReadPassword("db.pwd", "123456")),
		Port:            port,
		Timeout:         time.Second * 5,
		ReadTimeout:     time.Second * 30,
		ConnMaxLifetime: time.Hour * 4,
		MaxOpenConns:    5,
		MaxIdleConns:    5,
		MultiStatements: true,
		ParseTime:       true,
	}

	conn, err := conf.NewMysqlDb()
	if err != nil {
		golog.Error(err)
		panic(err)
	}
	Mconn, err = conn.Use(goconfig.ReadEnv("MYSQL_DB", goconfig.ReadString("db.db", "itflow")))
	if err != nil {
		golog.Warn(err)
	}
	_, err = Mconn.Query(bugsql)
	if err != nil {
		golog.Warn(err)
	}

	conf.DbName = goconfig.ReadEnv("MYSQL_DATABASE", goconfig.ReadString("db.db", "itflow"))
	Gorm, err = gorm.Open(mysql.Open(conf.GetMysqlDataSource()))
	if err != nil {
		panic(fmt.Sprintf("连接数据库失败: %v", err))
	}
	golog.Info("链接 gorm 成功")
	// 测试连接
	Gorm = Gorm.Debug()
	// if err != nil {
	// 	panic(fmt.Sprintf("Ping数据库失败: %v", err))
	// }

	fmt.Println("✅ 数据库连接成功!")
}
