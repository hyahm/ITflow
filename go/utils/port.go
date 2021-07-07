package utils

import (
	"fmt"
	"net"
)

func CheckPort(port int) error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	defer l.Close()
	return nil
}
