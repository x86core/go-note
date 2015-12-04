package main

import (
	"fmt"
)

func main() {
	//string->rune
	str := []rune("中国的GBK")

	for i := 0; i < len(str); i++ {

		fmt.Println(string(str[i]))

	}

}
