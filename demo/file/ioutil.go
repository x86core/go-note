package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	fi, err := ioutil.ReadFile("./file.gos")

	if err != nil { //}&& err != io.EOF {
		fmt.Println(err)
	}
	var s string
	io.WriteString(fi, s)

}
