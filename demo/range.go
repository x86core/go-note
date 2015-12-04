package main

import (
	"fmt"
)

func main() {

	var array []int = []int{5, 4, 3, 2, 1}

	for k, v := range array {

		fmt.Println(k, ":", v)
	}

	table := map[string]string{"a": "go", "b": "for", "c": "me"}

	for k := range table {
		fmt.Println(k, "->", table[k])
	}

	var table2 map[string]string = make(map[string]string) //nil map 不能存放key value

	table2["d"] = "must"
	table2["e"] = "be"
	table2["f"] = "you"
	fmt.Println("\n")
	for k, v := range table2 {
		fmt.Println(k, "->", v)
	}

}
