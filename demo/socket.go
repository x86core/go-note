package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

const (
	PORT = ":8080"
	HOST = "localhost:8080"
)

func main() {
	go Server()
	go Client()

	time.Sleep(1 * time.Second)
}

func Client() {
	conn, err := net.Dial("tcp", HOST)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(conn, "GET /HTTP/1.0\r\n\r\n")

	status, err := bufio.NewReader(conn).ReadString('\n')

	fmt.Println(status, err)
}

func Server() {
	in, err := net.Listen("tcp", PORT)

	if err != nil {

	}

	for {
		conn, err := in.Accept()

		if err != nil {
			fmt.Println(err)
		}

		go func() {
			conn.Write([]byte("string"))
		}()
	}
}
