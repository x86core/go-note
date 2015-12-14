package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fi, err := os.Open("./file.go")

	if err != nil {
		panic(err)
	}

	defer fi.Close()

	red := bufio.NewReader(fi)
	// buf := make([]byte, 1024)
	for {

		// b, _, err := red.ReadLine()
		// fmt.Printf("%s\n", string(b))

		b, err := red.ReadBytes('\n')

		fmt.Printf("%s", string(b))
		if err != nil && err != io.EOF {
			panic(err)
		}

		if err == io.EOF {
			break
		}

	}

	fmt.Println("--over--")
}
