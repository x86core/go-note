package main

import "fmt"

func main() {
	array := [10]string{"a", "b", "c", "d", "e", "f"}

	slice1 := array[:3]

	slice2 := array[3:]
	fmt.Printf("slice1 %v slice2 %v\n", slice1, slice2)
	slice1 = append(slice1, "BANG")
	fmt.Printf("append to slice1: %v\n", slice1)
	fmt.Printf("slice2 is now corrupt: %v\n", slice2)

	fmt.Printf("%p\n", &array)

	fmt.Printf("%p\n", slice1)

}
