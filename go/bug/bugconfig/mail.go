package bugconfig

import (
	"database/sql"
	"log"
)

func cacheemail() {
	err := Bug_Mysql.GetOne("select id,email,password,port,createuser,createbug,passbug from email").Scan(&CacheEmail.Id, &CacheEmail.EmailAddr, &CacheEmail.Password, &CacheEmail.Port, &CacheEmail.CreateUser, &CacheEmail.CreateBug, &CacheEmail.PassBug)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		log.Fatal(err)
	}
}
