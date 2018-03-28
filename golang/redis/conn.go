package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	address := "127.0.0.1:6379"
	network := "tcp"

	conn, err := net.Dial(network, address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	n, err := conn.Write([]byte("*1\r\n$4\r\nPING\r\n"))
	fmt.Println(n, err)

}
