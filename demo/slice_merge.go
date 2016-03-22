package main

import "fmt"

func main() {

	var sa []string

	var sb = []string{"a", "b", "c"}

	var sc = []string{"d", "e", "f"}

	sa = append(sa, sb...)
	sa = append(sa, sc...)

	fmt.Println(sa)

}
