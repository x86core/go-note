package main

import (
	"fmt"
	"time"
)

func main() {
	// selectFunc()
	// Timeout()
	waitDone()
}

func selectFunc() {
	ch := make(chan int, 1)

	for {
		select {
		case ch <- 0:
		case ch <- 1:
		}

		i := <-ch

		fmt.Println("input: ", i)
	}
}

func Timeout() {
	timeout := make(chan bool, 1)
	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "string"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		timeout <- true
	}()

	select {
	case <-ch:
		fmt.Println("input")
	case <-timeout:
		fmt.Println("timeout")
	}
}

func waitDone() {
	num := 5
	ch := make(chan int, num)

	go func() {
		for i := 0; i < num; i++ {
			ch <- i
			fmt.Println("go:", i)
		}
	}()

	v := <-ch
	fmt.Println("v: ", v)
	fmt.Println("over")
}
