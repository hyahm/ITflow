package main

import (
	"fmt"
	"time"
)

func Do(c func(args ...interface{}) ){
	start := time.Now()

	fmt.Println(time.Since(start).Seconds())
}