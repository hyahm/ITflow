package main

import (
	"fmt"

	"github.com/hyahm/gomysql"
)

func main() {
	conf := &gomysql.Sqlconfig{
		UserName: "cander",
		Password: "123456",
		Port:     3306,
		Host:     "192.168.234.128",
		DbName:   "bug",
	}
	db, err := conf.NewDb()
	if err != nil {
		panic(err)
	}
	var id int
	err = db.GetOne("select count(id) from bugs where dustbin=0 and uid=1 and sid=2").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}
