package main

import "fmt"

type a struct {
	Name string
}

var y map[*a]int

func main() {
	y = make(map[*a]int)
	x := &a{
		Name: "cander",
	}
	y[x] = 10

	delete(y, x)
	fmt.Println(x)

}
