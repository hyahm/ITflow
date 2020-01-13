package autodb

import (
	"fmt"
	"itflow/db"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hyahm/goconfig"
	"github.com/jinzhu/gorm"
)

func InitDb() {
	connstring := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
		goconfig.ReadString("mysql.user"),
		goconfig.ReadString("mysql.pwd"),
		goconfig.ReadString("mysql.host"),
		goconfig.ReadInt("mysql.port"),
		goconfig.ReadString("mysql.db"),
	)
	orm, err := gorm.Open("mysql", connstring)
	if err != nil {
		panic(err)
	}
	defer orm.Close()
	// 忽略s
	orm.SingularTable(true)
	// 自动迁移模式
	err = orm.AutoMigrate(&Apilist{},
		&Status{},
		&Projectname{},
		&Jobs{},
		&User{},
		&Header{},
		&Version{},
		&Rolegroup{},
		&Restfulname{},
		&Apiproject{},
		&Usergroup{},
		&Email{},
		&Statusgroup{},
		&Types{},
		&Level{},
		&Bugs{},
		&Options{},
		&Defaultvalue{},
		&Importants{},
		&Sharefile{},
		&Log{},
		&Headerlist{},
		&Informations{},
		&Environment{},
		&Roles{}).Error
	if err != nil {
		panic(err)
	}

	var count int
	row, err := db.Mconn.GetOne("select count(id) from user")
	if err != nil {
		panic(err)
	}
	err = row.Scan(&count)
	if err != nil {
		panic(err)
	}
	if count == 0 {
		_, err = db.Mconn.Insert("insert into user(nickname,password,email,createtime,realname,level) values(?,?,?,?,?,?)",
			"admin", "69ad5117e7553ecfa7f918a223426dd8da08a57f", "admin@qq.com", time.Now().Unix(), "admin", 0)
		if err != nil {
			panic(err)
		}
	}

	err = db.Mconn.GetOne("select count(status) from defaultvalue").Scan(&count)
	if err != nil {
		panic(err)
	}
	if count != 1 {
		//清空表
		_, err = db.Mconn.Update("truncate defaultvalue")
		if err != nil {
			panic(err)
		}
		_, err = db.Mconn.Insert("insert into defaultvalue(status,important,level) values(0,0,0)")
		if err != nil {
			panic(err)
		}
	}

	// 角色表
	//清空表

	_, err = db.Mconn.Update("truncate roles")
	if err != nil {
		panic(err)
	}
	_, err = db.Mconn.Insert("insert into roles(role) values(?),(?),(?),(?),(?),(?),(?),(?),(?),(?),(?)",
		"version", "project", "env", "status", "log", "statusgroup", "rolegroup", "important", "level",
		"position", "usergroup")
	if err != nil {
		panic(err)
	}

	// 增加类型
	row, err = db.Mconn.GetOne("select count(id) from  types ")
	if err != nil {
		panic(err)
	}
	err = row.Scan(&count)
	if err != nil {
		panic(err)
	}
	if count == 0 {
		_, err = db.Mconn.Insert("insert into types(`name`,`default`) values(?,?),(?,?),(?,?),(?,?),(?,?),(?,?),(?,?),(?,?),(?,?),(?,?),(?,?),(?,?),(?,?)",
			"int", "0", "string", "", "int64", "0", "double", "0.00", "bool", "false", "int8", "0", "int16", "0", "uint8", "0",
			"uint16", "0", "uint32", "0", "uint64", "0", "float32", "0", "float64", "0")
		if err != nil {
			panic(err)
		}
	}
	// if goconfig.ReadBool("apihelp") {
	// 	createapi()
	// }
}
