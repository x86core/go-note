package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, index int) {
	fmt.Println("Worker:", index)
	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	// start := make(chan bool)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}

	wg.Wait()

}
