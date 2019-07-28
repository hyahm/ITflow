package bugconfig

import (
	"database/sql"
	"fmt"
	"itflow/gadb"
	"log"
)

func cacheemail(db *gadb.Db) {
	err := db.GetOne("select id,email,password,port,createuser,createbug,passbug from email").Scan(&CacheEmail.Id, &CacheEmail.EmailAddr, &CacheEmail.Password, &CacheEmail.Port, &CacheEmail.CreateUser, &CacheEmail.CreateBug, &CacheEmail.PassBug)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		fmt.Println("必须有切只有一行")
		log.Fatal(err)
	}
}
