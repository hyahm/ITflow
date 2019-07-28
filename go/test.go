package main

import (
	"fmt"
	"gaencrypt"
	"io/ioutil"
	"log"
)

func main() {
	s , _ := ioutil.ReadFile("pri.key")
	fmt.Println(string(s))
	a ,err:= gaencrypt.RsaDecrypt("pje2oXUTk7pDqw87/SwNkty7ysEay3cIHDA288QgXdeOZaEiF7A6jbdByrGTSV4djRXsu2hSoqdzH7Ftb9YRfPzWRNp+iVoVkGIC4ZGktfRVViUNMSFx18e5Acsa3yg+3M7mEdHb1GzvJ+o2Lphm5P0Tpodyfov08dAIk9vfZkQ=", string(s),true)
	if err != nil {
		log.Fatal(err)
	}
	print("value:", string(a))
}
