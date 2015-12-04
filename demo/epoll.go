package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const MAX_CONN_NUM = 5

func ReadStream(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)

	for {
		_, err := conn.Read(buf)
		if err != nil {
			return
		}

		_, err = conn.Write(buf)
		if err != nil {
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer listener.Close()

	fmt.Println("running ....")

	var curNum int = 0

	connChan := make(chan net.Conn)

	tagConn := make(chan int)

	go func() {
		for tag := range tagConn {
			curNum += tag
		}
	}()

	go func() {
		c := time.Tick(1 * time.Minute)
		for a := range c {
			fmt.Println("cur num: ", a, curNum)
		}
	}()

	for i := 0; i < MAX_CONN_NUM; i++ {
		go func() {
			for conn := range connChan {
				tagConn <- 1
				ReadStream(conn)
				tagConn <- -1

			}
		}()
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Accept:", err.Error())
			return
		}

		connChan <- conn
	}
}
