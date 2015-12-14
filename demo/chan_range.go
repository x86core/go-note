package main

import (
	"fmt"
)

func FromChan(ch chan string) {
	ch <- "need to do"
	ch <- "well done"
	close(ch)
}

//routine pass data to outdoor
func main() {
	ch := make(chan string, 3) //buffer
	go FromChan(ch)
	ch <- "dd"
	for out := range ch {
		fmt.Println(out)
	}

	fmt.Println("--end--")
}
