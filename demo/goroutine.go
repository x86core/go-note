package main

import (
	"fmt"
	// "time"
)

func main() {
	ch := make(chan int, 0)
	routines := 10

	for i := 0; i < routines; i++ {
		// fmt.Println("before")
		go func(ch chan int, i int) {
			<-ch
			fmt.Printf("go routine %d recv \n", i)
		}(ch, i)
		fmt.Println("do....")
		ch <- 1

		// fmt.Println("after")
	}
	// what diffrent ....
	// close(ch)

	// time.Sleep(2 * time.Second)

}
