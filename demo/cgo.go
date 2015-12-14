package main

/*
#include <stdlib.h>
*/
import "C"
import "fmt"

func Random() int {
	var r C.long = C.random()
	return int(r)
}
func Seed(i int) {
	C.srandom(C.uint(i))
}

func main() {
	// Seed(20)
	fmt.Println(Random())
}
