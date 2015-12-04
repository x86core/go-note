package main

import (
	"fmt"
	"runtime"
)

func say(s string) {

	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Printf("%s %d\n", s, i)
	}

}

func main() {
	go say("world")

	say("go...")
}
