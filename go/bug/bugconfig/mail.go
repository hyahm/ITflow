package bugconfig

import (
	"database/sql"
	"itflow/db"
)

func cacheemail() {
	row, err := db.Mconn.GetOne("select id,email,password,port,createuser,createbug,passbug from email")
	if err != nil {
		panic(err)
	}
	err = row.Scan(&CacheEmail.Id, &CacheEmail.EmailAddr, &CacheEmail.Password, &CacheEmail.Port, &CacheEmail.CreateUser, &CacheEmail.CreateBug, &CacheEmail.PassBug)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		panic(err)
	}
}
