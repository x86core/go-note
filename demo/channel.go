package main

import "fmt"

/** 通过通信来共享 **/
func main() {
	a := []int{4, 1, 3, 8, 5}

	c := make(chan int)

	go sum(a[:len(a)], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}

	c <- total
}
