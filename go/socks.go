package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/armon/go-socks5"
)

func main() {
	port := 8000
	if len(os.Args) > 1 {
		p, err := strconv.Atoi(os.Args[1])
		if err == nil {
			port = p
		}
	}
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", fmt.Sprintf("0.0.0.0:%d", port)); err != nil {
		panic(err)
	}
}
