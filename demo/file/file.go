package main

import (
	"fmt"
	// "io"
	"os"
)

func main() {

	file, err := os.Open("./time.go")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("start to read...")
	defer file.Close()

	var i = 0 //读取次数
	for {
		b := make([]byte, 64)  // buffer 大小会，影响读取次数
		out, _ := file.Read(b) // byte

		if out == 0 {
			break
		}

		fmt.Print(string(b)) // byte to string
		i++
	}
	fmt.Println("--end--")
	fmt.Println(i) //
}
