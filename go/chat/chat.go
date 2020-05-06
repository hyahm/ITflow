package chat

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func Chat() {
	conn, err := net.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		panic(err)
	}
	ch := make(chan bool)
	go userRead(conn)
	go userWrite(conn)
	<-ch
}

func userRead(conn net.Conn) {
	for {
		output := bufio.NewScanner(conn)
		for output.Scan() {
			fmt.Println(output.Text())
		}
	}

}

func userWrite(conn net.Conn) {
	for {
		inputReader := bufio.NewReader(os.Stdin)
		input, err := inputReader.ReadString('\n')
		if err != nil {
			log.Println(err)
		}
		conn.Write([]byte(input))
	}

}
