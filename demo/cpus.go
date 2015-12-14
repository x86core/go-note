package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU()) // 多核
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		fmt.Println("one")
		wg.Done()
	}()
	go func() {
		fmt.Println("two")
		wg.Done()
	}()
	go func() {
		fmt.Println("three")
		wg.Done()
	}()

	wg.Wait()
}
