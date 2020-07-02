package cache

import (
	"database/sql"
	"itflow/db"
)

func cacheemail() {
	err := db.Mconn.GetOne("select id,email,password,port,createuser,createbug,passbug from email").
		Scan(&CacheEmail.Id, &CacheEmail.EmailAddr, &CacheEmail.Password, &CacheEmail.Port, &CacheEmail.CreateUser, &CacheEmail.CreateBug, &CacheEmail.PassBug)

	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		panic(err)
	}
}
