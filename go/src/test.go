package main

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
	"io/ioutil"
	"log"
)

func main() {
	jj, err := ioutil.ReadFile("C:\\Users\\lx\\Desktop\\email.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jj))
	name := gojsonq.New().JSONString(string(jj)).Find("mail_id")
	println(name.(string)) // Tom
}
