package main

var global int = 0
var c = make(chan int, 1)

func thread1() {
	<-c
	global = 1
	c <- 1

}

func thread2() {
	<-c
	global = 2
	c <- 1
}

func main() {
	c <- 1
	go thread1()
	go thread2()
}
